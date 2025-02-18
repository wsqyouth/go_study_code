package v1

// UserCreateRequest 定义创建用户的请求结构
type UserCreateRequest struct {
	Username string `json:"username" binding:"required" example:"zhangsan"`
	Email    string `json:"email" binding:"required,email" example:"zhangsan@example.com"`
}

// UserResponse 定义用户响应结构
type UserResponse struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type UserUpdateRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserListResponse struct {
	Total int64          `json:"total"`
	Users []UserResponse `json:"users"`
}
