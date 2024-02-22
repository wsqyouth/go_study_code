package repo

import (
	"context"
	"ddd_app/internal/domain/entity"
)

// UserInfoRepo 获取用户信息
type UserInfoRepo interface {
	GetUserInfoByID(ctx context.Context, userID int) (*entity.UserInfo, error)
}
