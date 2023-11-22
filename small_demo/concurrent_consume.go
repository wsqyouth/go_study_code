package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

type ResultStultents struct {
	UID  int
	Data string
}

func processRecord(record *ResultStultents) {
	// 模拟处理记录的操作
	fmt.Printf("Processing record with UID %d\n", record.UID)
	// 这里可以加入实际的处理逻辑
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func main() {
	// 创建通道切片
	numChannels := 128
	chs := make([]chan *ResultStultents, numChannels)

	// 初始化通道
	for i := range chs {
		chs[i] = make(chan *ResultStultents, 1024)
	}
	fmt.Println("chs len", len(chs))
	// 模拟从数据库中检索的结果
	pulsarMessageList := make([]*ResultStultents, 1000)
	for i := 0; i < len(pulsarMessageList); i++ {
		pulsarMessageList[i] = &ResultStultents{
			UID:  i,
			Data: fmt.Sprintf("data %d", i),
		}
	}

	// 将结果分发到对应的通道
	for _, each := range pulsarMessageList {
		chanIndex := each.UID % len(chs)
		ch := chs[chanIndex]
		ch <- each
	}

	// 使用errgroup进行并发消费
	var g errgroup.Group
	for _, ch := range chs {
		ch := ch // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			for msg := range ch {
				processRecord(msg)
			}
			return nil
		})
	}

	// 关闭所有通道
	for _, ch := range chs {
		close(ch)
	}

	// 等待所有goroutine结束
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All tasks done!")
	}
}

/*
在这个示例中，我们首先创建了一个包含128个channel的切片，并初始化了每个channel。然后，我们创建了一个包含1000个ResultStultents对象的切片，并将每个对象发送到一个由其UID值决定的channel。

然后，我们使用errgroup库创建了一个新的errgroup.Group对象，并为每个channel启动了一个goroutine来从channel中读取并处理数据。每个goroutine都会在其对应的channel被关闭并且所有数据都被处理完毕后结束。

最后，我们关闭了所有的channel，并使用errgroup.Group.Wait方法等待所有goroutine结束。如果任何goroutine在处理数据时返回了错误，Wait方法就会返回这个错误。

*/
