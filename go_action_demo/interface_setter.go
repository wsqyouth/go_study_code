package main

import (
	"fmt"
	"strings"
)

type ObjSetter struct {
	src    interface{}
	setter func(s string)
}

func setProcess(objSetters []ObjSetter) {
	for _, each := range objSetters {
		each.setter((each.src).(string))
	}
}
func main() {
	type Obj struct {
		a string
		b string
	}
	obj := Obj{
		a: "hello",
		b: "world",
	}
	objSetters := []ObjSetter{
		{
			src:    obj.a,
			setter: func(s string) { obj.a = strings.ToUpper(s) },
		},
		{
			src:    obj.b,
			setter: func(s string) { obj.b = strings.ToUpper(s) },
		},
	}
	setProcess(objSetters)
	fmt.Println(obj)
}
