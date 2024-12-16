/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package center

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"
)

// Reader interface
type Reader interface {
	Get(id id.ID) (*entity.Center, error)
	Search(tenantID id.ID, query string, page, limit int) ([]*entity.Center, error)
	List(tenantID id.ID, page, limit int) ([]*entity.Center, error)
	GetCount(id id.ID) (int, error)
}

// Writer center writer
type Writer interface {
	Create(e *entity.Center) (id.ID, error)
	Update(e *entity.Center) error
	Delete(id id.ID) error
	Upsert(e *entity.Center) (id.ID, error)
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetCenter(id id.ID) (*entity.Center, error)
	SearchCenters(tenantID id.ID, query string, page, limit int) ([]*entity.Center, error)
	ListCenters(tenantID id.ID, page, limit int) ([]*entity.Center, error)
	CreateCenter(tenantID id.ID, name string, mode entity.CenterMode, isEnabled bool) (id.ID, error)
	UpdateCenter(e *entity.Center) error
	DeleteCenter(id id.ID) error
	GetCount(id id.ID) int
	UpsertCenter(e *entity.Center) (id.ID, error)
}
