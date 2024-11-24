/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import "sudhagar/glad/entity"

// Reader defines read-only operations for products
type Reader interface {
	Get(id entity.ID) (*entity.Product, error)
	List(tenantID entity.ID, page, limit int) ([]*entity.Product, error)
	Search(tenantID entity.ID, query string) ([]*entity.Product, error)
	GetCount(tenantID entity.ID) ([]*entity.Product, error)
}

// Writer defines write-only operations for products
type Writer interface {
	Create(product *entity.Product) (entity.ID, error)
	Update(product *entity.Product) error
	Delete(id entity.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase defines the interface for product business logic
type UseCase interface {
	GetProduct(id entity.ID) (*entity.Product, error)
	SearchProducts(tenantID entity.ID, query string) ([]*entity.Product, error)
	ListProducts(tenantID entity.ID) ([]*entity.Product, error)
	CreateProduct(tenantID entity.ID,
		extID string,
		name string,
		title string,
		ctype string,
		baseProductID string,
		durationDays int32,
		visibility entity.ProductVisibility,
		maxAttendees int32,
		format entity.ProductFormat,
		isDeleted bool,
	) (entity.ID, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id entity.ID) error
	GetCount(id entity.ID) int
}
