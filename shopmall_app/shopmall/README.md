# shopmall

一个轻量级的 Gin Web 服务项目。

## 功能特性

- 基于 Gin 框架的 HTTP Web 服务
- 支持 RESTful API
- 简洁的项目结构
- 支持 Docker 部署

这个项目具有以下特点：

1. 清晰的分层架构：handler -> service -> repository
2. 使用 Gin 框架处理 HTTP 请求
3. 基于内存的存储实现
4. 简单的配置管理
5. 符合 DDD 设计思想
6. 易于理解和扩展

 ### 项目结构
 ```
 shopmall/
├── bin/                # 编译后的二进制文件
├── cmd/                # 主程序入口
├── internal/           # 内部包
├── pkg/                # 可重用的包
├── Dockerfile         # Docker 构建文件
├── Makefile          # 项目管理工具
└── README.md         # 项目文档
 ```

## 快速开始

### 前置要求

- Go 1.20 或更高版本
- Make

### 本地运行

1. 克隆项目
```bash
git clone <your-repository-url>
cd shopmall
```
2. 构建项目
```bash
make build
 ```

3. 运行服务
```bash
make run
 ```

 ### 测试
 注册用户：
 ```
 curl -X POST http://localhost:8000/api/v1/users -H "Content-Type: application/json" -d '{"username":"test","email":"test@example.com"}'

```
查询用户
 ```
 curl http://localhost:8000/api/v1/users/1
 ```
 更新用户：

```bash
curl -X PUT 'http://localhost:8000/api/v1/users/1' \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "newname",
    "email": "newemail@example.com"
  }'
```
获取用户列表：

```bash
curl -X GET 'http://localhost:8000/api/v1/users?page=1&page_size=10'
```

删除用户：
```bash
curl -X DELETE 'http://localhost:8000/api/v1/users/1'
```
