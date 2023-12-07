package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserService struct {
	db    *sql.DB
	cache *redis.Client
}

func NewUserService(db *sql.DB, cache *redis.Client) *UserService {
	return &UserService{
		db:    db,
		cache: cache,
	}
}

func (s *UserService) UpdateUser(ctx context.Context, user User) error {
	// Update database
	_, err := s.db.Exec("UPDATE user SET name = ?, age = ? WHERE id = ?", user.Name, user.Age, user.ID)
	if err != nil {
		return err
	}

	// Check cache 测试下cache是否存在这个user
	userKey := fmt.Sprintf("user:%d", user.ID)
	val, err := s.cache.Get(ctx, userKey).Result()
	if err == redis.Nil {
		fmt.Println("User does not exist in cache")
	} else if err != nil {
		return err
	} else {
		fmt.Println("User exists in cache:", val)
	}
	// Delete cache
	fmt.Println("DELETE FROM cache key: ", userKey)
	err = s.cache.Del(ctx, userKey).Err()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	db, err := sql.Open("mysql", "coopers:2019Youth@tcp(127.0.0.1:3306)/sql_test")
	if err != nil {
		panic(err)
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	service := NewUserService(db, cache)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			user := User{
				// ID:   id,
				ID:   1, // All goroutines update the same user
				Name: fmt.Sprintf("User%d", id),
				Age:  id + 20,
			}

			err := service.UpdateUser(ctx, user)
			if err != nil {
				fmt.Printf("Error updating user %d: %v\n", id, err)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All updates finished.")
}
