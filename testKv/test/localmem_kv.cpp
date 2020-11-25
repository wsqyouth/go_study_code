#include "localmem_kv.h"

#include <vector>
#include <string>
#include <cmath>
using namespace std;

///////////////////////////////////////////////////////////////////////////////

uint64_t IDSegment::GetMiddleId()
{
    return max_id - (uint64_t)round( (max_id-min_id) / 2);
}

void SeqSegment::DebugPrint()
{
    printf("DebugPrint\n");
    printf("cur_id:%lu, isSw:%d\n", cur_id, isSw);
    printf("id_segment[0], min_id:%lu, max_id:%lu\n", id_segment[0].min_id, id_segment[0].max_id);
    printf("id_segment[1], min_id:%lu, max_id:%lu\n", id_segment[1].min_id, id_segment[1].max_id);
}

int GetSecondTime()
{
    int cur_time = (int)time(NULL) / 60;
    return cur_time;
}

uint64_t GenerateNum()
{
    static uint64_t counter = 0;
    counter += 10;
}

void  StringSplit(const std::string& srcStr, const std::string& delim, std::vector<std::string>& vecResult)
{
	size_t last = 0;
	size_t index = srcStr.find_first_of(delim, last);
	while (index != std::string::npos)
	{
		vecResult.push_back(srcStr.substr(last, index - last));
		last = index + 1;
		index = srcStr.find_first_of(delim, last);
	}
	if (index - last > 0)
	{
		vecResult.push_back(srcStr.substr(last, index - last));
	}
}


LocalShmMemKV::LocalShmMemKV()
{
	m_pKVCache = NULL;
}
LocalShmMemKV::~LocalShmMemKV()
{

}

LocalShmMemKV* LocalShmMemKV::GetDetault()
{
	static LocalShmMemKV st_LocalShmMemKV;
	return &st_LocalShmMemKV;
}

int LocalShmMemKV::Init(key_t key, int flags, uint64_t num)
{
	printf("Enter Init\n");
	int shmid = 0;
	bool is_first = false;

	shmid = shmget(key, 0, flags);
	if(shmid == -1) {
		printf("Create new share memory\n");
		shmid = shmget(key, sizeof(KVHead) + sizeof(KVItem) * num, flags | IPC_CREAT);
		if(shmid == -1) {
			printf("shmid == -1, %s\n", strerror(errno));
			return -1;
		}
		is_first = true;
	} else {
		printf("Get existed share memory\n");
	}

	printf("shmid: %d\n", shmid);

	shmid_ds buf;
	if(shmctl(shmid, IPC_STAT, &buf) == -1) {
		printf("shmctl fail\n");
		return -1;
	}
	int real_size = buf.shm_segsz;
	if(real_size != (sizeof(KVHead) + sizeof(KVItem) * num)) {
		printf("real_size not match, real_size:%d, size:%lu\n", real_size, (sizeof(KVHead) + sizeof(KVItem) * num));
		return -1;
	}

	void *p_ret = shmat(shmid, 0, 0);
	if(p_ret == (void *)-1) {
		printf("shmat failed\n");
		return -1;
	}
	m_pKVCache = (KVCache *)p_ret;

	if(is_first) {
		printf("First use that share memory[%d], clear all\n", real_size);
		m_pKVCache->head.cur_list_end = 0;
		m_pKVCache->head.cur_list_head = 0;
		m_pKVCache->head.item_num = 0;
		m_pKVCache->head.magic = KV_MAGIC;
		for(int i = 0; i < num; i++) {
			//init value
			m_pKVCache->items[i].key_size = 0;
			m_pKVCache->items[i].used = 0;
			memset(m_pKVCache->items[i].key, 0, sizeof(m_pKVCache->items[i].key));
		}
	} else {
		printf("Use existed share memory, shm size; %d\n", real_size);
		if(m_pKVCache->head.magic != KV_MAGIC) {
			printf("magic wrong\n");
			return -1;
		}
	}

#if defined(USE_SPINLOCK)
	int status = pthread_spin_init(&m_pKVCache->head.mutex, PTHREAD_PROCESS_SHARED);
	if(status) {
		printf("init spin failed. %d\n", status);
		return -1;
	}
#elif defined(USE_KV_SPINLOCK)
	spin_lock_init(&m_pKVCache->head.spinlock);
#else
	pthread_mutexattr_t mutex_attr;
	int status = pthread_mutexattr_init(&mutex_attr);
	if(status) {
		printf("init mutexattr failed.%d\n", status);
		return -1;
	}

	status = pthread_mutexattr_setpshared(&mutex_attr, PTHREAD_PROCESS_SHARED);
	if(status) {
		printf("set mutex attr to shared failed. %d\n", status);
		return -1;
	}

	status = pthread_mutex_init(&m_pKVCache->head.mutex, &mutex_attr);
	if(status) {
		printf("init mutex failed. %d\n", status);
		return -1;
	}
#endif

	for(int i = 0; i< num; i++) {
#if defined(USE_SPINLOCK)
		status = pthread_spin_init(&m_pKVCache->items[i].mutex, PTHREAD_PROCESS_SHARED);
		if(status) {
			printf("init spin failed. %d\n", status);
			return -1;
		}
#elif defined(USE_KV_SPINLOCK)
		spin_lock_init(&m_pKVCache->items[i].spinlock);
#else
		status = pthread_mutex_init(&m_pKVCache->items[i].mutex, &mutex_attr);
		if(status) {
			printf("init mutex failed. %d\n", status);
			return -1;
		}
#endif

	}

	printf("Exit Init\n");
	return 0;
}

