package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// v.SetFloat(7.1) // Error: will panic.

	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface()) // 7.1
	fmt.Println(x)             // 7.1
}

/*
如上面这段代码，运行以后会崩溃，崩溃信息是 panic: reflect: reflect.Value.SetFloat using unaddressable value，
为什么这里 SetFloat() 会 panic 呢？这里给的提示信息是使用了不可寻址的 Value。

在上述代码中，调用 reflect.ValueOf 传进去的是一个值类型的变量，获得的 Value 其实是完全的值拷贝，这个 Value 是不能被修改的。
如果传进去是一个指针，获得的 Value 是一个指针副本，但是这个指针指向的地址的对象是可以改变的。
*/

/*
总结：
1. 反射可以从接口值中得到反射对象
* 通过实例获取 Value 对象，使用 reflect.ValueOf() 函数。
* 通过实例获取反射对象 Type，使用 reflect.TypeOf() 函数。
2. 反射可以从反射对象中获得接口值
* 将 Value 转换成空的 interface，内部存放具体类型实例。使用 interface() 函数。
* Value 也包含很多成员方法，可以将 Value 转换成简单类型实例，注意如果类型不匹配会 panic。
3. 若要修改反射对象，值必须可修改
* 指针类型 Type 转成值类型 Type。使用Elem方法
* 值类型 Type 转成指针类型 Type。PtrTo 返回的是指向 t 的指针类型 Type。
针对反射三定律的这个第三条，还需要特殊说明的是：Value 值的可修改性是什么意思。见上面demo
https://halfrost.com/go_reflection/
*/
