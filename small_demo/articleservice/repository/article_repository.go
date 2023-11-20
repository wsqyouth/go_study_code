package repository

import (
	"articleservice/entity"
	"context"
	"database/sql"
	"errors"
	"log"
	"sync"
)

type ArticleRepository interface {
	GetByID(ctx context.Context, id int64) (*entity.Article, error)
	Save(ctx context.Context, article *entity.Article) error
}

////////////////////////////////////////////////////////////////////////
type mysqlArticleRepository struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(Conn *sql.DB) ArticleRepository {
	return &mysqlArticleRepository{Conn}
}

func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int64) (*entity.Article, error) {
	// 实现从MySQL数据库获取文章的逻辑
	log.Printf("mysqlArticleRepository GetByID %+v", id)
	return nil, nil
}

func (m *mysqlArticleRepository) Save(ctx context.Context, article *entity.Article) error {
	// 实现保存文章到MySQL数据库的逻辑
	log.Printf("mysqlArticleRepository Save %+v", article)
	return nil
}

////////////////////////////////////////////////////////////////////////
type memoryArticleRepository struct {
	mu       sync.RWMutex
	articles map[int64]*entity.Article
}

func NewMemoryArticleRepository() ArticleRepository {
	return &memoryArticleRepository{
		articles: make(map[int64]*entity.Article),
	}
}

func (m *memoryArticleRepository) GetByID(ctx context.Context, id int64) (*entity.Article, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	article, ok := m.articles[id]
	if !ok {
		return nil, errors.New("article not found")
	}

	return article, nil
}

func (m *memoryArticleRepository) Save(ctx context.Context, article *entity.Article) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.articles[article.ID] = article
	return nil
}
