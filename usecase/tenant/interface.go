/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package tenant

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"
)

// Reader interface
type Reader interface {
	Get(id id.ID) (*entity.Tenant, error)
	GetByName(username string) (*entity.Tenant, error)
	List(page, limit int) ([]*entity.Tenant, error)
	GetCount() (int, error)
}

// Writer tenant writer
type Writer interface {
	Create(e *entity.Tenant) (id.ID, error)
	Update(e *entity.Tenant) error
	Delete(id id.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetTenant(id id.ID) (*entity.Tenant, error)
	ListTenants(page, limit int) ([]*entity.Tenant, error)
	CreateTenant(username, country string) (id.ID, error)
	UpdateTenant(e *entity.Tenant) error
	DeleteTenant(id id.ID) error
	Login(username, password string) (*entity.Tenant, error)
	GetCount() int
	// Thoughts: Need to validate token; use tenant id and token to validate
}
