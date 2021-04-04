package main

import (
	"encoding/json"
	"fmt"
)

type CardInfo struct {
	UseButton  uint32 `json:"use_button,omitempty"`
	ShopImgURL string `json:"shop_img_url,omitempty"`
}

type CrtInfoItem struct {
	CardInfo *CardInfo `json:"card_info,omitempty"`
}

func main() {

	cardInfo := &CardInfo{}
	var crtInfoItem CrtInfoItem
	crtInfoItem.CardInfo = cardInfo

	b, err := json.Marshal(crtInfoItem)
	if err != nil {
		fmt.Println("marshal err. crtInfoItem:%v", crtInfoItem)
	}
	fmt.Println(string(b))

}
