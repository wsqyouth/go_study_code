### 零、环境检查：
```
chmod +x init_env.sh
./init_env.sh
```
### 一、deploy.sh 职责
只负责编译和运行deploy.go脚本
### 二、deploy.go将处理以下任务：
* 编译业务代码app.go生成app二进制文件。
* 将app二进制文件上传到远程服务器。
* 在远程服务器上设置文件权限。
* 在远程服务器上启动app。
```azure
注意: 请确保正确配置了远程服务器的SSH密钥认证，以便无密码登录和执行命令。
```
### 三、开发者使用
```
./deploy.sh 9.134.23.138 cooperswang app
```
来构建并部署应用程序。

后续优化: 目前只支持app名称，后续支持将一个项目的应用编译生成二进制文件并上传部署