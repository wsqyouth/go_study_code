# 定义应用名称变量
ARG appName=shopmall

# 构建阶段
FROM golang:1.22-alpine AS builder

WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
ARG appName
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/${appName} ./cmd/server

# 运行阶段
FROM alpine:latest

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/shopmall /usr/local/bin/

# 创建配置文件目录
RUN mkdir -p /app/configs

# 从构建阶段复制配置文件
COPY --from=builder /app/configs/config.yaml /app/configs/

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["/usr/local/bin/shopmall"]