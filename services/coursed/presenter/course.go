/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"

	"github.com/ulule/deepcopier"
)

// Course data - TenantID is returned in the HTTP header
// X-GLAD-TenantID
type Course struct {
	ID           id.ID                `json:"id"`
	ExtID        *string              `json:"extID,omitempty"`
	CenterID     *id.ID               `json:"centerID,omitempty"`
	Name         *string              `json:"name,omitempty"`
	Notes        *string              `json:"notes,omitempty"`
	Timezone     *string              `json:"timezone,omitempty"`
	Address      *Address             `json:"address,omitempty"`
	Status       *entity.CourseStatus `json:"status,omitempty"`
	Mode         *entity.CourseMode   `json:"mode,omitempty"`
	MaxAttendees *int32               `json:"maxAttendees,omitempty"`
	NumAttendees *int32               `json:"numAttendees,omitempty"`
}

// Course teacher
type CourseTeacher struct {
	ID        id.ID `json:"id"`
	IsPrimary bool  `json:"is_primary"`
}

// CourseReq struct used to create & update the course via REST API
// TODO: Salesforce will send additional details: extID, URL (to be converted to shortURL), numAttendees
type CourseReq struct {
	Name         string               `json:"name"`
	CenterID     id.ID                `json:"centerID"`
	ProductID    id.ID                `json:"productID"`
	Mode         entity.CourseMode    `json:"mode"`
	Timezone     string               `json:"timezone"`
	Organizer    []id.ID              `json:"organizer"`
	Contact      []id.ID              `json:"contact"`
	Teacher      []CourseTeacher      `json:"teacher"`
	Notes        *string              `json:"notes"`
	Status       *entity.CourseStatus `json:"status"`
	MaxAttendees *int32               `json:"maxAttendees"`
	DateTime     []DateTime           `json:"date"`
	Address      *Address             `json:"address"`
	Notify       []id.ID              `json:"notify"`
}

// CourseResponse struct used as response to the create course request (REST API)
type CourseResponse struct {
	ID         id.ID   `json:"id"`
	DateTimeID []id.ID `json:"dateID"`
	ShortURL   *string `json:"shortURL,omitempty"`
}

// ToCourse creates course entity from course request
func (cr CourseReq) ToCourse(tenantID id.ID) (entity.Course, error) {

	var course entity.Course
	deepcopier.Copy(cr).To(&course)
	if cr.Address != nil {
		deepcopier.Copy(cr.Address).To(&course.Address)
	}

	course.TenantID = tenantID

	return course, nil
}

// ToCourseOrganizer creates course organizer from course request
func (cr CourseReq) ToCourseOrganizer() ([]*entity.CourseOrganizer, error) {
	var cos []*entity.CourseOrganizer
	for _, id := range cr.Organizer {
		co := entity.CourseOrganizer{
			ID: id,
		}

		cos = append(cos, &co)
	}

	return cos, nil
}

// ToCourseTeacher creates course teacher from course request
func (cr CourseReq) ToCourseTeacher() ([]*entity.CourseTeacher, error) {
	var cts []*entity.CourseTeacher
	for _, t := range cr.Teacher {
		ct := entity.CourseTeacher{
			ID:        t.ID,
			IsPrimary: t.IsPrimary,
		}

		cts = append(cts, &ct)
	}

	return cts, nil
}

// ToCourseContact creates course organizer from course request
func (cr CourseReq) ToCourseContact() ([]*entity.CourseContact, error) {
	var ccs []*entity.CourseContact
	for _, id := range cr.Contact {
		cc := entity.CourseContact{
			ID: id,
		}

		ccs = append(ccs, &cc)
	}

	return ccs, nil
}

// ToCourseNotify creates course organizer from course request
func (cr CourseReq) ToCourseNotify() ([]*entity.CourseNotify, error) {
	var cns []*entity.CourseNotify
	for _, id := range cr.Notify {
		cn := entity.CourseNotify{
			ID: id,
		}

		cns = append(cns, &cn)
	}

	return cns, nil
}

// ToCourseTiming creates course timing entity from course request
func (cr CourseReq) ToCourseTiming() ([]*entity.CourseTiming, error) {
	var cts []*entity.CourseTiming
	for _, dt := range cr.DateTime {
		ct := entity.CourseTiming{
			DateTime: entity.CourseDateTime{
				Date:      dt.Date,
				StartTime: dt.StartTime,
				EndTime:   dt.EndTime,
			},
		}

		cts = append(cts, &ct)
	}

	return cts, nil
}

// FromCourseEntity creates course response from course entity
func (c *Course) FromCourseEntity(e *entity.Course) error {

	c.ID = e.ID
	c.Name = &e.Name
	c.Mode = &e.Mode
	c.CenterID = &e.CenterID
	c.Notes = &e.Notes
	c.Timezone = &e.Timezone
	c.Status = &e.Status
	c.MaxAttendees = &e.MaxAttendees
	c.NumAttendees = &e.NumAttendees

	c.Address = &Address{}
	c.Address.CopyFrom(e.Address)

	return nil
}
