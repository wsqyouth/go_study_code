package entity

import (
	"time"
)

// Product represents the domain model for a product.
type Product struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
