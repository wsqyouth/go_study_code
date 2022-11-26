package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var resultJSON = []byte(`{"results":[
		 {
            "id": 10013,
            "sort_id": 7,
            "element_key": "social_switch",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 1000011,
            "sort_id": 8,
            "element_key": "social_cheer_text_type",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 1000909,
            "sort_id": 9,
            "element_key": "social_pag_animation_url",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 10040,
            "sort_id": 10,
            "element_key": "social_pag_animation_md5",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 1000893,
            "sort_id": 11,
            "element_key": "social_cheer_icon",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        },
        {
            "id": 1000894,
            "sort_id": 12,
            "element_key": "social_cheer_icon_dark",
            "count": 1,
            "use": "optional",
            "discard": "",
            "origin_id": 0,
            "use_old": ""
        }
	]}`)
	var specResult SpecResult
	_ = json.Unmarshal(resultJSON, &specResult)
	start := 30
	for i := 0; i < len(specResult.Results); i++ {
		specResult.Results[i].SortId = start
		start = start + 1

	}
	b, _ := json.Marshal(specResult.Results)

	fmt.Println(string(b))
}

type SpecResult struct {
	Results []struct {
		Id                  int    `json:"id" `
		SortId              int    `json:"sort_id"`
		ElementKey          string `json:"element_key"`
		Count               int    `json:"count"`
		Use                 string `json:"use"`
		Discard             string `json:"discard"`
		OriginId            int    `json:"origin_id"`
		UseOld              string `json:"use_old"`
		MinCount            int    `json:"min_count"`
		MaxCount            int    `json:"max_count"`
		SupportDynamicCount string `json:"support_dynamic_count"`
	}
}
