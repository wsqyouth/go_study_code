package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
	"github.com/yourusername/yourproject/entity"
)

func TestGetArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockArticleRepository(ctrl)
	svc := service.NewArticleService(repo)

	article := &entity.Article{ID: 1, Title: "Test", Content: "Test content"}

	gomonkey.ApplyMethod(reflect.TypeOf(repo), "GetByID", func(_ *mock.MockArticleRepository, _ context.Context, _ int64) (*entity.Article, error) {
		return article, nil
	})

	got, err := svc.GetArticle(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != article.ID || got.Title != article.Title || got.Content != article.Content {
		t.Fatalf("got %v, want %v", got, article)
	}
}
