/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/id"
	"time"
)

// Course date/time
type CourseDateTime struct {
	Date      string // Only date in YYYY-MM-DD format
	StartTime string // Only time in HH:MM:SS format (SS is optional, default 00)
	EndTime   string
}

// Course Timings
type CourseTiming struct {
	ID       id.ID
	CourseID id.ID
	ExtID    *string
	DateTime CourseDateTime

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewCourseDateTime creates a new course address
func NewCourseDateTime(
	date string,
	startTime string,
	endTime string,
) (*CourseDateTime, error) {

	dt := &CourseDateTime{
		Date:      date,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err := dt.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return dt, nil
}

// Validate validates course date/time
func (dt *CourseDateTime) Validate() error {
	if dt.Date == "" || dt.StartTime == "" || dt.EndTime == "" {
		return ErrInvalidEntity
	}
	return nil
}

// NewCourseTiming creates a new course timings
// Note: Tenant id is not needed here, because this is linked to the course internally.
// This object is not exposed externally via API. So, tenant ID can be mapped via the course.
func NewCourseTiming(
	courseID id.ID,
	extID *string,
	dateTime CourseDateTime,
) (*CourseTiming, error) {

	ct := &CourseTiming{
		ID:        id.New(),
		CourseID:  courseID,
		ExtID:     extID,
		DateTime:  dateTime,
		CreatedAt: time.Now(),
	}
	err := ct.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return ct, nil
}

// TODO: This must be renamed to Clone
// New creates a new course timing from existing course timing and overrides id and created & updated date
func (ct CourseTiming) New() (*CourseTiming, error) {
	courseTiming := &CourseTiming{}
	*courseTiming = ct

	courseTiming.ID = id.New()
	courseTiming.CreatedAt = time.Now()
	courseTiming.UpdatedAt = courseTiming.CreatedAt

	err := courseTiming.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return courseTiming, nil
}

// Validate validate course timings
func (ct *CourseTiming) Validate() error {
	if ct.CourseID == id.IDInvalid {
		return ErrInvalidEntity
	}

	return nil
}
