package rpuser

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserMemRepoImpl struct {
	sync.RWMutex
	users map[uint64]*User
	idSeq uint64
}

func NewUserMemRepoImpl() *UserMemRepoImpl {
	return &UserMemRepoImpl{
		users: make(map[uint64]*User),
	}
}

func (r *UserMemRepoImpl) Create(ctx context.Context, user *User) error {
	r.Lock()
	defer r.Unlock()

	user.ID = atomic.AddUint64(&r.idSeq, 1)
	r.users[user.ID] = user
	return nil
}

func (r *UserMemRepoImpl) GetByID(ctx context.Context, id uint64) (*User, error) {
	r.RLock()
	defer r.RUnlock()

	if user, ok := r.users[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found: %d", id)
}

func (r *UserMemRepoImpl) GetByUsername(ctx context.Context, username string) (*User, error) {
	r.RLock()
	defer r.RUnlock()
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found: %s", username)
}

func (r *UserMemRepoImpl) Update(ctx context.Context, user *User) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return fmt.Errorf("user not found: %d", user.ID)
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserMemRepoImpl) List(ctx context.Context, offset, limit int) (int64, []*User, error) {
	r.RLock()
	defer r.RUnlock()

	total := int64(len(r.users))
	if offset >= len(r.users) {
		return total, []*User{}, nil
	}

	users := make([]*User, 0, limit)
	count := 0
	for _, user := range r.users {
		if count >= offset && len(users) < limit {
			users = append(users, user)
		}
		count++
	}

	return total, users, nil
}

func (r *UserMemRepoImpl) Delete(ctx context.Context, id uint64) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.users[id]; !ok {
		return fmt.Errorf("user not found: %d", id)
	}

	delete(r.users, id)
	return nil
}
