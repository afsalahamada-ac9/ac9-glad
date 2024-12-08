/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"
)

// Center data - TenantID is returned in the HTTP header
// X-GLAD-TenantID
type Center struct {
	ID      id.ID             `json:"id"`
	Name    string            `json:"name"`
	ExtName string            `json:"extName"`
	Mode    entity.CenterMode `json:"mode"`
}
