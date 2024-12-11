/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package livedarshan

import (
	"ac9/glad/services/ldsd/entity"
	"time"
)

type Writer interface {
	Create(*entity.LiveDarshan) error
}

type Reader interface {
	Get(id int64) (*entity.LiveDarshan, error)
	GetAll() ([]*entity.LiveDarshan, error)
}

type Repository interface {
	Writer
	Reader
}

type Usecase interface {
	CreateLiveDarshan(id, date string, startTime time.Time, meetingUrl, createdBy string) (*entity.LiveDarshan, error)
	GetLiveDarshan(id int64) (*entity.LiveDarshan, error)
	GetAllLiveDarshan() ([]*entity.LiveDarshan, error)
}
