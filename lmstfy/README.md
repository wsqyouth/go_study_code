这个demo体现了以下核心设计思想：

1. __队列隔离__

- 每个队列(Queue)独立运行
- 有自己的channel缓冲区

2. __Subscriber模型__

- 每个Subscriber独立工作
- 包含自己的处理channel
- 可配置订阅线程数和处理线程数

3. __多级并发处理__

- 第一级：多个subscriber线程并发消费队列
- 第二级：多个processor线程并发处理消息

4. __优雅关闭__

- 使用context控制生命周期
- 可以优雅停止所有goroutine

运行这个demo，你会看到：

1. 不同队列的消息被不同的subscriber处理
2. 每个队列有多个订阅者线程并发消费
3. 每个subscriber有多个处理线程并发处理消息
4. 消息处理的完整流程：接收->处理->ACK


