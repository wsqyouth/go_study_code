package main

import "fmt"

type Content struct {
	Name    string   `json:"name,omitempty"`
	Changes []string `json:"changes,omitempty"`
}

func main() {
	//content := Content{
	//	Name:    "",
	//	Changes: []string{"hello"}, // nil时，len为0
	//}
	//fmt.Println(len(content.Changes))
	var content Content
	content.Name = "before"
	fmt.Println(content)

	// 直接赋值
	content = Content{
		Changes: []string{"hello"}, // nil时，len为0
	}
	fmt.Println(content)

}
