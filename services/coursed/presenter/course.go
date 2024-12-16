/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"

	"github.com/ulule/deepcopier"
)

// Course data - TenantID is returned in the HTTP header
// X-GLAD-TenantID
// type Course struct {
// 	ID           id.ID                `json:"id"`
// 	ExtID        *string              `json:"extID,omitempty"`
// 	Name         *string              `json:"name,omitempty"`
// 	CenterID     *id.ID               `json:"centerID,omitempty"`
// 	ProductID    *id.ID               `json:"productID,omitempty"`
// 	Notes        *string              `json:"notes,omitempty"`
// 	Timezone     *string              `json:"timezone,omitempty"`
// 	Address      *Address             `json:"address,omitempty"`
// 	Mode         *entity.CourseMode   `json:"mode,omitempty"`
// 	Status       *entity.CourseStatus `json:"status,omitempty"`
// 	MaxAttendees *int32               `json:"maxAttendees,omitempty"`
// 	NumAttendees *int32               `json:"numAttendees,omitempty"`
// }

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
	Notes        *string              `json:"notes,omitempty"`
	Timezone     string               `json:"timezone"`
	Address      *Address             `json:"address,omitempty"`
	Mode         entity.CourseMode    `json:"mode"`
	Status       *entity.CourseStatus `json:"status,omitempty"`
	MaxAttendees *int32               `json:"maxAttendees,omitempty"`
	Organizer    []id.ID              `json:"organizer,omitempty"`
	Contact      []id.ID              `json:"contact,omitempty"`
	Teacher      []CourseTeacher      `json:"teacher,omitempty"`
	Notify       []id.ID              `json:"notify,omitempty"`
	DateTime     []DateTime           `json:"date,omitempty"`
}

// CourseResponse struct used as response to the create course request (REST API)
type CourseResponse struct {
	ID         id.ID   `json:"id"`
	DateTimeID []id.ID `json:"dateID,omitempty"`
	ShortURL   *string `json:"shortURL,omitempty"`
}

type ImportCourseResponse struct {
	// todo: check the struct definition
	ID      id.ID  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
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

// Course data - TenantID is returned in the HTTP header
type Course struct {
	CourseResponse
	CourseReq
	NumAttendees *int32 `json:"numAttendees,omitempty"`
}

// FromEntityCourse creates course response from course entity
func (c *Course) FromEntityCourse(e *entity.Course) error {
	c.ID = e.ID
	c.Name = e.Name
	c.CenterID = e.CenterID
	c.ProductID = e.ProductID
	c.Notes = &e.Notes
	c.Timezone = e.Timezone
	c.Address = &Address{}
	c.Mode = e.Mode
	c.Status = &e.Status
	c.MaxAttendees = &e.MaxAttendees
	c.NumAttendees = &e.NumAttendees

	c.Address.CopyFrom(e.Address)

	return nil
}

// FromEntityCourseFull creates course response from course entity
func (c *Course) FromEntityCourseFull(cf *entity.CourseFull) error {
	c.FromEntityCourse(cf.Course)
	c.FromCourseOrganizer(cf.Cos)
	c.FromCourseTeacher(cf.Cts)
	c.FromCourseContact(cf.Ccs)
	c.FromCourseNotify(cf.Cns)
	c.FromCourseTiming(cf.CourseTiming)

	return nil
}

// FromCourseOrganizer creates course from course organizer
func (c *Course) FromCourseOrganizer(cos []*entity.CourseOrganizer) error {
	for _, co := range cos {
		c.Organizer = append(c.Organizer, co.ID)
	}

	return nil
}

// FromCourseTeacher creates course from course teacher
func (c *Course) FromCourseTeacher(cts []*entity.CourseTeacher) error {
	for _, ct := range cts {
		c.Teacher = append(c.Teacher, CourseTeacher{
			ID:        ct.ID,
			IsPrimary: ct.IsPrimary,
		})
	}

	return nil
}

// FromCourseContact creates course from course contact
func (c *Course) FromCourseContact(ccs []*entity.CourseContact) error {
	for _, cc := range ccs {
		c.Contact = append(c.Contact, cc.ID)
	}

	return nil
}

// FromCourseNotify creates course from course notify
func (c *Course) FromCourseNotify(cns []*entity.CourseNotify) error {
	for _, cn := range cns {
		c.Notify = append(c.Notify, cn.ID)
	}

	return nil
}

// FromCourseTiming creates course from course timing
func (c *Course) FromCourseTiming(cts []*entity.CourseTiming) error {
	for _, ct := range cts {
		c.DateTimeID = append(c.DateTimeID, ct.ID)
		c.DateTime = append(c.DateTime,
			DateTime{
				Date:      ct.DateTime.Date,
				StartTime: ct.DateTime.StartTime,
				EndTime:   ct.DateTime.EndTime,
			})
	}

	return nil
}

// ToEntity populates entity course from presenter course
func (c Course) ToEntity(e *entity.Course) error {
	deepcopier.Copy(c).To(e)
	l.Log.Infof("course full=%v, course=%v", c, e)
	return nil
}

// GladCourseToEntity populates entity course from glad entity
func GladCourseToEntity(gc glad.Course, e *entity.Course) error {
	deepcopier.Copy(gc).To(e)
	deepcopier.Copy(gc.Address).To(&e.Address)
	e.Mode = entity.CourseMode(gc.Mode)
	e.Status = entity.CourseStatus(gc.Status)
	e.ExtID = &gc.ExtID

	l.Log.Debugf("Course=%#v, entity.course=%#v", gc, e)
	return nil
}
