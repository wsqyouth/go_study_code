package service

import (
	"context"
	"database/sql"
	"productservice/entity"
	"productservice/repository"
)

// ProductService defines the interface for product service.
type ProductService interface {
	CreateProduct(ctx context.Context, name string) (*entity.Product, error)
	GetProductByID(ctx context.Context, id string) (*entity.Product, error)
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

////////////////////////////////////////////////////////////////////////
// NewProductService returns a new instance of ProductService.
func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

// CreateProduct creates a new product.
func (s *productService) CreateProduct(ctx context.Context, name string) (*entity.Product, error) {
	product := &entity.Product{
		ID:   generateID(), // Implement generateID to create a unique ID for the product
		Name: name,
	}
	err := s.repo.Insert(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProductByID returns a product by its ID.
func (s *productService) GetProductByID(ctx context.Context, id string) (*entity.Product, error) {
	return s.repo.GetByID(ctx, id)
}

// GetAllProducts returns all products.
func (s *productService) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	return s.repo.GetAll(ctx)
}

// generateID is a placeholder for an ID generation function.
func generateID() string {
	// Implement a proper ID generation logic
	return "some-unique-id"
}

////////////////////////////////////////////////////////////////////////

// NewProductServiceWithMySQLRepo returns a new instance of ProductService with MySQL repository.
func NewProductServiceWithMySQLRepo(db *sql.DB) ProductService {
	repo := repository.NewMySQLProductRepository(db)
	return NewProductService(repo)
}
