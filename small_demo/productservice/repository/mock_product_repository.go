package repository

import (
	"context"
	"productservice/entity"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Insert(ctx context.Context, product *entity.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *MockProductRepository) GetAll(ctx context.Context) ([]*entity.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Product), args.Error(1)
}

// 不使用 mockgen 来生成 mock，我们需要手动创建一个 mock 类型。
// 后续研究这种使用: mockgen -source=domain/product.go -destination=service/mock_product.go -package=service
