#!/usr/bin/python
#coding:utf-8

import sys
from optparse import OptionParser
import pprint
import time
import requests
import xml.dom.minidom

def mannual(ip, port, operator, uid, id):

    url = 'http://%s:%s/trpc.dp.dpsync.dpsync/ManualTool' % (ip,port)
    data = {'operator':operator,'uid':uid,'id':id}
    print(url)
    print(data)
    r = requests.post(url,data)
    print(r.url)
    print(r.status_code)
    resp = mmappsvr_pb2.MMBaseAppResponse()
    resp.ParseFromString(r.text.encode('utf8'))
    print('result:', resp.Result, 'errcode:', resp.ErrCode, 'message:', resp.ErrMsg)
    return

def main():
    global ip, port, operator, uid, id

    parser = OptionParser()
    parser.add_option("-s","--server",dest="server_name" , help="get check file in server" )

    parser.add_option("-i", "--ip",dest="ip", help="Synchronization container IP")
    parser.add_option("-p", "--port", dest="port", help="Synchronization container port")
    parser.add_option("-o", "--operator", dest="operator", help="Options:InsertCreativeByAid;SyncAdGroup;SyncAdCreativeByAid")
    parser.add_option("-z", "--uid", dest="uid", help="Advertiser UID")
    parser.add_option("-a", "--id", dest="id", help="Advertising aid")

    (options, args) = parser.parse_args()
    ip = options.ip
    port = options.port
    operator = options.operator
    uid = options.uid
    id = options.id
    if(not (ip  or port or operator or uid or id )):
        parser.print_help()
        sys.exit(-1)
    else:
        mannual(ip, port, operator, uid, id)


if __name__ == "__main__":
    main()



参数示例:
-i 11.160.168.232 -p 8007 -o InsertCreativeByAid -z 95918 -a 3149498648