int LocalShmMemKV::IsFull() {
	LockHead();
	uint64_t item_num = m_pKVCache->head.item_num;
	UnlockHead();
	return item_num == STAT_NUM;
};

// 从kv 或者 zk中取一批号段
int LocalShmMemKV::GetNextSeqUnite(const string &key, uint64_t & max_seq, int &step)
{
	int ret = 0;
	if(key != "010_media_id")  //kv取号段
	{
		//SeqDAO seq_dao(m_poConfig);
		//std::string kv_key;
		//GenKvKey(key, kv_key);
		//step = m_poConfig->GetSeqStep();
		//ret = seq_dao.GetNextSeq(kv_key, max_seq);
        max_seq = GenerateNum();
        step = 10;
	    printf("get kv max_seq:lu\n", max_seq);
	}
	else  //zk取号段
	{
		//step = m_poConfig->GetZkStep();
		//modify by flyyfyan 2020-04013 和中台端统一素材库节点，当前这个还未使用故可以直接更换
		//string node = "/ad/image_id";
		string node = "/ad/asset_media_material_id";
		//ret = GetZkSeq(m_poConfig, "NEXT", node, max_seq);
	    printf("node. %s\n",node.c_str());
	}
	printf("wsq. %d\n", ret);
	printf("wsqNew. %s\n", key.c_str());
	printf("wsqNew. %lu. %u\n", max_seq, step);
	string uid_list;
	int gray_percent;
	// AvatarAccess::AvatarAccess::GetFieldValue<std::string>(0, "mmocadcontroller/qq_advanced_version_uid_list", uid_list);
	// AvatarAccess::AvatarAccess::GetFieldValue<int>(0, "mmocadcontroller/qq_advanced_version_gray_percent", gray_percent);

	return ret;
}

