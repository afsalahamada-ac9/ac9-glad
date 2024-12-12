/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"strings"

	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
)

// inmemCourse in memory repo
type inmemCourse struct {
	m map[id.ID]*entity.Course
}

// newinmemCourse create new repository
func newInmemCourse() *inmemCourse {
	var m = map[id.ID]*entity.Course{}
	return &inmemCourse{
		m: m,
	}
}

// Create a course
func (r *inmemCourse) Create(e *entity.Course) (id.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get a course
func (r *inmemCourse) Get(courseID id.ID) (*entity.Course, error) {
	if r.m[courseID] == nil {
		return nil, glad.ErrNotFound
	}
	return r.m[courseID], nil
}

// Update a course
func (r *inmemCourse) Update(e *entity.Course) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search courses
func (r *inmemCourse) Search(tenantID id.ID,
	query string, page, limit int,
) ([]*entity.Course, error) {
	var courses []*entity.Course
	for _, j := range r.m {
		if j.TenantID == tenantID &&
			strings.Contains(strings.ToLower(j.Name), query) {
			courses = append(courses, j)
		}
	}

	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(courses) {
			return []*entity.Course{}, nil
		}
		if end > len(courses) {
			end = len(courses)
		}
		return courses[start:end], nil
	}

	return courses, nil
}

// List courses
func (r *inmemCourse) List(tenantID id.ID, page, limit int) ([]*entity.Course, error) {
	var courses []*entity.Course
	for _, j := range r.m {
		if j.TenantID == tenantID {
			courses = append(courses, j)
		}
	}
	return courses, nil
}

// Delete a course
func (r *inmemCourse) Delete(courseID id.ID) error {
	if r.m[courseID] == nil {
		return glad.ErrNotFound
	}
	r.m[courseID] = nil
	delete(r.m, courseID)
	return nil
}

// GetCount gets total courses for a given tenant
func (r *inmemCourse) GetCount(tenantID id.ID) (int, error) {
	count := 0
	for _, j := range r.m {
		if j.TenantID == tenantID {
			count++
		}
	}
	return count, nil
}

// GetByAccount a course
// TODO: implement
func (r *inmemCourse) GetByAccount(tenantID id.ID, accountID id.ID, page, limit int,
) (int, []*entity.Course, error) {
	return 0, nil, nil
}

// --------------------------------------------------------------------------------
// Course Organizer
// --------------------------------------------------------------------------------
// InsertCourseOrganizer creates course to organizer mapping
func (r *inmemCourse) InsertCourseOrganizer(courseID id.ID, cos []*entity.CourseOrganizer) error {
	// TODO
	return nil
}

// GetCourseOrganizer gets course organizer for the given course id
func (r *inmemCourse) GetCourseOrganizer(courseID id.ID) ([]*entity.CourseOrganizer, error) {
	// TODO
	return nil, nil
}

// UpdateCourseOrganizer updates course organizer for the given course id and the organizer
func (r *inmemCourse) UpdateCourseOrganizer(courseID id.ID, cos []*entity.CourseOrganizer) error {
	// TODO
	return nil
}

// DeleteCourseOrganizer deletes the given course organizers
func (r *inmemCourse) DeleteCourseOrganizer(courseID id.ID, cos []*entity.CourseOrganizer) error {
	// TODO
	return nil
}

// DeleteCourseOrganizerByCourse deletes course organizers using course id
func (r *inmemCourse) DeleteCourseOrganizerByCourse(courseID id.ID) error {
	// TODO
	return nil
}

// MultiGetCourseOrganizer gets course organizer for the given course ids
func (r *inmemCourse) MultiGetCourseOrganizer(courseIDList []id.ID,
) ([][]*entity.CourseOrganizer, error) {
	// TODO
	return nil, nil
}

// --------------------------------------------------------------------------------
// Course Teacher
// --------------------------------------------------------------------------------
// InsertCourseTeacher creates course to teacher mapping
func (r *inmemCourse) InsertCourseTeacher(courseID id.ID, cts []*entity.CourseTeacher) error {
	// TODO
	return nil
}

