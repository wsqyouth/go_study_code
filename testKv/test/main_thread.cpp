#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <semaphore.h>
#include <sys/types.h>
#include <dirent.h>
#include <pthread.h>
#include <errno.h>
#include <signal.h>
#include <time.h>
#include <stdio.h>
#include <assert.h>
#include <cstdlib>

#include "localmem_kv.h"

//conf param
const int SleepTimePerLoopUs = 10;

static void *ScanThread(void *args);

int GetFastRandom()
{
    return rand();
}

static void* ScanThread(void *arg)
{
//    while (1)
//    {
        //usleep(100 * 1000);
        //printf("detach thread running...!\n");
    //}
    //printf("Leave thread1!\n");
    
    //sleep 5s防止死锁
    sleep(5);

    LocalShmMemKV * pMemKV = LocalShmMemKV::GetDetault();
    if (!pMemKV)
    {
        return NULL;
    }

    int iLoopTimeUs = SleepTimePerLoopUs;
    int RandTime =  iLoopTimeUs / 5;
	printf("Trace: iLoopTimeUs:%d", iLoopTimeUs);

    while (1)
	{
		try
		{
			//扫描当前号段已用一半，另一号段用完的借点去更新另一号段
			pMemKV->ScanAndFlush();
		}
		catch (std::exception& e)
		{
			printf("ERR: got exception(%s)!!!", e.what());
		}


		//多机器上的线程休眠不同时间，错开相同时间处理，层次性处理
		int realLoopTimeUs = GetFastRandom() % (2*RandTime);
		usleep(iLoopTimeUs + realLoopTimeUs - RandTime);
	}
    return NULL;
}

int main(int argc, char** argv)
{
    pthread_t tid;

    pthread_create(&tid, NULL, ScanThread, NULL);
    pthread_detach(tid);  // 使线程处于分离状态

    LocalShmMemKV * pMemKV = LocalShmMemKV::GetDetault();
    if(!pMemKV)
    {
        printf("pMemKv is nullptr");
        return -1;
    }
    char buffer[128];
	char file_path[1024];
 	sprintf(file_path, "%s/.mem_cache_0", "./");
	if(access(file_path, F_OK) == -1) {
		sprintf(buffer, "touch %s", file_path);
		system(buffer);
	}

	key_t shm_key = ftok(file_path, 'L');
	assert(shm_key != -1);

	int ret = pMemKV->Init(shm_key, 0666);
	assert(ret != -1);
    uint64_t seq = 0;
    while(1)
    {
        string sKey = "001_snsad";
        int ret = pMemKV->GetAndInrcSeq(sKey, seq);
        if( ret || seq == 0 )
        {
            printf("error");
            return -2;
        }
        usleep(1);
        printf("main thread get seq:%lu!\n", seq);
    }

    return 0;
}

