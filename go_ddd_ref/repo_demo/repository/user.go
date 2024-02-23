// user.go

package repository

import (
	"context"
	"errors"
)

// IUserRepo defines the interface for user repository operations
type IUserRepo interface {
	GetUserByID(ctx context.Context, userID int64) (*User, error)
}

// User represents a user in the system
type User struct {
	ID   int64
	Name string
}

// userRepoImpl is the implementation of IUserRepo
type userRepoImpl struct {
	// Here you would have your data source (e.g., database connection)
}

// NewUserRepo creates a new instance of IUserRepo
func NewUserRepo() IUserRepo {
	return &userRepoImpl{}
}

// GetUserByID retrieves a user by their ID
func (repo *userRepoImpl) GetUserByID(ctx context.Context, userID int64) (*User, error) {
	// Here you would have your actual database call
	// For this example, we'll simulate a user retrieval
	if userID == 1 {
		return &User{ID: 1, Name: "John Doe"}, nil
	}
	return nil, errors.New("user not found")
}
