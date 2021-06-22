package main

import (
	"fmt"
	"reflect"
)

func main() {
	o := order{orderId: 20125067, customId: "grade"}
	createQuery(o)

	e := employee{"paopao", 001, "sz", 888, "china"}
	createQuery(e)

}

// createQuery 使用接口支持任意结构体生成sql语句
func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		v := reflect.ValueOf(q)

		query := fmt.Sprintf("insert into %s(", t)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int, reflect.String:
				if i == v.NumField()-1 {
					query = fmt.Sprintf("%s%s", query, v.Type().Field(i).Name)
				} else {
					query = fmt.Sprintf("%s%s,", query, v.Type().Field(i).Name)
				}
			default:
				fmt.Println("Invalid type")
				return
			}
		}
		query = fmt.Sprintf("%s) values(", query)

		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Invalid type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Printf("query: %v\n", query)
		return
	}
	fmt.Println("unsupported type")
}

// output:
// query: insert into order(orderId,customId) values(20125067, "grade")
// query: insert into employee(name,id,address,salary,country) values("paopao", 1, "sz", 888, "china")

type order struct {
	orderId  int
	customId string
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

////////////////////////////////////////////////////////////////
/*
ref: https://golangbot.com/reflection/
notes: Clear is better than clever. Reflection is never clear. --by Rob Pike

两大类型及两大方法
The concrete type of interface{} is represented by reflect.Type and the underlying value is represented by reflect.Value.
There are two functions reflect.TypeOf() and reflect.ValueOf() which return the reflect.Type and reflect.Value respectively.

Kind()和Type()的区别
Type represents the actual type of the interface{}, eg: main.Order
Kind represents the specific kind of the type. eg: struct.

结构体的方法
The NumField() method returns the number of fields in a struct and the Field(i int) method returns the reflect.Value of the ith field.

值方法
The methods Int and String help extract the reflect.Value as an int64 and string respectively.
*/
