package persistence

import (
	"fmt"
	"context"
	"ddd_app/internal/domain/entity"
	"ddd_app/internal/domain/repo"
)

type ArticleMapRepo struct {
	storage map[int][]entity.Article
}

func NewArticleMapRepo() repo.ArticleRepo {
	return &ArticleMapRepo{
		storage: make(map[int][]entity.Article),
	}
}

func (r *ArticleMapRepo) Save(ctx context.Context, article entity.Article) error {
	r.storage[article.UserID] = append(r.storage[article.UserID], article)
	fmt.Println("map save:",r.storage)
	return nil
}

func (r *ArticleMapRepo) GetArticlesByUserID(ctx context.Context, userID int) ([]entity.Article, error) {
	fmt.Println("map has:",r.storage)
	if articles, ok := r.storage[userID]; ok {
		return articles, nil
	}
	return nil, nil
}