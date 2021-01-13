package main

import "fmt"
import "encoding/json"
import "bytes"

type JumpApp struct {
	AppName string `json:"app_name"`
	AppPath string `json:"app_path"`
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
	fmt.Println("Before Marshal", jumpApp)
	messageJSON, _ := jumpApp.JSON()
	fmt.Println("After marshal", string(messageJSON))
}

//Note:
//https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and
