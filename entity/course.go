/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"time"
)

// Course mode
type CourseMode string

const (
	CourseInPerson CourseMode = "in-person"
	CourseOnline   CourseMode = "online"
	// Add new types here
)

// Course status
type CourseStatus string

const (
	CourseDraft            CourseStatus = "draft"
	CourseArchived         CourseStatus = "archived"
	CourseOpen             CourseStatus = "open"
	CourseExpenseSubmitted CourseStatus = "expense-submitted"
	CourseExpenseDeclined  CourseStatus = "expense-declined"
	CourseClosed           CourseStatus = "closed"
	CourseActive           CourseStatus = "active"
	CourseDeclined         CourseStatus = "declined"
	CourseSubmitted        CourseStatus = "submitted"
	CourseCanceled         CourseStatus = "canceled"
	CoursedInactive        CourseStatus = "inactive"
	// Add new types here
)

// Course Address
type CourseAddress struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string
	Country string
}

// Course date/time
type CourseDateTime struct {
	Date      string
	StartTime string
	EndTime   string
}

// Course data
type Course struct {
	ID        ID
	TenantID  ID
	CenterID  ID
	ProductID ID

	ExtID *string

	Name     string
	Notes    string
	Timezone string

	Address CourseAddress

	Status CourseStatus

	Mode CourseMode

	MaxAttendees int32
	NumAttendees int32

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewCourseAddress creates a new course address
func NewCourseAddress(street1 string,
	street2 string,
	city string,
	state string,
	zip string,
	country string) (*CourseAddress, error) {

	l := &CourseAddress{
		Street1: street1,
		Street2: street2,
		City:    city,
		State:   state,
		Zip:     zip,
		Country: country,
	}
	err := l.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return l, nil
}

// Validate validates course address
func (l *CourseAddress) Validate() error {
	if l.Street1 == "" || l.City == "" || l.State == "" || l.Zip == "" || l.Country == "" {
		return ErrInvalidEntity
	}
	return nil
}

// NewCourse create a new course
func NewCourse(tenantID ID,
	extID *string,
	centerID ID,
	productID ID,
	name string,
	notes string,
	timezone string,
	address CourseAddress,
	status CourseStatus,
	mode CourseMode,
	maxAttendees int32,
	numAttendees int32,
) (*Course, error) {
	c := &Course{
		ID:           NewID(),
		TenantID:     tenantID,
		ExtID:        extID,
		CenterID:     centerID,
		ProductID:    productID,
		Name:         name,
		Notes:        notes,
		Timezone:     timezone,
		Address:      address,
		Status:       status,
		Mode:         mode,
		MaxAttendees: maxAttendees,
		NumAttendees: numAttendees,
		CreatedAt:    time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return c, nil
}

// Validate validate course
func (c *Course) Validate() error {
	if c.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}
