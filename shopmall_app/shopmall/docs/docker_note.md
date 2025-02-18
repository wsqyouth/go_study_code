# Docker 操作笔记

## 1. 镜像管理

### 1.1 构建镜像
```bash
# 在当前目录下构建镜像
docker build -t shopmall:latest .

# 指定目录构建镜像
docker build -t shopmall:latest /Users/sq.wang/Work/GoLandProj/shopmall
```

### 1.2 查看镜像
```bash
# 查看所有本地镜像
docker images

# 查看指定镜像详细信息
docker inspect shopmall:latest
```

### 1.3 删除镜像
```bash
# 删除指定镜像
docker rmi shopmall:latest

# 强制删除镜像
docker rmi -f shopmall:latest
```

## 2. 容器管理

### 2.1 获取容器 ID
```bash
# 查看所有运行中的容器及其 ID
docker ps

# 查看所有容器（包括已停止的）及其 ID
docker ps -a

# 通过容器名称过滤查找容器 ID
docker ps -a --filter "name=shopmall"

# 只显示容器 ID
docker ps -aq

# 获取最近创建的容器 ID
docker ps -l -q
```

### 2.2 运行容器
```bash
# 后台运行容器并映射端口
docker run -d -p 8080:8000 shopmall:latest

# 运行容器并进入交互式终端
docker run -it shopmall:latest /bin/bash
```

### 2.2 容器状态查看
```bash
# 查看运行中的容器
docker ps

# 查看所有容器（包括已停止的）
docker ps -a

# 查看容器资源使用情况
docker stats
```

### 2.3 容器日志查看
```bash
# 查看容器日志
docker logs <container_id>

# 实时查看日志
docker logs -f <container_id>

# 查看最近的100行日志
docker logs --tail 100 <container_id>
```

### 2.4 容器操作
```bash
# 停止容器
docker stop <container_id>

# 启动已停止的容器
docker start <container_id>

# 重启容器
docker restart <container_id>

# 删除容器
docker rm <container_id>

# 强制删除运行中的容器
docker rm -f <container_id>
```

### 2.5 进入容器
```bash
# 进入运行中的容器
docker exec -it <container_id> /bin/bash
```

## 3. 服务验证

### 3.1 验证服务状态
```bash
# 测试 HTTP 服务
curl http://localhost:8080

# 查看端口监听状态
netstat -an | grep 8080
```

注意事项：
1. 执行 Docker 命令前，请确保 Docker Desktop 已启动并正常运行
2. 容器端口映射时，确保宿主机对应端口未被占用
3. 删除镜像前，请确保相关的容器已经停止并删除