package svuser

import (
	"context"
	v1 "shopmall/api/v1"
	"shopmall/internal/entity/etuser"
	"shopmall/internal/repository/rpuser"
)

// UserServiceImpl 实现 Service 接口
type UserServiceImpl struct {
	userRepo rpuser.Repository
}

func NewUserService(repo rpuser.Repository) UserService {
	return &UserServiceImpl{userRepo: repo}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, req *v1.UserCreateRequest) (*v1.UserResponse, error) {
	// 创建领域实体
	user := etuser.NewUserEntity(req.Username, req.Email)

	// 验证领域规则
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// 转换为仓储实体并保存
	repoUser := &rpuser.User{
		Username: user.Username,
		Email:    user.Email,
	}
	if err := s.userRepo.Create(ctx, repoUser); err != nil {
		return nil, err
	}

	// 更新领域实体ID
	user.ID = repoUser.ID

	// 转换为 API 响应
	return convertToUserResponse(user), nil
}

func (s *UserServiceImpl) GetUser(ctx context.Context, id uint64) (*v1.UserResponse, error) {
	// 从仓储层获取数据
	repoUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 转换为领域实体
	user := etuser.NewUserEntity(repoUser.Username, repoUser.Email)
	user.ID = repoUser.ID

	return convertToUserResponse(user), nil
}

func (s *UserServiceImpl) GetUserByUsername(ctx context.Context, username string) (*v1.UserResponse, error) {
	// 从仓储层获取数据
	repoUser, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	// 转换为领域实体
	user := etuser.NewUserEntity(repoUser.Username, repoUser.Email)
	user.ID = repoUser.ID
	return convertToUserResponse(user), nil
}

// convertToUserResponse 将实体转换为 API 响应
func convertToUserResponse(userEntity *etuser.UserEntity) *v1.UserResponse {
	return &v1.UserResponse{
		ID:       userEntity.ID,
		Username: userEntity.Username,
		Email:    userEntity.Email,
	}
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, id uint64, req *v1.UserUpdateRequest) (*v1.UserResponse, error) {
	// 先获取现有用户
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 更新字段（只更新非空字段）
	if req.Username != "" {
		existingUser.Username = req.Username
	}
	if req.Email != "" {
		existingUser.Email = req.Email
	}

	// 保存更新
	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		return nil, err
	}

	// 转换为领域实体
	user := etuser.NewUserEntity(existingUser.Username, existingUser.Email)
	user.ID = existingUser.ID

	return convertToUserResponse(user), nil
}

func (s *UserServiceImpl) ListUsers(ctx context.Context, page, pageSize int) (*v1.UserListResponse, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 获取用户列表
	total, users, err := s.userRepo.List(ctx, offset, pageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	userResponses := make([]v1.UserResponse, 0, len(users))
	for _, u := range users {
		user := etuser.NewUserEntity(u.Username, u.Email)
		user.ID = u.ID
		userResponses = append(userResponses, *convertToUserResponse(user))
	}

	return &v1.UserListResponse{
		Total: total,
		Users: userResponses,
	}, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, id uint64) error {
	return s.userRepo.Delete(ctx, id)
}
