package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var resultJSON = []byte(`{"results":[
		 {
            "id": 1000750,
            "sort_id": 7,
            "element_key": "info_bar_switch",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 10028,
            "sort_id": 8,
            "element_key": "info_bar_type",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 1000885,
            "sort_id": 9,
            "element_key": "info_bar_title",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 10016,
            "sort_id": 10,
            "element_key": "info_bar_description",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 10026,
            "sort_id": 11,
            "element_key": "info_bar_head_description",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 20182,
            "sort_id": 12,
            "element_key": "info_bar_head_image_id",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        }
	]}`)
	var specResult SpecResult
	_ = json.Unmarshal(resultJSON, &specResult)
	start := 9
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
