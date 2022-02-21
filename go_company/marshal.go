package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	avatarImgUrlList := []string{
		"5ffeefb400097d4200000000a7f357090000008d00004eec", "5ffeefb400097d4200000000a7f357090000008d00004eec"}
	avatarBytes, err := json.Marshal(avatarImgUrlList)
	if err != nil {
		fmt.Printf("marshal Avatars error: %v, avatarImgList: %v", err,
			avatarImgUrlList)
		return
	}
	fmt.Println((avatarBytes))
	fmt.Println(string(avatarBytes))
}
