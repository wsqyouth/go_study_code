### 一、环境及运行
go mod init articleservice

go run main.go

### 二、测试
```
设置
curl -X POST -H "Content-Type: application/json" -d '{"id":1,"title":"Test","content":"Test content"}' http://localhost:8080/articles

拉取
curl http://localhost:8080/articles/1
```

### 三、功能点总结
* 使用了gin框架
* 使用了自定义中间件添加耗时监控
* 使用了仓储模式及依赖注入
* 使用了单元测试

* 参考项目代码: https://github.com/bxcodec/go-clean-arch/blob/master/article/delivery/http/middleware/middleware.go


另外一个项目的,存放测试下
![Alt text](image.png)