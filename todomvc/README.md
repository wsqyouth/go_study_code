# Go TodoMVC

Go TodoMVC backend is based on [gin-gonic](https://github.com/gin-gonic/gin) and [go-gorm](https://github.com/go-gorm/gorm) development, and the frontend is based on [TodoMVC/VanillaJS](https://todomvc.com/examples/vanillajs/) development.


## develop

Dependencies need to be installed before development.

```
make install
```

Run it.
```bash
make dev
```

## build
```
make build
```


----
这里看到`https://ide.cloud.tencent.com/tty/icqhnl/`内部的项目前后端结合的不错，就搬过来方便后续学习下。

环境命令：
go env -w  GO111MODULE=off
go env -w  GO111MODULE=auto  //因为我这里使用了多module,仅在当前命令行运行验证
go mod tidy