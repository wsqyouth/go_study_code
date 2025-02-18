package rpuser

import (
	"context"

	"gorm.io/gorm"
)

// UserDBRepoImpl 实现 Repository 接口
type UserDBRepoImpl struct {
	db *gorm.DB
}

// NewUserDBRepoImpl 创建仓储实例
func NewUserDBRepoImpl(db *gorm.DB) Repository {
	return &UserDBRepoImpl{db: db}
}

func (r *UserDBRepoImpl) Create(ctx context.Context, user *User) error {
	return r.db.Create(user).Error
}

func (r *UserDBRepoImpl) GetByID(ctx context.Context, id uint64) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserDBRepoImpl) GetByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDBRepoImpl) Update(ctx context.Context, user *User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserDBRepoImpl) List(ctx context.Context, offset, limit int) (total int64, users []*User, err error) {
	var count int64
	if err = r.db.WithContext(ctx).Model(&User{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}

	var userList []*User
	if err = r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&userList).Error; err != nil {
		return 0, nil, err
	}

	return count, userList, nil
}

func (r *UserDBRepoImpl) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&User{}, id).Error
}