// GetCourseTeacher gets course teacher for the given course id
func (r *inmemCourse) GetCourseTeacher(courseID id.ID) ([]*entity.CourseTeacher, error) {
	// TODO
	return nil, nil
}

// UpdateCourseTeacher updates course teacher for the given course id and the teacher
func (r *inmemCourse) UpdateCourseTeacher(courseID id.ID, cos []*entity.CourseTeacher) error {
	// TODO
	return nil
}

// DeleteCourseTeacher deletes the given course teachers
func (r *inmemCourse) DeleteCourseTeacher(courseID id.ID, cos []*entity.CourseTeacher) error {
	// TODO
	return nil
}

// DeleteCourseTeacherByCourse deletes course teachers using course id
func (r *inmemCourse) DeleteCourseTeacherByCourse(courseID id.ID) error {
	// TODO
	return nil
}

// MultiGetCourseTeacher gets course teacher for the given course ids
func (r *inmemCourse) MultiGetCourseTeacher(courseIDList []id.ID,
) ([][]*entity.CourseTeacher, error) {
	// TODO
	return nil, nil
}

// --------------------------------------------------------------------------------
// Course Contact
// --------------------------------------------------------------------------------
// InsertCourseContact creates course to contact mapping
func (r *inmemCourse) InsertCourseContact(courseID id.ID, ccs []*entity.CourseContact) error {
	// TODO
	return nil
}

// GetCourseContact gets course contact for the given course id
func (r *inmemCourse) GetCourseContact(courseID id.ID) ([]*entity.CourseContact, error) {
	// TODO
	return nil, nil
}

// UpdateCourseContact updates course contact for the given course id and the contact
func (r *inmemCourse) UpdateCourseContact(courseID id.ID, cos []*entity.CourseContact) error {
	// TODO
	return nil
}

// DeleteCourseContact deletes the given course contacts
func (r *inmemCourse) DeleteCourseContact(courseID id.ID, cos []*entity.CourseContact) error {
	// TODO
	return nil
}

// DeleteCourseContactByCourse deletes course contacts using course id
func (r *inmemCourse) DeleteCourseContactByCourse(courseID id.ID) error {
	// TODO
	return nil
}

// MultiGetCourseContact gets course contact for the given course ids
func (r *inmemCourse) MultiGetCourseContact(courseIDList []id.ID,
) ([][]*entity.CourseContact, error) {
	// TODO
	return nil, nil
}

// --------------------------------------------------------------------------------
// Course Notify
// --------------------------------------------------------------------------------
// InsertCourseNotify creates course to notify mapping
func (r *inmemCourse) InsertCourseNotify(courseID id.ID, cns []*entity.CourseNotify) error {
	// TODO
	return nil
}

// GetCourseNotify gets course notify for the given course id
func (r *inmemCourse) GetCourseNotify(courseID id.ID) ([]*entity.CourseNotify, error) {
	// TODO
	return nil, nil
}

// UpdateCourseNotify updates course notify for the given course id and the notify
func (r *inmemCourse) UpdateCourseNotify(courseID id.ID, cos []*entity.CourseNotify) error {
	// TODO
	return nil
}

// DeleteCourseNotify deletes the given course notifys
func (r *inmemCourse) DeleteCourseNotify(courseID id.ID, cos []*entity.CourseNotify) error {
	// TODO
	return nil
}

// DeleteCourseNotifyByCourse deletes course notifys using course id
func (r *inmemCourse) DeleteCourseNotifyByCourse(courseID id.ID) error {
	// TODO
	return nil
}

// MultiGetCourseNotify gets course notify for the given course ids
func (r *inmemCourse) MultiGetCourseNotify(courseIDList []id.ID,
) ([][]*entity.CourseNotify, error) {
	// TODO
	return nil, nil
}
