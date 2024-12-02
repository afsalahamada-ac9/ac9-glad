/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"strings"
	"time"

	"ac9/glad/entity"
)

// Service course usecase
type Service struct {
	cRepo  CourseRepository
	ctRepo CourseTimingRepository
}

// NewService creates new service
func NewService(cr CourseRepository, ctr CourseTimingRepository) *Service {
	return &Service{
		cRepo:  cr,
		ctRepo: ctr,
	}
}

// CreateCourse creates a course
func (s *Service) CreateCourse(
	course entity.Course,
	cos []*entity.CourseOrganizer,
	cts []*entity.CourseTeacher,
	ccs []*entity.CourseContact,
	cns []*entity.CourseNotify,
	courseTiming []*entity.CourseTiming,
) (entity.ID, []entity.ID, error) {
	c, err := course.New()
	if err != nil {
		return entity.IDInvalid, nil, err
	}

	courseID, err := s.cRepo.Create(c)
	if err != nil {
		return courseID, nil, err
	}

	for i, ct := range courseTiming {
		ct.CourseID = courseID
		courseTiming[i], err = ct.New()
		if err != nil {
			return entity.IDInvalid, nil, err
		}
	}

	// Note: These can be executed in parallel using go-routines
	err = s.cRepo.InsertCourseOrganizer(courseID, cos)
	if err != nil {
		return courseID, nil, err
	}

	err = s.cRepo.InsertCourseTeacher(courseID, cts)
	if err != nil {
		return courseID, nil, err
	}

	err = s.cRepo.InsertCourseContact(courseID, ccs)
	if err != nil {
		return courseID, nil, err
	}

	err = s.cRepo.InsertCourseNotify(courseID, cns)
	if err != nil {
		return courseID, nil, err
	}

	var courseTimingID []entity.ID
	for _, ct := range courseTiming {
		ctID, err := s.ctRepo.Create(ct)
		if err != nil {
			return courseID, nil, err
		}
		courseTimingID = append(courseTimingID, ctID)
	}

	return courseID, courseTimingID, err
}

// GetCourse retrieves a course
// TODO: Retrieve organizer, teacher, contact, notify and timing and return it
func (s *Service) GetCourse(id entity.ID) (*entity.Course, error) {
	t, err := s.cRepo.Get(id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchCourses searches course
func (s *Service) SearchCourses(tenantID entity.ID,
	query string, page, limit int,
) ([]*entity.Course, error) {
	courses, err := s.cRepo.Search(tenantID, strings.ToLower(query), page, limit)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, entity.ErrNotFound
	}
	return courses, nil
}

// ListCourses lists course
func (s *Service) ListCourses(tenantID entity.ID, page, limit int) ([]*entity.Course, error) {
	courses, err := s.cRepo.List(tenantID, page, limit)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, entity.ErrNotFound
	}
	return courses, nil
}

// DeleteCourse deletes a course
// Note: Since delete is cascaded to dependent tables, no need to call those functions explicitly
func (s *Service) DeleteCourse(id entity.ID) error {
	t, err := s.GetCourse(id)
	if t == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.cRepo.Delete(id)
}

// GetCount gets total course count
func (s *Service) GetCount(tenantID entity.ID) int {
	count, err := s.cRepo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}

// UpdateCourse updates course
func (s *Service) UpdateCourse(
	course entity.Course,
	cos []*entity.CourseOrganizer,
	cts []*entity.CourseTeacher,
	ccs []*entity.CourseContact,
	cns []*entity.CourseNotify,
	courseTiming []*entity.CourseTiming,
) error {

	err := course.Validate()
	if err != nil {
		return err
	}

	courseID := course.ID

	// This may not be needed
	course.UpdatedAt = time.Now()
	err = s.cRepo.Update(&course)
	if err != nil {
		return err
	}

	for _, ct := range courseTiming {
		ct.CourseID = courseID
	}

	// Note: These can be executed in parallel using go-routines
	err = s.cRepo.UpdateCourseOrganizer(courseID, cos)
	if err != nil {
		return err
	}

	err = s.cRepo.UpdateCourseTeacher(courseID, cts)
	if err != nil {
		return err
	}

	err = s.cRepo.UpdateCourseContact(courseID, ccs)
	if err != nil {
		return err
	}

	err = s.cRepo.UpdateCourseNotify(courseID, cns)
	if err != nil {
		return err
	}

	for _, ct := range courseTiming {
		err := s.ctRepo.Update(ct)
		if err != nil {
			return err
		}
	}

	return err
}
