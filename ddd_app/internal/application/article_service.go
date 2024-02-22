// ddd_app/internal/application/article_service.go
package application

import (
	"context"
	"ddd_app/internal/domain/entity"
	"ddd_app/internal/domain/repo"
)

type ArticleService struct {
	articleRepo repo.ArticleRepo
}

func NewArticleService(r repo.ArticleRepo) *ArticleService {
	return &ArticleService{articleRepo: r}
}

func (s *ArticleService) WriteArticle(ctx context.Context, article entity.Article) error {
	return s.articleRepo.Save(ctx, article)
}

func (s *ArticleService) GetArticlesByUserID(ctx context.Context, userID int) ([]entity.Article, error) {
	return s.articleRepo.GetArticlesByUserID(ctx, userID)
}