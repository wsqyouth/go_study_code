#include "sequtil.h"
#include "math.h"
#include <string>
#include <assert.h>

#include "iLog.h"
#include "iSvrkit.h"

using namespace std;

namespace mmocadbusinessseq
{

uint64_t IDSegment::GetMiddleId()
{
    return max_id - (uint64_t)round( (max_id-min_id) / 2);
}

void SeqSegment::DebugPrint()
{
    MMDEBUG("DebugPrint");
    MMDEBUG("cur_id:%lu, isSw:%d", cur_id, isSw);
    MMDEBUG("id_segment[0], min_id:%lu, max_id:%lu", id_segment[0].min_id, id_segment[0].max_id);
    MMDEBUG("id_segment[1], min_id:%lu, max_id:%lu", id_segment[1].min_id, id_segment[1].max_id);
}

void GenBusinessKey( uint32_t iBizuin, const string & type, string & strKey )
{
    strKey = Comm::StrFormat( "%03u_%s", iBizuin, type.c_str() );
    MMDEBUG("BusinessKey:%s", strKey.c_str());
}

void GenKvKey( const string & business_key, string & strKey )
{
    strKey = Comm::StrFormat( "mmocadbusinessseq_%s", business_key.c_str() );
    MMDEBUG("KvKey:%s", strKey.c_str());
}

void GenAdBusinessKey( uint32_t iBizuin, const string & type, string & strKey )
{
    strKey = Comm::StrFormat( "mmocadbusinessseq_%03u_%s", iBizuin, type.c_str() );
    MMDEBUG("AdBusinessKey:%s", strKey.c_str());
}

int GetSecondTime()
{
    int cur_time = (int)time(NULL) / 60;
    return cur_time;
}



}

