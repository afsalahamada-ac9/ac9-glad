/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package metadata

import (
	entity "ac9/glad/services/mediad/entity"
)

// Writer interface
type Writer interface {
	CreateQuote(e *entity.Quote) error
	CreateMedia(e *entity.Media) error
}

// Repository interface
type Repository interface {
	Writer
}

// UseCase interface
type UseCase interface {
	CreateMetadata(version int64, url string, total string, metadataType string) error
}
