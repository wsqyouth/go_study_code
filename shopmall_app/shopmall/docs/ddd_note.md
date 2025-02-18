# DDD 架构设计笔记

## 项目结构

internal/
├── entity/          # 领域实体
├── service/         # 业务逻辑
├── repository/      # 数据访问
└── handler/         # HTTP 处理

## 分层职责

### 1. Entity 层（领域实体层）
- 定义核心业务实体和值对象
- 包含业务规则和验证逻辑
- 不依赖其他层，保持独立性
- 示例：用户实体、订单实体、商品实体等

```go
// entity/user.go
type User struct {
    ID       uint64
    Username string
    Email    string
}

func (u *User) Validate() error {
    if u.Username == "" {
        return errors.New("username cannot be empty")
    }
    return nil
}
```

### 2. Repository 层（数据访问层）
- 定义数据访问接口
- 实现数据持久化逻辑
- 处理数据库操作和查询
- 不包含业务逻辑

```go
// repository/user.go
type UserRepository interface {
    Create(ctx context.Context, user *entity.User) error
    GetByID(ctx context.Context, id uint64) (*entity.User, error)
    Update(ctx context.Context, user *entity.User) error
}
```

### 3. Service 层（业务逻辑层）
- 协调领域实体和仓储层
- 实现核心业务逻辑
- 处理事务和数据转换
- 依赖 Entity 和 Repository 层

```go
// service/user.go
type UserService struct {
    repo repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    return s.repo.Create(ctx, user)
}
```

### 4. Handler 层（HTTP 处理层）
- 处理 HTTP 请求和响应
- 参数验证和转换
- 调用 Service 层处理业务逻辑
- 处理错误和返回结果

```go
// handler/user.go
type UserHandler struct {
    userService *service.UserService
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := h.userService.CreateUser(c, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, user)
}
```

## 依赖关系

依赖方向：handler -> service -> (entity, repository)

- Handler 层依赖 Service 层
- Service 层依赖 Entity 和 Repository 层
- Entity 层不依赖任何其他层
- Repository 层仅依赖 Entity 层

## 最佳实践

1. **接口隔离**
   - 每层都定义清晰的接口
   - 依赖接口而不是具体实现
   - 便于单元测试和模拟

2. **依赖注入**
   - 使用依赖注入管理层之间的依赖
   - 避免直接实例化依赖对象
   - 提高代码的可测试性和灵活性

3. **错误处理**
   - 定义领域特定的错误类型
   - 在适当的层处理错误
   - 提供清晰的错误信息

4. **数据验证**
   - 在 Entity 层进行业务规则验证
   - 在 Handler 层进行请求参数验证
   - 避免重复验证

## 扩展性

这种架构设计便于后续添加新的服务模块，如：
- etorder（电子订单服务）
- svorder（服务订单系统）
- rporder（零售订单平台）

每个新服务都可以遵循相同的分层结构，保持一致性和可维护性。

## 常见问题解决方案

1. **循环依赖**
   - 使用接口打破循环依赖
   - 重新思考业务边界
   - 考虑提取共享服务

2. **业务逻辑位置**
   - 简单验证放在 Entity 层
   - 复杂业务逻辑放在 Service 层
   - 避免在 Repository 层放置业务逻辑

3. **数据一致性**
   - 使用事务管理复杂操作
   - 定义清晰的数据边界
   - 合理使用乐观锁和悲观锁