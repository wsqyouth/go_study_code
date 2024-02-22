// ddd_app/internal/application/blog_service.go
package application

import (
	"context"
	"ddd_app/internal/domain/entity"
	"ddd_app/internal/domain/repo"
	"ddd_app/internal/infrastructure/persistence"
	"fmt"
)

// BlogService 博客服务
type BlogService struct {
	articleRepo  repo.ArticleRepo
	userInfoRepo repo.UserInfoRepo
	// ... 服务聚合,可以添加其他的repo
}

// NewBlogService 创建博客服务实例
func NewBlogService(article repo.ArticleRepo, user repo.UserInfoRepo) *BlogService {
	return &BlogService{articleRepo: article, userInfoRepo: user}
}

// WriteArticle 发表博客
func (s *BlogService) WriteArticle(ctx context.Context, article entity.Article) (err error) {
	userID := article.UserID
	// 首先检查用户是否存在
	user, err := s.userInfoRepo.GetUserInfoByID(ctx, userID)
	if err != nil {
		if err == persistence.ErrUserNotFound {
			return fmt.Errorf("userID: %d does not exist", userID)
		}
		return err
	}
	// 如果用户存在，发表他的文章
	err = s.articleRepo.Save(ctx, article)
	if err != nil {
		return err
	}
	fmt.Printf("user:%v, article:%v", user, article)
	return nil
}

// GetArticlesByUserID 根据用户id获取博客
func (s *BlogService) GetArticlesByUserID(ctx context.Context, userID int) ([]entity.Article, error) {
	// 首先检查用户是否存在
	user, err := s.userInfoRepo.GetUserInfoByID(ctx, userID)
	if err != nil {
		if err == persistence.ErrUserNotFound {
			return nil, fmt.Errorf("userID: %d does not exist", userID)
		}
		return nil, err
	}

	// 如果用户存在，获取他的文章
	articles, err := s.articleRepo.GetArticlesByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
