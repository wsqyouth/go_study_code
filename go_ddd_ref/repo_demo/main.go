package main

import (
	"context"
	"fmt"
	"repo_demo/repository"
)

func main() {
	ctx := context.Background()
	user, err := repository.NewUserRepo().GetUserByID(ctx, 1)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("user:", user)
}
