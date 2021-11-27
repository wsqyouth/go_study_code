package main

import "fmt"

func main() {
	fmt.Println("vm-go")
}

curl -X POST  http://127.0.0.1:8007/trpc.dp.dpsync.dpsync/SearchTemplate  -H 'Content-Type: application/json'   -H 'cache-control: no-cache'   -d '{"uid":169051,"adcreative_template_id":641,"promoted_object_type_list":["PROMOTED_OBJECT_TYPE_LEAD_AD","PROMOTED_OBJECT_TYPE_ECOMMERCE","PROMOTED_OBJECT_TYPE_APP_IOS","PROMOTED_OBJECT_TYPE_MINI_GAME_WECHAT","PROMOTED_OBJECT_TYPE_APP_ANDROID"],"site_id":21,"buying_type":"BUYINGTYPE_AUCTION"}'
