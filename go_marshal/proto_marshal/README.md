1. 根据proto文件生成pb文件
protoc --go_out=. --go_opt=paths=source_relative my_message.proto
2. main代码进行验证
go run .  