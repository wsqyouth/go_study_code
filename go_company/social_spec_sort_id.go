package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//{
	//	"id": 1000249,
	//	"sort_id": 10,
	//	"element_key": "data_monitor_switch",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},
	//{
	//	"id": 1000837,
	//	"sort_id": 11,
	//	"element_key": "data_monitor_ext_click_url",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},
	//{
	//	"id": 1000909,
	//	"sort_id": 12,
	//	"element_key": "data_monitor_ext_exposure_url",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},

	//{
	//	"id": 10002,
	//	"sort_id": 7,
	//	"element_key": "social_skill_first_comment_switch",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},
	//{
	//	"id": 1000845,
	//	"sort_id": 8,
	//	"element_key": "social_skill_first_comment",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},
	//{
	//	"id": 10030,
	//	"sort_id": 9,
	//	"element_key": "hidden_comment_switch",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},
	//{
	//	"id": 10004,
	//	"sort_id": 13,
	//	"element_key": "guide_group_switch",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},
	//{
	//	"id": 10024,
	//	"sort_id": 14,
	//	"element_key": "guide_group_id",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//}
	//{
	//	"id": 1000759,
	//	"sort_id": 4,
	//	"element_key": "brand_name",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//},
	//{
	//	"id": 1000129,
	//	"sort_id": 5,
	//	"element_key": "brand_img",
	//	"count": 1,
	//	"use": "optional",
	//	"discard": "",
	//	"origin_id": 0,
	//	"use_old": "",
	//	"min_count": 0,
	//	"max_count": 0,
	//	"support_dynamic_count": ""
	//}
	var resultJSON = []byte(`{"results":[
        {
            "id": 10002,
            "sort_id": 20,
            "element_key": "social_skill_first_comment_switch",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        },
        {
            "id": 1000845,
            "sort_id": 21,
            "element_key": "social_skill_first_comment",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        },
        {
            "id": 10030,
            "sort_id": 22,
            "element_key": "hidden_comment_switch",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        },
        {
            "id": 1000249,
            "sort_id": 23,
            "element_key": "data_monitor_switch",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        },
        {
            "id": 1000837,
            "sort_id": 24,
            "element_key": "data_monitor_ext_click_url",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        },
        {
            "id": 1000909,
            "sort_id": 25,
            "element_key": "data_monitor_ext_exposure_url",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        },
        {
            "id": 10045,
            "sort_id": 26,
            "element_key": "guide_group_switch",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        },
        {
            "id": 1000876,
            "sort_id": 27,
            "element_key": "guide_group_id",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": "",
            "min_count": 0,
            "max_count": 0,
            "support_dynamic_count": ""
        }
	]}`)
	var specResult SpecResult
	_ = json.Unmarshal(resultJSON, &specResult)
	start := 25
	for i := 0; i < len(specResult.Results); i++ {
		specResult.Results[i].SortId = start
		start = start + 1

	}
	b, _ := json.Marshal(specResult.Results)

	fmt.Println(string(b))
}

type SpecResult struct {
	Results []struct {
		Id         int    `json:"id" `
		SortId     int    `json:"sort_id"`
		ElementKey string `json:"element_key"`
		Count      int    `json:"count"`
		Use        string `json:"use"`
		Discard    string `json:"discard"`
		OriginId   int    `json:"origin_id"`
		UseOld     string `json:"use_old"`
	}
}
