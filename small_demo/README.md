### 针对small_demo项目,存放一些笔记和附图。

代码`concurrent_consume`,使用了类似下图的思想: 将相同topic的消息放入同一个队列中,然后使用多个线程消费队列中的消息。
![Alt text](image.png)