int LocalShmMemKV::GetNextSeq(const int cur_pos, uint64_t & seq)
{
	// 外部请求步长固定为1 不支持设置
	auto & tmp_seq = m_pKVCache->items[cur_pos].seq_segment;
	printf("wsq. cur_id:%lu, max_id:%lu, min_id:%lu\n", tmp_seq.cur_id,tmp_seq.id_segment[tmp_seq.isSw].max_id,tmp_seq.id_segment[tmp_seq.isSw].max_id);
	printf("wsq. m_pKVCache->items[cur_pos].key:%s\n",m_pKVCache->items[cur_pos].key);
	//左开右闭区间（min_id,max_id]
	if(tmp_seq.cur_id < tmp_seq.id_segment[tmp_seq.isSw].max_id )
	{
		tmp_seq.cur_id += 1;
		seq = tmp_seq.cur_id;
		printf("hit this segment. isSw:%d, cur_id:%lu, max_id:%lu\n", tmp_seq.isSw, tmp_seq.cur_id, tmp_seq.id_segment[tmp_seq.isSw].max_id );
	}
	else if( tmp_seq.cur_id == tmp_seq.id_segment[tmp_seq.isSw].max_id &&  tmp_seq.cur_id <= tmp_seq.id_segment[1-tmp_seq.isSw].min_id)
	{
		//切换号段
		tmp_seq.cur_id = tmp_seq.id_segment[1-tmp_seq.isSw].min_id + 1;
		seq = tmp_seq.cur_id;
		tmp_seq.isSw = 1 - tmp_seq.isSw ;
		printf("hit another segment. isSw:%d, cur_id:%lu, max_id:%lu\n", tmp_seq.isSw, tmp_seq.cur_id, tmp_seq.id_segment[tmp_seq.isSw].max_id);
	}
	else
	{
		//需要从kv or zk中取号段;
		uint64_t max_seq = 0;
		int step = 0;
		int ret = GetNextSeqUnite(m_pKVCache->items[cur_pos].key, max_seq, step);
		printf("wsq. m_pKVCache->items[cur_pos].key:%s,step:%d\n",m_pKVCache->items[cur_pos].key,step);
		if( 0 != ret )
		{
			//查询kv失败
			printf( "ERR: call kv failed! ret:%d, key:%s\n", ret, m_pKVCache->items[cur_pos].key);
			return ret;
		}
		//刷新缓存号段
		tmp_seq.id_segment[tmp_seq.isSw].max_id = max_seq;
		tmp_seq.id_segment[tmp_seq.isSw].min_id = max_seq - step;
		tmp_seq.cur_id = tmp_seq.id_segment[tmp_seq.isSw].min_id + 1;
		seq = tmp_seq.cur_id;
		printf("hit kv. isSw:%d, cur_id:%lu, max_id:%lu\n", tmp_seq.isSw, tmp_seq.cur_id, tmp_seq.id_segment[tmp_seq.isSw].max_id);
	}
	return 0;
}

int LocalShmMemKV::GetAndInrcSeq(std::string &key, uint64_t & seq)
{
	LockHead();
	int list_head = m_pKVCache->head.cur_list_head;
	int list_end = m_pKVCache->head.cur_list_end;
	uint64_t item_num = m_pKVCache->head.item_num;
	UnlockHead();

	int need_loop_num = (list_end - list_head + STAT_NUM) % STAT_NUM;
	if(item_num == STAT_NUM) {
		need_loop_num = STAT_NUM;
	}
	printf("wsq. list_head:%d,list_end:%d,item_num:%lu\n", list_head,list_end,item_num);
	int cur_pos = list_head;
	int ret = 0;
	for(int i = 0; i < need_loop_num; i++)
	{
		LockItem(cur_pos);
		if(m_pKVCache->items[cur_pos].used)
		{
			UnlockItem(cur_pos);
			if(m_pKVCache->items[cur_pos].key_size == key.length() && (strncmp(m_pKVCache->items[cur_pos].key, key.c_str(), key.length()) == 0))
			{
				LockItem(cur_pos);
				ret = GetNextSeq(cur_pos, seq);
				UnlockItem(cur_pos);
				return ret;
			}
			else
			{
				cur_pos++;
				if(cur_pos >= STAT_NUM)
				{
					cur_pos = cur_pos % STAT_NUM;
				}
				continue;
			}
		}
		else
		{
			UnlockItem(cur_pos);
		}

		cur_pos++;
		if(cur_pos >= STAT_NUM)
		{
			cur_pos = cur_pos % STAT_NUM;
		}
	}

	LockHead();
	if(m_pKVCache->head.item_num == STAT_NUM) {
		//no bucket
		UnlockHead();
		printf("all bucket ia full.\n");
		return -1;
	}

	//new item
	if(m_pKVCache->items[m_pKVCache->head.cur_list_end].used) {
		//no bucket
		UnlockHead();
		printf("used: %d, no bucket\n", m_pKVCache->items[m_pKVCache->head.cur_list_end].used);
		return -1;
	}
	m_pKVCache->head.item_num++;
	list_end = m_pKVCache->head.cur_list_end++;
	if(m_pKVCache->head.cur_list_end >= STAT_NUM) {
		m_pKVCache->head.cur_list_end = m_pKVCache->head.cur_list_end % STAT_NUM;
	}
	UnlockHead();

	LockItem(list_end);
	//第一次找不到key, 同样需要从kv中拿号段
	uint64_t max_seq = 0;
	int step = 0;
	ret = GetNextSeqUnite(key, max_seq, step);
	if( 0 != ret )
	{
		//查询kv失败
		UnlockItem(list_end);
		printf( "ERR: call kv failed! ret:%d, key:%s\n", ret, key.c_str());
		return ret;
	}

	memcpy(m_pKVCache->items[list_end].key, key.c_str(), key.length());
	m_pKVCache->items[list_end].key_size = key.length();
	m_pKVCache->items[list_end].used = 1;
	//刷新缓存号段
	auto & tmp_seq = m_pKVCache->items[cur_pos].seq_segment;
	tmp_seq.id_segment[tmp_seq.isSw].max_id = max_seq;
	tmp_seq.id_segment[tmp_seq.isSw].min_id = max_seq - step;
	tmp_seq.cur_id = tmp_seq.id_segment[tmp_seq.isSw].min_id + 1;
	seq =  m_pKVCache->items[cur_pos].seq_segment.cur_id;
	printf("hit kv and set mem cache. isSw:%d, cur_id:%lu, max_id:%lu\n", tmp_seq.isSw, tmp_seq.cur_id, tmp_seq.id_segment[tmp_seq.isSw].max_id);
	UnlockItem(list_end);

	return 0;
}


