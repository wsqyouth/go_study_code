package main

import (
    "context"
    "fmt"
    "strconv"
    "time"
)

func main() {
    ctx,cancel := context.WithTimeout(context.Background(), 30*time.Second) //30s

    //consumer
    msgList := make(chan string, 10)

    go Consumer(ctx,msgList,"A")
    go Consumer(ctx,msgList,"B")
    go Consumer(ctx,msgList,"C")
    go Consumer(ctx,msgList,"D")
    go Consumer(ctx,msgList,"E")

    c:=20
    for i :=0 ; i < c; i++{
        for {
            if len(msgList) < 10 {
                Producer(strconv.Itoa(i), msgList)
                break
            }
        }
    }

    time.Sleep(30 * time.Second)
    cancel()
    fmt.Println("done")
}

func Producer(msg string , msgList chan<- string){
    msgList<-msg
}

func Consumer(ctx context.Context, msgList<-chan string, n string){
    for {
        select {
        case msg := <- msgList:
            sendMsg(n + " == " + msg)
        case <- ctx.Done():
            return 
        }
    }
}

func sendMsg(msg string) {
    time.Sleep(time.Second * 1)
    fmt.Println(msg)
}
