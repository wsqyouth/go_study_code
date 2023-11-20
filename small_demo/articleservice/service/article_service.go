package service

import (
	"articleservice/entity"
	"articleservice/repository"
	"context"
)

type ArticleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) GetArticle(ctx context.Context, id int64) (*entity.Article, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ArticleService) CreateArticle(ctx context.Context, article *entity.Article) error {
	return s.repo.Save(ctx, article)
}