int LocalShmMemKV::FlushSegment(const int cur_pos)
{
	auto & tmp_seq = m_pKVCache->items[cur_pos].seq_segment;
	printf("FlushSegment segment key:%s\n", m_pKVCache->items[cur_pos].key);
	tmp_seq.DebugPrint();
	//左开右闭区间（min_id,max_id]
	if(tmp_seq.cur_id >= tmp_seq.id_segment[tmp_seq.isSw].GetMiddleId() && tmp_seq.cur_id >= tmp_seq.id_segment[1-tmp_seq.isSw].max_id )
	{
		uint64_t max_seq = 0;
		int step = 0;
		int ret = GetNextSeqUnite(m_pKVCache->items[cur_pos].key, max_seq, step);
		if(0 != ret)
		{
			//查询kv or zk 失败
			printf( "ERR: flsuh kv_key failed! ret:%d, key:%s\n", ret, m_pKVCache->items[cur_pos].key);
			return ret;
		}
		//刷新缓存号段
		tmp_seq.id_segment[1-tmp_seq.isSw].max_id = max_seq;
		tmp_seq.id_segment[1-tmp_seq.isSw].min_id = max_seq - step;
		printf( "ERR: flsuh kv_key suc! key:%s, min_id:%lu, max_id:%lu\n", m_pKVCache->items[cur_pos].key, tmp_seq.id_segment[1-tmp_seq.isSw].min_id, max_seq);
	}
	return 0;
}


int LocalShmMemKV::ScanAndFlush()
{
	LockHead();
	int list_head = m_pKVCache->head.cur_list_head;
	int list_end = m_pKVCache->head.cur_list_end;
	uint64_t item_num = m_pKVCache->head.item_num;
	UnlockHead();

	if(item_num == 0)
	{
		//空数组直接返回
		printf("Shm Mem Kv empty\n");
		return 0;
	}

	int need_loop_num = (list_end - list_head + STAT_NUM) % STAT_NUM;
	if(item_num == STAT_NUM) {
		need_loop_num = STAT_NUM;
	}

	int cur_pos = list_head;
	int ret = 0;
	for(int i = 0; i < need_loop_num; i++)
	{
		LockItem(cur_pos);
		if(m_pKVCache->items[cur_pos].used)
		{
			ret = FlushSegment(cur_pos);
			UnlockItem(cur_pos);
		}
		else
		{
			UnlockItem(cur_pos);
		}

		cur_pos++;
		if(cur_pos >= STAT_NUM)
		{
			cur_pos = cur_pos % STAT_NUM;
		}
	}

	return 0;
}



/*
 * Not thread safe
 */
