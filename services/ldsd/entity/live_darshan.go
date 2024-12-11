/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"

	"time"
)

// LiveDarshan contains live darshan information
type LiveDarshan struct {
	ID         id.ID
	TenantID   id.ID
	Date       time.Time // YYYY-MM-DD format
	StartTime  time.Time // HH:mm:ss format
	MeetingURL string
	CreatedBy  id.ID
	UpdatedBy  id.ID // last updated

	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewLiveDarshan creates a new live darhsan entity
func NewLiveDarshan(
	tenantID id.ID,
	date string,
	startTime string,
	meetingUrl string,
	createdBy id.ID,
) (*LiveDarshan, error) {

	ld := &LiveDarshan{
		ID:         id.New(),
		TenantID:   tenantID,
		MeetingURL: meetingUrl,
		CreatedBy:  createdBy,
		UpdatedBy:  createdBy,
	}

	tDate, err := time.Parse("2006-01-02", date) // Using "YYYY-MM-DD" format
	if err != nil {
		l.Log.Warnf("Date value=%v is not valid", ld.Date)
		return nil, glad.ErrInvalidValue
	}

	tTime, err := time.Parse("15:04:00", startTime)
	if err != nil {
		l.Log.Warnf("Start time value=%v is not valid", ld.StartTime)
		return nil, glad.ErrInvalidValue
	}

	ld.Date = tDate
	ld.StartTime = tTime

	err = ld.Validate()
	return ld, err
}

// Validate validates the live darshan parameters
func (ld *LiveDarshan) Validate() error {

	if ld.MeetingURL == "" {
		l.Log.Warnf("Meeting URL is required")
		return glad.ErrMissingParam
	}

	return nil
}
