/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package metadata

import "ac9/glad/services/mediad/entity"

type Writer interface {
	Create(s *entity.Metadata) error
}

type Reader interface {
	Get(contentType entity.ContentType) (*entity.Metadata, error)
}

type Repository interface {
	Writer
	Reader
}

type Usecase interface {
	CreateMetadata(url string, total int, contentType string) (*entity.Metadata, error)
	GetMetadata(contentType string) (*entity.Metadata, error)
}