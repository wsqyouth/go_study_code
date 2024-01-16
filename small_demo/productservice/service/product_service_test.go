package service

import (
	"context"
	"productservice/entity"
	"productservice/repository"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService_CreateProduct(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	ctx := context.Background()
	productName := "Test Product"
	expectedProduct := &entity.Product{
		ID:   "some-unique-id",
		Name: productName,
	}

	// 使用gomonkey打桩generateID函数
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	patches.ApplyFunc(generateID, func() string {
		return "some-unique-id"
	})

	mockRepo.On("Insert", ctx, mock.AnythingOfType("*entity.Product")).Return(nil)

	product, err := service.CreateProduct(ctx, productName)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductByID(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	ctx := context.Background()
	productID := "some-unique-id"
	expectedProduct := &entity.Product{
		ID:   productID,
		Name: "Test Product",
	}

	mockRepo.On("GetByID", ctx, productID).Return(expectedProduct, nil)

	product, err := service.GetProductByID(ctx, productID)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetAllProducts(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	ctx := context.Background()
	expectedProducts := []*entity.Product{
		{ID: "id1", Name: "Test Product 1"},
		{ID: "id2", Name: "Test Product 2"},
	}

	mockRepo.On("GetAll", ctx).Return(expectedProducts, nil)

	products, err := service.GetAllProducts(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	mockRepo.AssertExpectations(t)
}
