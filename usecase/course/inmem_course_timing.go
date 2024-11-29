/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"sudhagar/glad/entity"
)

// inmemCourseTiming in memory repo
type inmemCourseTiming struct {
	m map[entity.ID]*entity.CourseTiming
}

// newInmem create new repository
func newInmemCourseTiming() *inmemCourseTiming {
	var m = map[entity.ID]*entity.CourseTiming{}
	return &inmemCourseTiming{
		m: m,
	}
}

// Create a course
func (r *inmemCourseTiming) Create(e *entity.CourseTiming) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get a course
func (r *inmemCourseTiming) Get(id entity.ID) (*entity.CourseTiming, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update a course
func (r *inmemCourseTiming) Update(e *entity.CourseTiming) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// List courses
func (r *inmemCourseTiming) GetByCourse(courseID entity.ID) ([]*entity.CourseTiming, error) {
	var courses []*entity.CourseTiming
	for _, j := range r.m {
		if j.CourseID == courseID {
			courses = append(courses, j)
		}
	}
	return courses, nil
}

// Delete a course
func (r *inmemCourseTiming) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	delete(r.m, id)
	return nil
}

// GetCount gets total courses for a given tenant
func (r *inmemCourseTiming) GetCount() (int, error) {
	count := 0

	for range r.m {
		count++
	}
	return count, nil
}
