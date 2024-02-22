package persistence

import (
	"context"
	"ddd_app/internal/domain/entity"
	"ddd_app/internal/domain/repo"
	"errors"
	"fmt"
)

// UserRepoImpl user persistence impl
type UserRepoImpl struct {
	// 依赖的db句柄,或者其他资源
	UserInfoConf map[int]entity.UserInfo
}

func NewUserInfoRepo() repo.UserInfoRepo {
	// 这里暂时使用一份固定的人员配置,即要求userId必须在这个map里面
	mockUserInfoConf := make(map[int]entity.UserInfo)
	mockUserInfoConf[1] = entity.UserInfo{ID: 1, Name: "adamin"}
	return &UserRepoImpl{
		UserInfoConf: mockUserInfoConf,
	}
}

// GetUserInfo 通过userID获取用户基本信息
func (u *UserRepoImpl) GetUserInfoByID(ctx context.Context, userID int) (*entity.UserInfo, error) {
	if userID == 0 {
		return nil, ErrUserNotFound
	}
	// 检查该userID是否在UserInfoConf里，不在则报错
	userInfo, ok := u.UserInfoConf[userID]
	if !ok {
		return nil, ErrUserNotFound
	}
	fmt.Println("userInfo:", userInfo)
	return &userInfo, nil
}

var (
	// ErrUserNotFound user not found
	ErrUserNotFound = errors.New("user not found")
)
