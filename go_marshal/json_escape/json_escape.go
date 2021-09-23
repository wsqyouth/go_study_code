package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"log"
)

//学习对特殊符号的编码问题
//Note:  https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and

type JumpApp struct {
	AppName string `json:"app_name"`
	AppPath string `json:"app_path"`
}

func (t *JumpApp) JSONWithReplace() (string, error) {
	b, e := json.Marshal(t)
	if e != nil {
		log.Fatal(e)
	}
	content := string(b)
	content = strings.Replace(content, "\\u003c", "<", -1)
	content = strings.Replace(content, "\\u003e", ">", -1)
	content = strings.Replace(content, "\\u0026", "&", -1)
	return content, nil
}

func (t *JumpApp) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func main() {
	jumpApp := JumpApp{}
	jumpApp.AppName = "gh_6ee8536f381b"
	jumpApp.AppPath = "pages/detail/detail.html?id=1112&share=true"

	//方法1：替换
	messageStr, e := jumpApp.JSONWithReplace()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(messageStr)

	//方法2：
	fmt.Println("Before Marshal", jumpApp)
	messageJSON, _ := jumpApp.JSON()
	fmt.Println("After marshal", string(messageJSON))
}
