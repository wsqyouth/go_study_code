#include <iostream>
#include "localmem_kv.h"
using namespace std;

int main() {
    cout << "hello world" << endl;
    LocalShmMemKV * pMemKV = LocalShmMemKV::GetDetault();
    if(!pMemKV)
    {
        printf("pMemKv is nullptr");
        return -1;
    }

    uint64_t seq = 0;
    string sKey = "001_snsad";
    int ret = pMemKV->GetAndInrcSeq(sKey, seq);
    if( ret || seq == 0 )
    {
        printf("error");
        return -2;
    }
    return 0;
}
