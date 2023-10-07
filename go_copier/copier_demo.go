package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	source := Person{
		Name: "John Doe",
		Age:  30,
	}

	destination := new(Person)
	if err := copier.Copy(&destination, &source); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Copied person:\n %+v\n", destination)
	}
}

/*
代码创建了两个具有不同地址但相似字段的结构体实例。通过调用 `copier.Copy()` 函数，我们将源结构体的所有字段值复制到目标结构体实例中。
*/
