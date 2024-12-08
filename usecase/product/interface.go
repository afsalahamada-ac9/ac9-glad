/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package product

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"
)

// Reader defines read-only operations for products
type Reader interface {
	Get(id id.ID) (*entity.Product, error)
	List(tenantID id.ID, page, limit int) ([]*entity.Product, error)
	Search(tenantID id.ID, q string, page, limit int) ([]*entity.Product, error)
	GetCount(tenantID id.ID) (int, error)
}

// Writer defines write-only operations for products
type Writer interface {
	Create(product *entity.Product) (id.ID, error)
	Update(product *entity.Product) error
	Delete(id id.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase defines the interface for product business logic
type UseCase interface {
	GetProduct(id id.ID) (*entity.Product, error)
	SearchProducts(tenantID id.ID, q string, page, limit int) ([]*entity.Product, error)
	ListProducts(tenantID id.ID, page, limit int) ([]*entity.Product, error)
	CreateProduct(tenantID id.ID,
		extID string,
		extName string,
		title string,
		ctype string,
		baseProductExtID string,
		durationDays int32,
		visibility entity.ProductVisibility,
		maxAttendees int32,
		format entity.ProductFormat,
		isAutoApprove bool,
	) (id.ID, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id id.ID) error
	GetCount(id id.ID) int
}
