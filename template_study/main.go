package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Http server failed. err:%v\n", err)
		return
	}
	fmt.Println("vim-go")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed")
		return
	}

	// 利用指定的模板进行渲染,并将结果写入w
	tmpl.Execute(w, "清水河畔")
}
