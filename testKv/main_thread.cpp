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


//conf param
const int SleepTimePerLoopUs = 10;

static void *ScanThread(void *args);

static void* ScanThread(void *arg)
{
    while (1)
    {
        usleep(100 * 1000);
        printf("detach thread running...!\n");
    }
    printf("Leave thread1!\n");

    return NULL;
}

int main(int argc, char** argv)
{
    pthread_t tid;

    pthread_create(&tid, NULL, ScanThread, NULL);
    pthread_detach(tid);  // 使线程处于分离状态
    while(1)
    {
        sleep(1);
        printf("main thread!\n");
    }

    return 0;
}

