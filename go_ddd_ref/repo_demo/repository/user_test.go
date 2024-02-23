// user_test.go

package repository_test

import (
	"context"
	"reflect"
	"testing"

	"repo_demo/repository"
	mock_repository "repo_demo/repository/mock"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
)

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockIUserRepo(ctrl)
	userID := int64(1)
	expectedUser := &repository.User{ID: userID, Name: "John Doe"}

	userRepo := repository.NewUserRepo()
	patches := gomonkey.ApplyFunc(repository.NewUserRepo, func() repository.IUserRepo {
		return mockUserRepo
	}).ApplyMethod(reflect.TypeOf(mockUserRepo), "GetUserByID", func(_ *mock_repository.MockIUserRepo, ctx context.Context, id int64) (*repository.User, error) {
		return &repository.User{ID: userID, Name: "John Doe"}, nil
	})
	defer patches.Reset()

	user, err := userRepo.GetUserByID(context.Background(), userID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.ID != expectedUser.ID || user.Name != expectedUser.Name {
		t.Errorf("expected user %+v, got %+v", expectedUser, user)
	}
}
