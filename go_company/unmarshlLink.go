package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var resultJSON = []byte(`[{
	"creative_template_id": 1465,
	"site_id": 102,
	"buying_type": 2,
	"product_type_page_type_list": [{
	"product_type": 31,
	"landing_page_type_list": [{
	"landing_page_type": "40008",
	"link_name_type_list": [{
	"text": "应用下载量",
	"value": "2"
	}]
	}]
	}]
	}]`)

	var newWechatLinkNameTypeConfNamed NewWechatLinkNameTypeConfNamed
	_ = json.Unmarshal(resultJSON, &newWechatLinkNameTypeConfNamed)
	fmt.Println(newWechatLinkNameTypeConfNamed)

	var newWechatLinkNameTypeConf NewWechatLinkNameTypeConf
	_ = json.Unmarshal(resultJSON, &newWechatLinkNameTypeConf)
	fmt.Println(newWechatLinkNameTypeConf)

}

type NewWechatLinkNameTypeConfNamed []struct {
	CreativeTemplateID      int                       `json:"creative_template_id"`
	SiteID                  int                       `json:"site_id"`
	BuyingType              int                       `json:"buying_type"`
	ProductTypePageTypeList []ProductTypePageTypeList `json:"product_type_page_type_list"`
}
type LinkNameTypeList struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}
type LandingPageTypeList struct {
	LandingPageType  string             `json:"landing_page_type"`
	LinkNameTypeList []LinkNameTypeList `json:"link_name_type_list"`
}
type ProductTypePageTypeList struct {
	ProductType         int                   `json:"product_type"`
	LandingPageTypeList []LandingPageTypeList `json:"landing_page_type_list"`
}

// NewWechatLinkNameTypeConf
type NewWechatLinkNameTypeConf []struct {
	CreativeTemplateID      uint32 `json:"creative_template_id"`
	SiteID                  uint32 `json:"site_id"`
	BuyingType              uint32 `json:"buying_type"`
	ProductTypePageTypeList []struct {
		ProductType         uint32 `json:"product_type"`
		LandingPageTypeList []struct {
			LandingPageType  string `json:"landing_page_type"`
			LinkNameTypeList []struct {
				Text  string `json:"text"`
				Value string `json:"value"`
			} `json:"link_name_type_list"`
		} `json:"landing_page_type_list"`
	} `json:"product_type_page_type_list"`
}

//ref:https://oktools.net/json2go
