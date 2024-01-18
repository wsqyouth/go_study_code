// main.go

package main

import (
	"context"
	"fmt"
	"mockgen_demo/repository"
)

func main() {
	ctx := context.Background()

	// NewProductRepository 返回了一个实现了 ProductRepository 接口的实例
	productRepo := repository.NewProductRepository()
	MysqlProductDemo(ctx, productRepo)
}

// MysqlProductDemo 使用 productRepo 来执行一些操作
func MysqlProductDemo(ctx context.Context, productRepo repository.ProductRepository) {
	product := &repository.Product{
		ID:   "1",
		Name: "Example Product",
	}
	err := productRepo.Insert(ctx, product)
	if err != nil {
		fmt.Printf("Error inserting product: %v\n", err)
		return
	}

	retrievedProduct, err := productRepo.GetByID(ctx, "1")
	if err != nil {
		fmt.Printf("Error retrieving product: %v\n", err)
		return
	}
	fmt.Printf("Retrieved product: %+v\n", retrievedProduct)

	allProducts, err := productRepo.GetAll(ctx)
	if err != nil {
		fmt.Printf("Error retrieving all products: %v\n", err)
		return
	}
	fmt.Printf("All products: %+v\n", allProducts)
}
