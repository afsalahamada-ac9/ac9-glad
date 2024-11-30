/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"sudhagar/glad/entity"
)

// CourseReader course reader
type CourseReader interface {
	Get(id entity.ID) (*entity.Course, error)
	Search(tenantID entity.ID, query string, page, limit int) ([]*entity.Course, error)
	List(tenantID entity.ID, page, limit int) ([]*entity.Course, error)
	GetCount(id entity.ID) (int, error)
}

// CourseWriter course writer
type CourseWriter interface {
	Create(e *entity.Course) (entity.ID, error)
	Update(e *entity.Course) error
	Delete(id entity.ID) error
	CreateCourseOrganizer(entity.ID, []*entity.CourseOrganizer) error
	CreateCourseTeacher(entity.ID, []*entity.CourseTeacher) error
	CreateCourseContact(entity.ID, []*entity.CourseContact) error
	CreateCourseNotify(entity.ID, []*entity.CourseNotify) error
}

// Course repository interface
type CourseRepository interface {
	CourseReader
	CourseWriter
}

// CourseTimingReader course timing reader
type CourseTimingReader interface {
	Get(id entity.ID) (*entity.CourseTiming, error)
	GetByCourse(courseID entity.ID) ([]*entity.CourseTiming, error)
	GetCount() (int, error)
}

// CourseTimingWriter course timing writer
type CourseTimingWriter interface {
	Create(e *entity.CourseTiming) (entity.ID, error)
	Update(e *entity.CourseTiming) error
	Delete(id entity.ID) error
}

// Course timing repository interface
type CourseTimingRepository interface {
	CourseTimingReader
	CourseTimingWriter
}

// UseCase interface
type UseCase interface {
	GetCourse(id entity.ID) (*entity.Course, error)
	SearchCourses(tenantID entity.ID, query string, page, limit int) ([]*entity.Course, error)
	ListCourses(tenantID entity.ID, page, limit int) ([]*entity.Course, error)
	CreateCourse(
		course entity.Course,
		cos []*entity.CourseOrganizer,
		cts []*entity.CourseTeacher,
		ccs []*entity.CourseContact,
		cns []*entity.CourseNotify,
		courseTimings []*entity.CourseTiming,
	) (entity.ID, []entity.ID, error)
	UpdateCourse(e *entity.Course) error
	DeleteCourse(id entity.ID) error
	GetCount(id entity.ID) int
}
