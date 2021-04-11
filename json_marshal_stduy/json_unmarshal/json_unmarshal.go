package main

import (
	"encoding/json"
	"fmt"
)

//重点学习使用marshal提取json成员类型信息，思考：若为通用结构如何提取成员类型？

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
