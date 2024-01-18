// main_test.go

package main

import (
	"context"
	"mockgen_demo/repository"
	mock_repository "mockgen_demo/repository/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestMysqlProductDemo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockProductRepository(ctrl)
	ctx := context.Background()
	product := &repository.Product{
		ID:        "1",
		Name:      "Example Product",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 设置预期的行为和返回值
	any := gomock.Any()
	mockRepo.EXPECT().Insert(any, any).Return(nil).Times(1)
	mockRepo.EXPECT().GetByID(ctx, "1").Return(product, nil).Times(1)
	mockRepo.EXPECT().GetAll(ctx).Return([]*repository.Product{product}, nil).Times(1)

	// 调用 MysqlProductDemo 函数
	MysqlProductDemo(ctx, mockRepo)

}

// 注: 这里使用了依赖注入,如果不能使用则只能使用patches.ApplyMethod(reflect.TypeOf(mockRepo),进行注入处理
