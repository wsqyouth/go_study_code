package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var resultJSON = []byte(`[
  {
    "creative_template_id": 1465,
    "site_id": 102,
    "buying_type": 1,
    "product_type_page_type_list": [
      {
        "product_type": 31,
        "landing_page_type_list": [
          {
            "landing_page_type": "40008",
            "link_name_type_list": [
              {
                "text": "应用下载量",
                "value": "35"
              }],
              "link_page_type_list": [
              {
                "text": "微信小程序",
                "value": "40101"
              }]
          }]
      }]
  }]`)

	//var newWechatLinkNameTypeConfNamed NewWechatLinkNameTypeConfNamed
	//_ = json.Unmarshal(resultJSON, &newWechatLinkNameTypeConfNamed)
	//fmt.Println(newWechatLinkNameTypeConfNamed)

	var newWechatLinkNameTypeConf NewWechatLinkNameTypeConf
	_ = json.Unmarshal(resultJSON, &newWechatLinkNameTypeConf)
	fmt.Println(newWechatLinkNameTypeConf)

	creativeTemplateID := uint32(1465)
	siteID := uint32(102)
	buyingType := uint32(1)
	productType := uint32(31)
	landingPageType := "40008"
	linkNameTypeStr := "35"
	//for _, each := range newWechatLinkNameTypeConf {
	//	if each.CreativeTemplateID == creativeTemplateID && each.SiteID == siteID && each.BuyingType == buyingType {
	//		for _, productTypeDataTypeList := range each.ProductTypePageTypeList {
	//			if productType == productTypeDataTypeList.ProductType {
	//				for _, landingPageTypeItem := range productTypeDataTypeList.LandingPageTypeList {
	//					if landingPageType == landingPageTypeItem.LandingPageType {
	//						var linkNameTypeList []string
	//						for _, linkNameTypeItem := range landingPageTypeItem.LinkNameTypeList {
	//							linkNameTypeList = append(linkNameTypeList, linkNameTypeItem.Value)
	//						}
	//						if InArray(linkNameTypeStr, linkNameTypeList) {
	//							fmt.Println("match")
	//						}
	//					}
	//				}
	//			}
	//		}
	//	}
	//}

	weChatLinkNameTypeMap := make(map[string][]string)
	weChatLinkPageTypeMap := make(map[string][]string)

	for _, each := range newWechatLinkNameTypeConf {
		for _, productTypeDataTypeList := range each.ProductTypePageTypeList {
			for _, landingPageTypeItem := range productTypeDataTypeList.LandingPageTypeList {
				var linkNameTypeList, linkPageTypeList []string
				for _, linkNameTypeItem := range landingPageTypeItem.LinkNameTypeList {
					linkNameTypeList = append(linkNameTypeList, linkNameTypeItem.Value)
				}
				for _, linkPageTypeItem := range landingPageTypeItem.LinkPageTypeList {
					linkPageTypeList = append(linkPageTypeList, linkPageTypeItem.Value)
				}
				key := strconv.Itoa(int(each.CreativeTemplateID)) + strconv.Itoa(int(each.SiteID)) +
					strconv.Itoa(int(each.BuyingType)) + strconv.Itoa(int(productTypeDataTypeList.ProductType)) +
					landingPageTypeItem.LandingPageType
				weChatLinkNameTypeMap[key] = linkNameTypeList
				weChatLinkPageTypeMap[key] = linkPageTypeList
			}
		}
	}
	key := strconv.Itoa(int(creativeTemplateID)) + strconv.Itoa(int(siteID)) +
		strconv.Itoa(int(buyingType)) + strconv.Itoa(int(productType)) + landingPageType

	fmt.Println(key)
	fmt.Println(weChatLinkNameTypeMap)
	fmt.Println(weChatLinkPageTypeMap)
	if linkNameTypeList, ok := weChatLinkNameTypeMap[key]; ok {
		fmt.Println(linkNameTypeList)
		if InArray(linkNameTypeStr, linkNameTypeList) {
			fmt.Println("match")
		}
	}

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
			LinkPageTypeList []struct {
				Text  string `json:"text"`
				Value string `json:"value"`
			} `json:"link_page_type_list"`
		} `json:"landing_page_type_list"`
	} `json:"product_type_page_type_list"`
}

//ref:https://oktools.net/json2go

//InArray 判断是否在数组中
func InArray(needle interface{}, haystack interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				exists = true
				return exists
			}
		}
	}

	return exists
}
