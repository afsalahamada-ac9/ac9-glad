/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"strings"
	"time"

	"sudhagar/glad/entity"
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
	courseTimings []*entity.CourseTiming,
) (entity.ID, error) {
	c, err := course.New()
	if err != nil {
		return entity.IDInvalid, err
	}

	courseID, err := s.cRepo.Create(c)
	if err != nil {
		return courseID, err
	}

	return courseID, err
}

// GetCourse retrieves a course
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

// UpdateCourse updates course
func (s *Service) UpdateCourse(c *entity.Course) error {
	err := c.Validate()
	if err != nil {
		return err
	}
	c.UpdatedAt = time.Now()
	return s.cRepo.Update(c)
}

// GetCount gets total course count
func (s *Service) GetCount(tenantID entity.ID) int {
	count, err := s.cRepo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
