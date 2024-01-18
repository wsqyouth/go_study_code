1. 记得初始化
go mod init mockgen_demo

2. 生成mock文件
go get github.com/golang/mock/
go get github.com/golang/mock/mockgen/model
mockgen -destination=repository/mock/mock_product_repository.go -package=repository mockgen_demo/repository ProductRepository

这条命令的组成部分解释如下：
-destination=repository/mock/mock_product_repository.go：指定生成的 mock 文件的路径和文件名。
-package=repository：指定生成的 mock 文件所属的包名。
mockgen_demo/repository：指定要 mock 的接口所在的包的导入路径。由于您的 go.mod 文件在项目根目录下，mockgen_demo 应该是您的模块名。
ProductRepository：指定要 mock 的接口名。
确保您在项目根目录下运行此命令，这样 mockgen 工具可以正确地找到 go.mod 文件和接口定义.