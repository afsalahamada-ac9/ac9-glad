/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package live_darshan

import (
	"ac9/glad/pkg/id"
	"ac9/glad/services/ldsd/entity"
)

type Writer interface {
	Create(*entity.LiveDarshan) error
	Delete(ldID id.ID) error
	Update(*entity.LiveDarshan) error
}

type Reader interface {
	Get(ldID id.ID) (*entity.LiveDarshan, error)
	List(tenantID id.ID, page, limit int) ([]*entity.LiveDarshan, error)
	GetCount(tenantID id.ID) (int, error)
}

type Repository interface {
	Writer
	Reader
}

type UseCase interface {
	CreateLiveDarshan(
		tenantID id.ID,
		date string,
		startTime string,
		meetingURL string,
		createdBy id.ID,
	) (*entity.LiveDarshan, error)
	GetLiveDarshan(ldID id.ID) (*entity.LiveDarshan, error)
	ListLiveDarshan(tenantID id.ID, page, limit int) ([]*entity.LiveDarshan, error)
	UpdateLiveDarshan(e *entity.LiveDarshan) error
	DeleteLiveDarshan(ldID id.ID) error
	GetCount(tenantID id.ID) int
}
