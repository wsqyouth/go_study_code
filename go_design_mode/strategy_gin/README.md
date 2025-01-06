### 代码说明

1. **策略实现**：
    - `BasicStrategy` 和 `JWTStrategy` 实现了 `AuthStrategy` 接口，分别模拟了Basic和JWT认证策略的逻辑。
    - 在实际应用中，这里应该加入具体的token解析和验证逻辑。

2. **自动策略选择**：
    - `AutoStrategy` 包含了 `basic` 和 `jwt` 两种策略，通过 `Authorization` 头部自动选择合适的策略进行认证。

3. **HTTP服务器**：
    - 使用Go内置的 `net/http` 包创建一个简单的HTTP服务器，监听 `/auth` 路径，通过 `AutoStrategy` 处理请求。

4. **日志输出**：
    - 使用 `log.Println` 模拟输出认证策略的选择和错误处理。

该示例展示了如何独立实现一个策略模式，解析并验证不同类型的认证token，并在运行时动态选择合适的策略进行处理。


### 测试脚本
### 1. Basic Authentication 示例

假设我们使用`Basic`作为认证类型，`curl`命令如下：

```bash
curl -v -H "Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=" http://localhost:8080/auth
```


### 2. Bearer (JWT) Authentication 示例

假设我们使用`Bearer`作为认证类型，`curl`命令如下：

```bash
curl -v -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" http://localhost:8080/auth
```
在服务器端日志中，你应该看到相应的策略选择日志，并在客户端收到相应的响应消息：

- 对于Basic认证，响应消息为：`Basic Authentication Successful`
- 对于Bearer认证，响应消息为：`JWT Authentication Successful`

通过这些步骤，你可以验证两个不同的认证策略在示例代码中的运行效果。

### 备注
- 该代码是阅读 iam 项目中的中间件代码提取出来的，用于展示策略模式的使用。