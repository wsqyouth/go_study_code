package main

import (
	"context"
	"database/sql"
	"productservice/entity"
	"productservice/service"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

func TestMysqlProductDemo(t *testing.T) {
	db, mock, err := sqlmock.New() // 创建模拟数据库连接
	assert.Nil(t, err)
	defer db.Close()
	tests := []struct {
		name           string
		patchFunc      func(*gomonkey.Patches, *service.MockProductService)
		expectedOutput string
		expectError    bool
	}{
		{
			name: "successful demo",
			patchFunc: func(patches *gomonkey.Patches, mockService *service.MockProductService) {
				expectedProduct := &entity.Product{ID: "1", Name: "Example Product"}
				patches.ApplyMethod(reflect.TypeOf(mockService), "CreateProduct", func(_ *service.MockProductService, _ context.Context, _ string) (*entity.Product, error) {
					return expectedProduct, nil
				}).ApplyMethod(reflect.TypeOf(mockService), "GetProductByID", func(_ *service.MockProductService, _ context.Context, _ string) (*entity.Product, error) {
					return expectedProduct, nil
				}).ApplyMethod(reflect.TypeOf(mockService), "GetAllProducts", func(_ *service.MockProductService, _ context.Context) ([]*entity.Product, error) {
					return []*entity.Product{expectedProduct}, nil
				})
			},
			expectedOutput: "Created product: {ID:1 Name:Example Product}\nRetrieved product: {ID:1 Name:Example Product}\nAll products: [{ID:1 Name:Example Product}]\n",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockService := new(service.MockProductService)
			dbPatches := gomonkey.NewPatches()
			defer dbPatches.Reset()
			dbPatches.ApplyFunc(sql.Open, func(driverName, dataSourceName string) (*sql.DB, error) {
				return db, nil // 返回nil表示成功打开数据库
			})
			dbPatches.ApplyFunc(service.NewProductServiceWithMySQLRepo, func(db *sql.DB) service.ProductService {
				return mockService
			})

			if tt.patchFunc != nil {
				servicePatches := gomonkey.NewPatches()
				defer servicePatches.Reset()
				tt.patchFunc(servicePatches, mockService)
			}

			// 调用MysqlProductDemo函数
			MysqlProductDemo(ctx)

			// 验证mockService的预期行为是否被调用
			mockService.AssertExpectations(t)
			mock.ExpectationsWereMet() // 确保所有模拟数据库操作的预期都已满足

		})
	}
}
