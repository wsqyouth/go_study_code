// mock_product_service.go

package service

import (
	"context"
	"productservice/entity"

	"github.com/stretchr/testify/mock"
)

// MockProductService 是 ProductService 接口的一个 mock 实现
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) CreateProduct(ctx context.Context, name string) (*entity.Product, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *MockProductService) GetProductByID(ctx context.Context, id string) (*entity.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *MockProductService) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Product), args.Error(1)
}
