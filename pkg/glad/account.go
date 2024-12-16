/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package glad

import (
	"time"
)

// Note: Ideally these should be proto files and we should use grpc between services
// TODO: ctypes support to be added later
type Account struct {
	ExtID        string    `json:"extID"`
	CognitoID    string    `json:"cognitoID"`
	Username     string    `json:"username"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	FullPhotoURL string    `json:"fullPhotoURL"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type AccountResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}
