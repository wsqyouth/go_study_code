package main

import (
	"log"
	"os"
	"text/template"
)

// Data 是我们将应用于模板的数据的类型
type Data struct {
	Name string
	Age  int
}

func main() {
	// 创建一个新的模板
	tpl := `Hello, my name is {{.Name}} and I am {{.Age}} years old.`

	// 创建一个新的模板实例，并解析模板字符串
	t, err := template.New("test").Parse(tpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// 创建一些将应用于模板的数据
	d := Data{
		Name: "Alice",
		Age:  25,
	}

	// 将数据应用于模板并将结果打印到控制台
	err = t.Execute(os.Stdout, d)
	if err != nil {
		log.Print("Execute: ", err)
		return
	}
	/*
		// 创建一个文件
		f, err := os.Create("output.txt")
		if err != nil {
			log.Fatal("Create: ", err)
			return
		}
		defer f.Close()

		// 将数据应用于模板并将结果写入文件
		err = t.Execute(f, d)
		if err != nil {
			log.Print("Execute: ", err)
			return
		}
	*/
	log.Println("\nTemplate has been applied and result has been printed to console.")
}
