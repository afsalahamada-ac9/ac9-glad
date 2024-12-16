/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package glad

import "time"

// Note: Ideally these should be proto files and we should use grpc between services
type CourseAddress struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type Course struct {
	ExtID        string        `json:"extID"`
	CenterExtID  string        `json:"centerExtID"`
	ProductExtID string        `json:"productExtID"`
	Name         string        `json:"name"`
	Notes        string        `json:"notes"`
	Timezone     string        `json:"timezone"`
	Address      CourseAddress `json:"address"`
	Status       string        `json:"status"`
	Mode         string        `json:"mode"`
	MaxAttendees int           `json:"maxAttendees"`
	NumAttendees int           `json:"numAttendees"`
	URL          string        `json:"url"`
	CheckoutURL  string        `json:"checkoutURL"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
}

type CourseResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}
