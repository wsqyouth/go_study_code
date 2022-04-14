package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	//urlOp()
	//fmt.Println(ss)
	//strReplace()
	strReplaceAll()
}

// urlOp 学习url相关操作方法
func urlOp() {
	srcURL := "http://tmp-qq.com/a?ck=123"
	srcURLObj, err := url.Parse(srcURL)
	if err != nil {
		fmt.Println("parse err")
	}
	fmt.Println(srcURLObj)

	srcParam, err := url.ParseQuery(srcURLObj.RawQuery)
	if err != nil {
		fmt.Println("ParseQuery err")
	}
	fmt.Println(srcParam)

	//找key
	fmt.Println(srcParam.Get("ck")) //123
}

func strReplace() {
	replaceMatch := []string{"男", "女", "我", "你"}
	replacer := strings.NewReplacer(replaceMatch...)
	s1 := replacer.Replace("我是男生")
	fmt.Println(s1)
}

func strReplaceAll() {
	srcStr := `{"mask_enable":false,"mask_img":"","mask_hide_time":0,"cover_enable":false,"cover_img":"","cover_show_time":0,"end_cover_info":{"title_img_url":"https://wxsnsdythumb.wxs.qq.com/141/20204/snscosdownload/SZ/reserved/624168080005744d000000000f5588090000008d00004eec?m=828ec849523dc23db757e02c1db59694&ck=828ec849523dc23db757e02c1db59694&sha256=53aa836383583d47868d67291426f53714267c6e95b6643c4bb4c7be77987c66","desc":"5555","action_title":"了解更多","mask_img":"https://wxsnsdythumb.wxs.qq.com/141/20204/snscosdownload/SZ/reserved/62331d9f000c08af00000000225588090000008d00004eec?m=03aa45440b3c0511eb1670e49df8cd40&ck=03aa45440b3c0511eb1670e49df8cd40&sha256=8d1ea682b8b08d5719d3b599498ea84d3d46b5b1c879a1e48db93b2e460301dd","ambient_image_url":"https://wxsnsdythumb.wxs.qq.com/141/20204/snscosdownload/SZ/reserved/621dd83300044e38000000006a5588090000008d00004eec?m=456990aa51827663ef3ce40c87087732&ck=456990aa51827663ef3ce40c87087732&sha256=297ff6623afd54417a48cdcb050e86296549b3f3dcb239fd5f26bff7b43d79e8"}}`
	fmt.Println(srcStr)
	strNew := strings.ReplaceAll(srcStr, `"mask_enable":false"`, "")
	fmt.Println(strNew)
}
