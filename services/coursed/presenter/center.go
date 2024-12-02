/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"sudhagar/glad/entity"
)

// Center data - TenantID is returned in the HTTP header
// X-GLAD-TenantID
type Center struct {
	ID      entity.ID         `json:"id"`
	Name    string            `json:"name"`
	ExtName string            `json:"extName"`
	Mode    entity.CenterMode `json:"mode"`
}
