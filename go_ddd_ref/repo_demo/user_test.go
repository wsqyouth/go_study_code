// user_test.go

package repository_test

import (
	"context"
	"reflect"
	"testing"

	"repo_demo/repository"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
)

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockIUserRepo(ctrl)
	userID := int64(1)
	expectedUser := &repository.User{ID: userID, Name: "John Doe"}

	// Using gomock
	mockUserRepo.EXPECT().GetUserByID(gomock.Any(), userID).Return(expectedUser, nil)

	userRepo := repository.NewUserRepo()
	patches := gomonkey.ApplyMethod(reflect.TypeOf(userRepo), "GetUserByID", func(_ *repository.userRepoImpl, ctx context.Context, id int64) (*repository.User, error) {
		return mockUserRepo.GetUserByID(ctx, id)
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
