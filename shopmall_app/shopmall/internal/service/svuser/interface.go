package svuser

import (
	"context"
	v1 "shopmall/api/v1"
)

// UserService 定义用户服务接口
type UserService interface {
	CreateUser(ctx context.Context, req *v1.UserCreateRequest) (*v1.UserResponse, error)
	GetUser(ctx context.Context, id uint64) (*v1.UserResponse, error)
	GetUserByUsername(ctx context.Context, username string) (*v1.UserResponse, error)
	UpdateUser(ctx context.Context, id uint64, req *v1.UserUpdateRequest) (*v1.UserResponse, error)
	ListUsers(ctx context.Context, page, pageSize int) (*v1.UserListResponse, error)
	DeleteUser(ctx context.Context, id uint64) error
}

// 确保 UserServiceImpl 实现了 UserService 接口
var _ UserService = (*UserServiceImpl)(nil)
