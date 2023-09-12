package main

import (
	"fmt"
	"reflect"
)

type ProductSpec struct {
	ProductImgList []string
}

func main() {
	ce := &ProductSpec{}
	keyArr := []string{"ProductSpec", "ProductImgList"}
	elementData := struct{ Value string }{Value: "new_element"}
	// Check if ce is nil
	if ce == nil {
		fmt.Println("ce is nil")
		return
	}
	// Get the value of ce
	value := reflect.ValueOf(ce)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	// Check if the field exists
	field := value.FieldByName("ProductImgList")
	if !field.IsValid() {
		fmt.Printf("Field %s does not exist\n", "ProductImgList")
		return
	}

	// Initialize the slice if it's nil
	if ce.ProductImgList == nil {
		ce.ProductImgList = make([]string, 0)
	}
	ce.ProductImgList = append(ce.ProductImgList, "old_element")
	elementProductVal := reflect.ValueOf(&ce.ProductImgList).Elem()
	ProductImgList, ok := elementProductVal.Interface().([]string)
	if !ok {
		fmt.Printf("not support key: %s, val: %s\n", keyArr[1], elementProductVal)
		return
	}
	ProductImgList = append(ProductImgList, elementData.Value)
	elementProductVal.Set(reflect.ValueOf(ProductImgList))
	fmt.Println(ce.ProductImgList) // Output: [old_element new_element]
}

/*
思考：针对指针类型的变量，当我们通过反射获取接口值后,如何安全的找到成员变量并对其修改呢？
1. 获取接口值value，实现对其的解引用 [反射的第一定律是：“反射可以从接口值（interface）得到反射对象]
2. 检查对应的成员变量是否存在,不存在则报错
3. 通过反射值取出原有成员变量的内容，然后添加新的内容 [反射的第二定律是：可以从反射对象得到接口值（interface）]
4. 重新set进去。[反射的第三定律是：“要修改反射对象，该值必须可以修改”。]

第三条定律看上去与第一、第二条均无直接关联，但却是必不可少的，因为反射在工程实践中，目的一就是可以获取到值和类型，其二就是要能够修改他的值。

参考：https://golang1.eddycjy.com/posts/ch2/reflect/#:~:text=Go%20%E8%AF%AD%E8%A8%80%E4%B8%AD%E7%9A%84%E5%8F%8D%E5%B0%84,reflection%20object%20to%20interface%20value.
*/
