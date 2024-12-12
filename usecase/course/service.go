/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"database/sql"
	"strings"
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
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
) (id.ID, []id.ID, error) {
	c, err := course.New()
	if err != nil {
		return id.IDInvalid, nil, err
	}

	courseID, err := s.cRepo.Create(c)
	if err != nil {
		return courseID, nil, err
	}

	for i, ct := range courseTiming {
		ct.CourseID = courseID
		courseTiming[i], err = ct.Clone()
		if err != nil {
			return id.IDInvalid, nil, err
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

	var courseTimingID []id.ID
	for _, ct := range courseTiming {
		ctID, err := s.ctRepo.Create(ct)
		if err != nil {
			return courseID, nil, err
		}
		courseTimingID = append(courseTimingID, ctID)
	}

	return courseID, courseTimingID, err
}

// GetCourse retrieves a course and related information
func (s *Service) GetCourse(courseID id.ID) (*entity.CourseFull, error) {
	course, err := s.cRepo.Get(courseID)
	if course == nil {
		err = glad.ErrNotFound
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	cos, err := s.cRepo.GetCourseOrganizer(courseID)
	if err != nil {
		l.Log.Warnf("Unable to fetch organizers for course id=%v", courseID)
		return nil, err
	}

	cts, err := s.cRepo.GetCourseTeacher(courseID)
	if err != nil {
		l.Log.Warnf("Unable to fetch teachers for course id=%v", courseID)
		return nil, err
	}

	ccs, err := s.cRepo.GetCourseContact(courseID)
	if err != nil {
		l.Log.Warnf("Unable to fetch contact for course id=%v", courseID)
		return nil, err
	}

	cns, err := s.cRepo.GetCourseNotify(courseID)
	if err != nil {
		l.Log.Warnf("Unable to fetch notify for course id=%v", courseID)
		return nil, err
	}

	courseTiming, err := s.ctRepo.GetByCourse(courseID)
	if err != nil {
		l.Log.Warnf("Unable to fetch course timings for course id=%v", courseID)
		return nil, err
	}

	return entity.NewCourseFull(*course, cos, cts, ccs, cns, courseTiming), err
}

// SearchCourses searches course
// TODO: Return full information
func (s *Service) SearchCourses(tenantID id.ID,
	query string, page, limit int,
) ([]*entity.Course, error) {
	courses, err := s.cRepo.Search(tenantID, strings.ToLower(query), page, limit)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, glad.ErrNotFound
	}
	return courses, nil
}

// ListCourses lists course
// TODO: Return full information
func (s *Service) ListCourses(tenantID id.ID, page, limit int) ([]*entity.Course, error) {
	courses, err := s.cRepo.List(tenantID, page, limit)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, glad.ErrNotFound
	}
	return courses, nil
}

// DeleteCourse deletes a course
// Note: Since delete is cascaded to dependent tables, no need to call those functions explicitly
func (s *Service) DeleteCourse(id id.ID) error {
	err := s.cRepo.Delete(id)
	if err == sql.ErrNoRows {
		return glad.ErrNotFound
	}

	return err
}

// GetCount gets total course count
func (s *Service) GetCount(tenantID id.ID) int {
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

// GetCourseByAccount retrieves course and related information using account id
func (s *Service) GetCourseByAccount(courseID id.ID,
	accountID id.ID,
	page, limit int,
) (int, []*entity.CourseFull, error) {
	count, courseList, err := s.cRepo.GetByAccount(courseID, accountID, page, limit)
	if count == 0 {
		err = glad.ErrNotFound
		l.Log.Warnf("%v", err)
		return count, nil, err
	}
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	// Generate course ids
	var courseIDList []id.ID
	for _, course := range courseList {
		courseIDList = append(courseIDList, course.ID)
	}

	// Get all organizers, teachers, contact, notify and timings for N courses
	cosList, err := s.cRepo.MultiGetCourseOrganizer(courseIDList)
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	ctsList, err := s.cRepo.MultiGetCourseTeacher(courseIDList)
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	ccsList, err := s.cRepo.MultiGetCourseContact(courseIDList)
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	cnsList, err := s.cRepo.MultiGetCourseNotify(courseIDList)
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	courseTimingList, err := s.ctRepo.MultiGetCourseTiming(courseIDList)
	if err != nil {
		l.Log.Warnf("%v", err)
		return count, nil, err
	}

	// Merge all these in CourseFull entity and return a list of CourseFull entity & total count
	var cfList []*entity.CourseFull
	for i := range courseList {
		l.Log.Debugf("Course id=%v", courseList[i].ID)
		cfList = append(cfList,
			entity.NewCourseFull(*courseList[i], cosList[i], ctsList[i], ccsList[i], cnsList[i], courseTimingList[i]))
	}

	return count, cfList, err
}
