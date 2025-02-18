package etuser

// User 领域实体
type UserEntity struct {
	ID       uint64
	Username string
	Email    string
}

// NewUser 创建用户实体
func NewUserEntity(username, email string) *UserEntity {
	return &UserEntity{
		Username: username,
		Email:    email,
	}
}

// Validate 验证用户实体
func (u *UserEntity) Validate() error {
	// 在这里添加领域规则验证
	return nil
}
