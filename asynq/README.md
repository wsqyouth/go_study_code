假定在上运行Redis服务器localhost:6379。在开始之前，请确保已安装并运行Redis。

### 代码实现：
producer.go 将使用四种方式创建异步处理任务
tasks.go 定义异步处理的任务结构，以及消费者要处理的Handler
consumer.go 将处理producer创建的任务


### 代码运行
go get github.com/hibiken/asynq/tools/asynqmon
watch asynq stat

cd producer  && go run .
cd consumer  && go run .