package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Name     string
	Age      int64
	Birthday *time.Time
}

type ElemData struct {
	Value string
	Key   string
}

/*
目标:从一些元素内核里提取数据到User对象
参考:https://github.com/datawhalechina/go-talent/blob/master/10.%E5%8F%8D%E5%B0%84%E6%9C%BA%E5%88%B6.md
https://darjun.github.io/2021/05/27/godailylib/reflect/
*/
func main() {
	// 定义底层元素,该元素和db模型对应
	elemData := &ElemData{
		Key:   "Name",
		Value: "paopao",
	}
	var user User
	ce := reflect.ValueOf(&user)
	ty := reflect.TypeOf(&user)
	fmt.Println(ty)
	if ce.Kind() == reflect.Ptr && ce.IsNil() {
		//反射的字段如果是指针，并且为空的时候，需要初始化一个对象
		v := reflect.New(ce.Type().Elem())
		ce.Set(v)
	}
	keyArr := []string{"Name"}
	for _, key := range keyArr {
		fmt.Printf("key: %v, kind: %v,set: %v\n", key, ce.Kind() != reflect.Invalid, ce.Elem().FieldByName(key).CanSet())
		if ce.Kind() != reflect.Invalid && ce.Elem().FieldByName(key).CanSet() {
			ce.Elem().FieldByName(key).Set(reflect.ValueOf(elemData.Value))
		}
	}

	//定律2: 将“反射类型对象”转换为“接口类型变量”
	fmt.Println(ce.Interface().(*User))
}
