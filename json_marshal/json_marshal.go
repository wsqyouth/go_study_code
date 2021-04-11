//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type CardInfo struct {
//	UseButton  uint32 `json:"use_button,omitempty"`
//	ShopImgURL string `json:"shop_img_url,omitempty"`
//}
//
//type CrtInfoItem struct {
//	CardInfo *CardInfo `json:"card_info,omitempty"`
//}
//
//func main() {
//
//	cardInfo := &CardInfo{}
//	var crtInfoItem CrtInfoItem
//	crtInfoItem.CardInfo = cardInfo
//
//	b, err := json.Marshal(crtInfoItem)
//	if err != nil {
//		fmt.Println("marshal err. crtInfoItem:%v", crtInfoItem)
//	}
//	fmt.Println(string(b))
//
//}

package main

import (
	"encoding/json"
	"fmt"
)

type CardInfo struct {
	UseButton uint32 `json:"use_button,omitempty"`
	NewButton int64  `json:"new_button,omitempty"`
}

func main() {

	cardInfo := &CardInfo{
		UseButton: uint32(1),
		NewButton: int64(2),
	}

	b, err := json.Marshal(cardInfo)
	if err != nil {
		fmt.Println("marshal err. crtInfoItem:%v", cardInfo)
	}
	fmt.Println(string(b))

	var cardInfoOut CardInfo
	err = json.Unmarshal(b, &cardInfoOut)
	if err != nil {
		fmt.Println("unmarshal error: %v, b: %v", err, string(b))
	}
	fmt.Printf("%T   %v\n", cardInfo.UseButton, cardInfo.UseButton)
	fmt.Printf("%T   %v\n", cardInfo.NewButton, cardInfo.NewButton)
}
