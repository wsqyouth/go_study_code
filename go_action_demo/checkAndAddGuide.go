package main

import (
	"fmt"
	"strings"
)

func main() {

	oldDestURL := "https://mp.weixin.qq.com/mp/ad_biz_info?__biz=MzA3MTA0MTczNA==&sn=ad7a166d37c83c13e25f621918a21cc0#wechat_redirect"
	fmt.Println(strings.Contains(oldDestURL, "&guide_group_id=")) //true
	newDestURL := addGroupIDAttentionDestUrl(oldDestURL)
	fmt.Println(newDestURL)
	fmt.Println(strings.Contains(newDestURL, "&guide_group_id=")) //true

}

func addGroupIDAttentionDestUrl(oldDestURL string) (newDestURL string) {
	if strings.Contains(oldDestURL, "&guide_group_id=") {
		return oldDestURL
	}
	newDestURL = oldDestURL
	isSwitch, guideGroupID := true, "2179308857666109443"
	if isSwitch && guideGroupID != "" {
		ss := strings.Split(oldDestURL, "#wechat_redirect")
		if len(ss) > 1 {
			newDestURL = ss[0] + "&guide_group_id=" + guideGroupID + "#wechat_redirect"
			for _, val := range ss[1:] {
				newDestURL += val
			}
		}
	}
	fmt.Printf("addGroupIDAttentionDestUrl: %v\n", newDestURL)
	return newDestURL
}
