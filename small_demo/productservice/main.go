package main

import (
	"context"
	"database/sql"
	"fmt"
	"productservice/repository"
	"productservice/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	// InMemoryProductDemo(ctx)
	MysqlProductDemo(ctx)
}

func MysqlProductDemo(ctx context.Context) {
	// Create a MySQL database connection.
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		return
	}
	defer db.Close()

	// Initialize the service with the MySQL repository.
	productService := service.NewProductServiceWithMySQLRepo(db)

	// Use the service to create a product.
	product, err := productService.CreateProduct(ctx, "Example Product")
	if err != nil {
		fmt.Printf("Error creating product: %v\n", err)
		return
	}
	fmt.Printf("Created product: %+v\n", product)

	// Retrieve the product by ID.
	retrievedProduct, err := productService.GetProductByID(ctx, product.ID)
	if err != nil {
		fmt.Printf("Error retrieving product: %v\n", err)
		return
	}
	fmt.Printf("Retrieved product: %+v\n", retrievedProduct)

	// Retrieve all products.
	allProducts, err := productService.GetAllProducts(ctx)
	if err != nil {
		fmt.Printf("Error retrieving all products: %v\n", err)
		return
	}
	fmt.Printf("All products: %+v\n", allProducts)
}

func InMemoryProductDemo(ctx context.Context) {

	// Initialize the repository
	repo := repository.NewInMemoryProductRepository()

	// Initialize the service with the repository
	productService := service.NewProductService(repo)

	// Use the service to create a product
	product, err := productService.CreateProduct(ctx, "Example Product")
	if err != nil {
		fmt.Printf("Error creating product: %v\n", err)
		return
	}
	fmt.Printf("Created product: %+v\n", product)

	// Retrieve the product by ID
	retrievedProduct, err := productService.GetProductByID(ctx, product.ID)
	if err != nil {
		fmt.Printf("Error retrieving product: %v\n", err)
		return
	}
	fmt.Printf("Retrieved product: %+v\n", retrievedProduct)

	// Retrieve all products
	allProducts, err := productService.GetAllProducts(ctx)
	if err != nil {
		fmt.Printf("Error retrieving all products: %v\n", err)
		return
	}
	fmt.Printf("All products len: %+v\n", len(allProducts))
}
