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
type Product struct {
	ExtID            string    `json:"extID"`
	ExtName          string    `json:"extName"`
	Title            string    `json:"title"`
	CType            string    `json:"ctype"`
	BaseProductExtID string    `json:"baseProductExtID"`
	DurationDays     int32     `json:"durationDays"`
	Visibility       string    `json:"visibility"`
	MaxAttendees     int32     `json:"maxAttendees"`
	Format           string    `json:"format"`
	IsAutoApprove    bool      `json:"isAutoApprove"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type ProductResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}
