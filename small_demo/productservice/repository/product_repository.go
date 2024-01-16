package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"productservice/entity"
	"sync"
	"time"
)

// ProductRepository defines the interface for product repository.
type ProductRepository interface {
	Insert(ctx context.Context, product *entity.Product) error
	GetByID(ctx context.Context, id string) (*entity.Product, error)
	GetAll(ctx context.Context) ([]*entity.Product, error)
}

////////////////////////////////////////////////////////////////////////

// InMemoryProductRepository is an in-memory implementation of ProductRepository.
type InMemoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*entity.Product
}

// NewInMemoryProductRepository returns a new instance of InMemoryProductRepository.
func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*entity.Product),
	}
}

// Insert inserts a new product into the repository.
func (r *InMemoryProductRepository) Insert(ctx context.Context, product *entity.Product) error {
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
func (r *InMemoryProductRepository) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if product, exists := r.products[id]; exists {
		return product, nil
	}
	return nil, errors.New("product not found")
}

// GetAll returns all products in the repository.
func (r *InMemoryProductRepository) GetAll(ctx context.Context) ([]*entity.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	products := make([]*entity.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

////////////////////////////////////////////////////////////////////////

// MySQLProductRepository is a MySQL implementation of ProductRepository.
type MySQLProductRepository struct {
	db *sql.DB
}

// NewMySQLProductRepository returns a new instance of MySQLProductRepository.
func NewMySQLProductRepository(db *sql.DB) *MySQLProductRepository {
	return &MySQLProductRepository{
		db: db,
	}
}

// Insert inserts a new product into the MySQL database.
func (r *MySQLProductRepository) Insert(ctx context.Context, product *entity.Product) error {
	query := "INSERT INTO products (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, product.ID, product.Name, product.CreatedAt, product.UpdatedAt)
	return err
}

// GetByID returns a product by its ID from the MySQL database.
func (r *MySQLProductRepository) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	query := "SELECT id, name, created_at, updated_at FROM products WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var product entity.Product
	err := row.Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// GetAll returns all products from the MySQL database.
func (r *MySQLProductRepository) GetAll(ctx context.Context) ([]*entity.Product, error) {
	query := "SELECT id, name, created_at, updated_at FROM products"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

////////////////////////////////////////////////////////////////////////
