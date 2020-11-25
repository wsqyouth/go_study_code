/*
 * 原方案每次获取ID都得读写各一次kv，单次请求需要10ms+,QPS峰值能支撑120左右， 对于原生页每次请求需要调10次左右，耗时不可接受；
 * 改为利用批量获取，每次获取一个segment(step=10000决定大小)号段的值, 缓存在共享内存里。
 * 用完之后再去kv获取新的号段，可以大大减轻耗时。
 * 双buffer优化：我们希望取号段的过程能够做到无阻塞，不需要在取号段的时候阻塞请求线程，
 * 即当号段消费到某个点时就异步的把下一个号段加载到内存中。
 * 而不需要等到号段用尽的时候才去更新号段。这样做就可以很大程度上的降低系统的TP999指标
 * 采用双buffer的方式，Leaf服务内部有两个号段缓存区segment。当前号段已下发50%时，如果下一个号段未更新，
 * 则另启一个更新线程去更新下一个号段。当前号段全部下发完后，如果下个号段准备好了则切换到下个号段为当前segment接着下发，
 * 这里的异步线程其实是采取轮询（每10s轮询一次）循环往复。
 * 目前单次请求降到500us左右， Qps可达到2000；
 * 这里有三种锁可选择mutex互斥锁，spin自旋锁, kv_spin ckv自己实现的自旋锁
 * 压测的时候发现采用spin自旋锁会出现cpu飙到400%(基本必现)，导致快速拒绝, kv_spin（小概率出现），且我们的锁性能不需要追求极致
 * 故这里采用mutex互斥锁
 */

 /* 2019 09-09
  * 改成单线程，无须加锁
  */

#ifndef MEM_KV_H_
#define MEM_KV_H_

#include <iostream>
#include <string>
#include <string.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <pthread.h>


#define KV_MAGIC     0xAABBCCDD
#define KEY_SIZE     64
#define STAT_NUM     (2048)

#ifndef USE_KV_SPINLOCK
//#define USE_KV_SPINLOCK
#endif


struct KVHead {
	uint32_t magic;
	uint64_t item_num;
#if defined(USE_SPINLOCK)
	pthread_spinlock_t mutex;
#elif defined(USE_KV_SPINLOCK)
	kv_spinlock_t spinlock;
#else
	pthread_mutex_t mutex;
#endif
	int cur_list_head;
	int cur_list_end;
};


    struct IDSegment
    {
        uint64_t max_id;
        uint64_t min_id;
        IDSegment()
        {
            min_id = 0;
            max_id = 0;
        }
        uint64_t GetMiddleId();
    };

    struct SeqSegment
    {
        uint64_t cur_id;
        IDSegment id_segment[2];
        int isSw;
        SeqSegment()
        {
            cur_id = 0;
            isSw = 0;
        }
        void DebugPrint();
    };

struct KVItem {
	KVItem() {
		key_size = 0;
		memset(&key, 0, KEY_SIZE);
		used = 0;
	}

	uint32_t key_size;
	char key[KEY_SIZE];

	SeqSegment seq_segment;
	int used;

#if defined(USE_SPINLOCK)
	pthread_spinlock_t mutex;
#elif defined(USE_KV_SPINLOCK)
	kv_spinlock_t spinlock;
#else
	pthread_mutex_t mutex;
#endif
};

struct KVCache {
	KVHead head;
	KVItem items[0];
};

class LocalShmMemKV {
public:
	LocalShmMemKV();
	virtual ~LocalShmMemKV();
	virtual int Init(key_t key, int flags, uint64_t num = STAT_NUM);

public:
	virtual int DelAndGoNext();
	virtual int GetAndInrcSeq(std::string &key, uint64_t & seq);

	virtual int GetByPos(std::string &key, uint64_t &count, int &key_time, int pos);
	virtual int GetItem(std::string &key, uint64_t &count, int &key_time, int pos);

public:
	virtual int GetHeadPos(){
		LockHead();
		int head = m_pKVCache->head.cur_list_head;
		UnlockHead();
		return head;
	};
	virtual int GetEndPos(){
		LockHead();
		int end = m_pKVCache->head.cur_list_end;
		UnlockHead();
		return end;
	};
	virtual int IsFull();
	virtual int Size(){return m_pKVCache->head.item_num;};
	virtual int BucketSize(){return STAT_NUM;};
	virtual int ResetCache();

	int GetNextSeq(const int cur_pos, uint64_t & seq);
	int ScanAndFlush();
	int FlushSegment(const int cur_pos);

	int GetNextSeqUnite(const string &key, uint64_t & seq, int &step);

public:
	static LocalShmMemKV* GetDetault();

public:
	virtual void DebugPrint();

private:
	KVCache *m_pKVCache;

private:
	void inline LockHead()
	{
//#if defined(USE_SPINLOCK)
//		pthread_spin_lock(&m_pKVCache->head.mutex);
//#elif defined(USE_KV_SPINLOCK)
//		spin_lock(&m_pKVCache->head.spinlock);
//#else
//		pthread_mutex_lock(&m_pKVCache->head.mutex);
//#endif
	}

	void inline UnlockHead()
	{
//#if defined(USE_SPINLOCK)
//		pthread_spin_unlock(&m_pKVCache->head.mutex);
//#elif defined(USE_KV_SPINLOCK)
//		spin_unlock(&m_pKVCache->head.spinlock);
//#else
//		pthread_mutex_unlock(&m_pKVCache->head.mutex);
//#endif
	}

	void inline LockItem(int index)
	{
//#ifdef USE_SPINLOCK
//		pthread_spin_lock(&m_pKVCache->items[index].mutex);
//#elif defined(USE_KV_SPINLOCK)
//		spin_lock(&m_pKVCache->items[index].spinlock);
//#else
//		pthread_mutex_lock(&m_pKVCache->items[index].mutex);
//#endif
	}

	void inline UnlockItem(int index)
	{
//#ifdef USE_SPINLOCK
//		pthread_spin_unlock(&m_pKVCache->items[index].mutex);
//#elif defined(USE_KV_SPINLOCK)
//		spin_unlock(&m_pKVCache->items[index].spinlock);
//#else
//		pthread_mutex_unlock(&m_pKVCache->items[index].mutex);
//#endif
	}
};


#endif /* MEM_KV_H_ */
