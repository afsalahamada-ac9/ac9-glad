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
	"strings"
	"time"
)

// Course mode
type CourseMode string

const (
	CourseInPerson CourseMode = "in-person"
	CourseOnline   CourseMode = "online"
	CourseNotSet   CourseMode = "not-set"
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

// Course organizer
type CourseOrganizer struct {
	id.ID
}

// Course teacher
type CourseTeacher struct {
	ID        id.ID
	IsPrimary bool
}

// Course contact
type CourseContact struct {
	id.ID
}

// Course notify
type CourseNotify struct {
	id.ID
}

// Course data
type Course struct {
	ID        id.ID
	TenantID  id.ID
	CenterID  id.ID
	ProductID id.ID

	// TODO: Check whether sql.NullString is a better option
	ExtID *string

	Name     string
	Notes    string
	Timezone string

	Address CourseAddress
	Status  CourseStatus
	Mode    CourseMode

	MaxAttendees int32
	NumAttendees int32

	URL         string
	CheckoutURL string

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
		return nil, glad.ErrInvalidEntity
	}
	return l, nil
}

// Validate validates course address
func (ca *CourseAddress) Validate() error {
	if (ca.Street1 == "" && ca.Street2 == "") ||
		ca.City == "" || ca.State == "" ||
		ca.Zip == "" || ca.Country == "" {
		l.Log.Warnf("Address is not valid=%v", ca)
		return glad.ErrInvalidEntity
	}
	return nil
}

// NewCourse create a new course
func NewCourse(tenantID id.ID,
	extID *string,
	centerID id.ID,
	productID id.ID,
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
		ID:           id.New(),
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
		return nil, glad.ErrInvalidEntity
	}
	return c, nil
}

// New creates a new course from existing course and overrides id and created & updated date
func (c Course) New() (*Course, error) {
	course := &c

	course.ID = id.New()
	course.CreatedAt = time.Now()
	course.UpdatedAt = course.CreatedAt

	err := course.Validate()
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return course, nil
}

// Transform fixes the data issues, if any
func (c *Course) Transform() {
	c.Status = CourseStatus(strings.ToLower(string(c.Status)))
	c.Mode = CourseMode(strings.ToLower(string(c.Mode)))

	if c.Mode == "" {
		// TODO: Add a metric to keep track of this
		c.Mode = CourseNotSet
	}
}

// Validate validate course
func (c *Course) Validate() error {
	if c.Name == "" {
		l.Log.Warnf("Course name is empty; extID=%v", c.ExtID)
		return glad.ErrInvalidEntity
	}
	return c.Address.Validate()
}
