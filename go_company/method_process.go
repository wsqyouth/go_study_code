package main

import (
    "context"
    "fmt"
)

type Obj struct {
    a int
    b int
    sum int 
    sub int
}

func main() {
    var obj Obj
    obj = Obj{
        a:100,
        b:20,
    }
    err := Process(context.Background(),&obj)
    if err != nil {
        fmt.Println(err)
        panic("Process error")
    }
    fmt.Println("hello",obj)
}

// Process 针对一个对象呃逆多个成员组装生成其他成员的批量处理方法
func Process(ctx context.Context, obj *Obj) (err error){
    type setField func(ctx context.Context, obj *Obj) (err error)
    funcs := []setField{
        SetSum,
        SetSub,
    }

    for _,each := range funcs {
        err = each(ctx,obj)
        if err != nil {
            return err
        }
    }
    return nil
}


func SetSum(ctx context.Context, obj *Obj) (err error){
    obj.sum = obj.a + obj.b
    return nil
}
func SetSub(ctx context.Context, obj *Obj) (err error){
    obj.sub = obj.a - obj.b
    return nil
}
