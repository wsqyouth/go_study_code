package repo

import (
	"context"
	"ddd_app/internal/domain/entity"
)

type ArticleRepo interface {
	Save(ctx context.Context, article entity.Article) error
	GetArticlesByUserID(ctx context.Context, userID int) ([]entity.Article, error)
}