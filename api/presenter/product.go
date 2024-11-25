/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"sudhagar/glad/entity"
)

// Product data - TenantID is returned in the HTTP header
// X-GLAD-TenantID
type Product struct {
	ID            entity.ID                `json:"id"`
	Name          string                   `json:"name"`
	Title         string                   `json:"title"`
	CType         string                   `json:"ctype"`
	BaseProductID string                   `json:"baseProductId,omitempty"`
	DurationDays  int32                    `json:"durationDays,omitempty"`
	Visibility    entity.ProductVisibility `json:"visibility,omitempty"`
	MaxAttendees  int32                    `json:"maxAttendees,omitempty"`
	Format        entity.ProductFormat     `json:"format,omitempty"`
}
