package main

import (
	"fmt"
)

type T struct {
	a int
	b float64
	c string
}

func main() {
	fmt.Println("hello")
	t := &T{a: 7, b: -2.35, c: "abc\tdebf"}
	fmt.Printf("%v\n", t) //&{7 -2.35 abc   debf}
	//打印了具体字段名
	fmt.Printf("%+v\n", t) //&{a:7 b:-2.35 c:abc     debf}
	fmt.Printf("%#v\n", t) //&main.T{a:7, b:-2.35, c:"abc\tdebf"}

	str := "hello 世界"
	fmt.Printf("%q\n", str)  //"hello 世界"
	fmt.Printf("%#q\n", str) //`hello 世界`
	fmt.Printf("%+q\n", str) //"hello \u4e16\u754c"
	fmt.Printf("%x\n", str)  //68656c6c6f20e4b896e7958c
	fmt.Printf("% x\n", str) //68 65 6c 6c 6f 20 e4 b8 96 e7 95 8c

}

//referenc:https://jingwei.link/2019/10/26/effectivego-map-fmt-append.html
