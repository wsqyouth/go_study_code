package rpuser

import "context"

// Repository 定义用户仓储接口
type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint64) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Update(ctx context.Context, user *User) error
	List(ctx context.Context, offset, limit int) (total int64, users []*User, err error)
	Delete(ctx context.Context, id uint64) error
}

// 确保 UserMemRepoImpl 实现了 Repository 接口
var _ Repository = (*UserMemRepoImpl)(nil)
