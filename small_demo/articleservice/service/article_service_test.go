package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"articleservice/entity"
	"articleservice/repository"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
)

func TestGetArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockArticleRepository()
	svc := NewArticleService(repo)

	article := &entity.Article{ID: 1, Title: "Test", Content: "Test content"}

	gomonkey.ApplyMethod(reflect.TypeOf(repo), "GetByID", func(_ *mockArticleRepository, _ context.Context, _ int64) (*entity.Article, error) {
		return article, nil
	})

	got, err := svc.GetArticle(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(got, article)
	if got.ID != article.ID || got.Title != article.Title || got.Content != article.Content {
		t.Fatalf("got %v, want %v", got, article)
	}
}

type mockArticleRepository struct {
}

func NewMockArticleRepository() repository.ArticleRepository {
	return &mockArticleRepository{}
}
func (m *mockArticleRepository) GetByID(ctx context.Context, id int64) (*entity.Article, error) {
	return nil, nil
}

func (m *mockArticleRepository) Save(ctx context.Context, article *entity.Article) error {
	return nil
}

/*
go test  -gcflags=-l -run TestGetArticle
使用gomonkey进行反射
*/
