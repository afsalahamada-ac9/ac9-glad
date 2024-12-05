/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/entity"
)

// Participant data - TenantID is returned in the HTTP header (may be not, as account is global?)
// X-GLAD-TenantID
type Participant struct {
	ID      entity.ID `json:"id"`
	Email   string    `json:"email,omitempty"`
	Account *Account  `json:"account,omitempty"`
}
