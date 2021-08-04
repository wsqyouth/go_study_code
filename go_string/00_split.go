package main

import (
	"fmt"
	"strings"
)

func main() {
	oldDestURL := "-&from-#---"
	args := fmt.Sprintf("&from=moments&tid=%d", 123)
	ss := oldDestURL
	// 包括"#"的话，取"#"前面的部分
	ss = strings.Split(ss, "#")[0]
	// 包括"&from"的话，取"&from"前面的部分
	ss = strings.Split(ss, "&from")[0]
	// 原生页关注组件为了能够预览出关注按钮，增加了&preview=1参数，这里需要去掉
	ss = strings.ReplaceAll(ss, "&preview=1", "")
	ss = ss + args + "#wechat_redirect"
	fmt.Println(ss)
}