int LocalShmMemKV::GetItem(std::string &key, uint64_t &count, int &key_time, int pos)
{
	printf("Enter GetItem");

	if(pos < 0 || pos >= STAT_NUM) {
		printf("panic, pos: %d not at the right range[0, %d)\n", pos, STAT_NUM);
		return -1;
	}

	LockItem(pos);
	if(m_pKVCache->items[pos].used == 0) {
		UnlockItem(pos);
		printf("This item is not used, skip\n");
		return -2;
	}
	UnlockItem(pos);

	std::vector<std::string> vec;
	key.assign(m_pKVCache->items[pos].key, m_pKVCache->items[pos].key_size);
	StringSplit("_", key, vec);
	if(vec.size() != 2) {
		m_pKVCache->items[pos].used = 0; //delete bad key
		printf("bad name at share memory: %s, skip\n", key.c_str());
		return  -2;
	}

	key_time = atoi(vec[1].c_str());
	if((key_time + 1) < GetSecondTime()) {
		//count = m_pKVCache->items[pos].count; //have a time diff, so no need mutex
		printf("Exit GetItem\n");
		return 0;
	}

	//finish the loop
	printf("Exit GetItem\n");
	return -3;
}
/*
 * Not thread safe
 */
int LocalShmMemKV::GetByPos(std::string &key, uint64_t &count, int &key_time, int pos)
{
	printf("Enter GetByPos\n");

	if(pos < 0 || pos >= STAT_NUM) {
		printf("panic, pos: %d not at the right range[0, %d)\n", pos, STAT_NUM);
		return -1;
	}

	if(m_pKVCache->items[pos].used == 0) {
		printf("This item is not used, skip\n");
		return -2;
	}

	std::vector<std::string> vec;
	key.assign(m_pKVCache->items[pos].key, m_pKVCache->items[pos].key_size);
	StringSplit("_", key, vec);
	if(vec.size() != 2) {
		m_pKVCache->items[pos].used = 0; //delete bad key
		printf("bad name at share memory: %s, skip\n", key.c_str());
		return  -2;
	}

	key_time = atoi(vec[1].c_str());
	if(key_time < GetSecondTime()) {
		//count = m_pKVCache->items[pos].count; //have a time diff, so no need mutex
		printf("Exit GetByPos\n");
		return (pos + 1) % STAT_NUM;
	}

	//finish the loop
	printf("Exit GetByPos\n");
	return -3;
}


/*
 * Not thread safe
 */
int LocalShmMemKV::DelAndGoNext()
{
	printf("Enter DelAndGoNext\n");

	LockHead();
	m_pKVCache->head.item_num--;
	int list_head = m_pKVCache->head.cur_list_head++;
	if(m_pKVCache->head.cur_list_head >= STAT_NUM) {
		m_pKVCache->head.cur_list_head = m_pKVCache->head.cur_list_head % STAT_NUM;
	}
	UnlockHead();

	LockItem(list_head);
	m_pKVCache->items[list_head].used = 0;
	UnlockItem(list_head);

	printf("Exit DelAndGoNext\n");
	return 0;
}

int LocalShmMemKV::ResetCache()
{
	return -1;
}

void LocalShmMemKV::DebugPrint()
{
	printf("Enter DebugPrint\n");

	printf("\nhead {\n\t{item_num, %lu}\n\t{cur_list_head, %d}\n\t{cur_list_end, %d}\n\t{magic, %x}\n}\n",
			m_pKVCache->head.item_num, m_pKVCache->head.cur_list_head, m_pKVCache->head.cur_list_end, m_pKVCache->head.magic);

	std::string key;
	int need_loop_num = (m_pKVCache->head.cur_list_end - m_pKVCache->head.cur_list_head + STAT_NUM) % STAT_NUM;
	int cur_pos = m_pKVCache->head.cur_list_head;
	for(int i = 0; i < need_loop_num; i++) {
		printf("\n{\n\t{pos, %d}\n\t{key, %s}\n\t{used, %d}\n\t{used, %d}\n}",
				cur_pos, key.assign(m_pKVCache->items[cur_pos].key, m_pKVCache->items[cur_pos].key_size).c_str(),
				m_pKVCache->items[cur_pos].used, m_pKVCache->items[cur_pos].used);
		cur_pos++;
		if(cur_pos >= STAT_NUM) {
			cur_pos = cur_pos % STAT_NUM;
		}
	}

	for(int i = 0; i < STAT_NUM; i++) {
		if(m_pKVCache->items[i].used == 0) {
			continue;
		}
		printf("\n{\n\t{pos, %d}\n\t{key, %s}\n\t{used, %d}\n\t{used, %d}\n}",
				i, key.assign(m_pKVCache->items[i].key, m_pKVCache->items[i].key_size).c_str(),
				m_pKVCache->items[i].used,	m_pKVCache->items[i].used);
	}

	printf("Exit DebugPrint\n");
}

