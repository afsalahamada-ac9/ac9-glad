/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/id"
	"fmt"
	"time"
)

type LiveDarshan struct {
	ID         int64
	Date       string
	StartTime  time.Time
	MeetingURL string
	CreatedBy  id.ID
}

func NewLiveDarshan(id int64, date string, startTime time.Time, meetingUrl string, createdBy id.ID) (*LiveDarshan, error) {
	ld := &LiveDarshan{
		ID:         id,
		Date:       date,
		StartTime:  startTime,
		MeetingURL: meetingUrl,
		CreatedBy:  createdBy,
	}

	return ld, nil
}

func (ld *LiveDarshan) Validate() error {
	_, err := time.Parse("2006-01-02", ld.Date) // Using "YYYY-MM-DD" format
	if err != nil {
		return fmt.Errorf("date format is invalid: %w", err)
	}

	if ld.MeetingURL == "" {
		return fmt.Errorf("meeting URL is required")
	}

	return nil
}
