/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package sf_import

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
)

// Repository interface
type Repository interface {
}

// Client interface to other glad services
type Client interface {
}

// UseCase defines the interface for product business logic
type UseCase interface {
	ImportProduct(tenantID id.ID,
		p []*glad.Product,
	) ([]*glad.ProductResponse, error)
	ImportCenter(tenantID id.ID,
		p []*glad.Center,
	) ([]*glad.CenterResponse, error)
}
