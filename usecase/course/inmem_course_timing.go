/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"log"
)

// inmemCourseTiming in memory repo
type inmemCourseTiming struct {
	m map[id.ID]*entity.CourseTiming
}

// newInmem create new repository
func newInmemCourseTiming() *inmemCourseTiming {
	var m = map[id.ID]*entity.CourseTiming{}
	return &inmemCourseTiming{
		m: m,
	}
}

// Create a course
func (r *inmemCourseTiming) Create(e *entity.CourseTiming) (id.ID, error) {
	// log.Printf("Create course timing; id=%v, details=%#v", e.ID, e)
	r.m[e.ID] = e
	return e.ID, nil
}

// Get a course
func (r *inmemCourseTiming) Get(id id.ID) (*entity.CourseTiming, error) {
	if r.m[id] == nil {
		log.Printf("Inmem CourseTiming Get() id=%#v, not found", id)
		return nil, glad.ErrNotFound
	}
	return r.m[id], nil
}

// Update a course
func (r *inmemCourseTiming) Update(e *entity.CourseTiming) error {
	// log.Printf("Update course timing; id=%v, details=%#v", e.ID, e)
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// List courses
func (r *inmemCourseTiming) GetByCourse(courseID id.ID) ([]*entity.CourseTiming, error) {
	var courses []*entity.CourseTiming
	for _, j := range r.m {
		if j.CourseID == courseID {
			courses = append(courses, j)
		}
	}
	return courses, nil
}

// Delete a course
func (r *inmemCourseTiming) Delete(id id.ID) error {
	if r.m[id] == nil {
		return glad.ErrNotFound
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
