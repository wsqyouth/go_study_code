// product_repository.go

package repository

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Product represents the domain model for a product.
type Product struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ProductRepository defines the interface for product repository.
type ProductRepository interface {
	Insert(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context) ([]*Product, error)
}

// NewProductRepository returns a new UserMessageRepo instance.
func NewProductRepository() ProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*Product),
	}
}

////////////////////////////////////////////////////////////////////////

// InMemoryProductRepository is an in-memory implementation of ProductRepository.
type InMemoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*Product
}

// NewInMemoryProductRepository returns a new instance of InMemoryProductRepository.
func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*Product),
	}
}

// Insert inserts a new product into the repository.
func (r *InMemoryProductRepository) Insert(ctx context.Context, product *Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.products[product.ID]; exists {
		return errors.New("product already exists")
	}
	product.CreatedAt = time.Now()
	product.UpdatedAt = product.CreatedAt
	r.products[product.ID] = product
	return nil
}

// GetByID returns a product by its ID.
func (r *InMemoryProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if product, exists := r.products[id]; exists {
		return product, nil
	}
	return nil, errors.New("product not found")
}

// GetAll returns all products in the repository.
func (r *InMemoryProductRepository) GetAll(ctx context.Context) ([]*Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	products := make([]*Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

////////////////////////////////////////////////////////////////////////
