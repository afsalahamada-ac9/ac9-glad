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
}

// CourseOrganizerWriter course organizer writer
type CourseOrganizerWriter interface {
	InsertCourseOrganizer(entity.ID, []*entity.CourseOrganizer) error
	UpdateCourseOrganizer(entity.ID, []*entity.CourseOrganizer) error
	DeleteCourseOrganizer(entity.ID, []*entity.CourseOrganizer) error
	DeleteCourseOrganizerByCourse(entity.ID) error
}

// CourseOrganizerReader course organizer reader
type CourseOrganizerReader interface {
	GetCourseOrganizer(entity.ID) ([]*entity.CourseOrganizer, error)
}

// CourseTeacherWriter course teacher writer
type CourseTeacherWriter interface {
	InsertCourseTeacher(entity.ID, []*entity.CourseTeacher) error
	// UpdateCourseTeacher(entity.ID, []*entity.CourseTeacher) error
	// DeleteCourseTeacher(entity.ID, []*entity.CourseTeacher) error
	// DeleteCourseTeacherByCourse(entity.ID) error
}

// CourseTeacherReader course teacher reader
type CourseTeacherReader interface {
	// GetCourseTeacher(entity.ID) ([]*entity.CourseTeacher, error)
}

// CourseContactWriter course contact writer
type CourseContactWriter interface {
	InsertCourseContact(entity.ID, []*entity.CourseContact) error
	// UpdateCourseContact(entity.ID, []*entity.CourseContact) error
	// DeleteCourseContact(entity.ID, []*entity.CourseContact) error
	// DeleteCourseContactByCourse(entity.ID) error
}

// CourseContactReader course contact reader
type CourseContactReader interface {
	// GetCourseContact(entity.ID) ([]*entity.CourseContact, error)
}

// CourseNotifyWriter course notify writer
type CourseNotifyWriter interface {
	InsertCourseNotify(entity.ID, []*entity.CourseNotify) error
	// UpdateCourseNotify(entity.ID, []*entity.CourseNotify) error
	// DeleteCourseNotify(entity.ID, []*entity.CourseNotify) error
	// DeleteCourseNotifyByCourse(entity.ID) error
}

// CourseNotifyReader course notify reader
type CourseNotifyReader interface {
	// GetCourseNotify(entity.ID) ([]*entity.CourseNotify, error)
}

// Course repository interface
type CourseRepository interface {
	CourseReader
	CourseWriter
	CourseOrganizerWriter
	CourseOrganizerReader
	CourseTeacherWriter
	CourseTeacherReader
	CourseContactWriter
	CourseContactReader
	CourseNotifyWriter
	CourseNotifyReader
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
	UpdateCourse(
		course entity.Course,
		cos []*entity.CourseOrganizer,
		cts []*entity.CourseTeacher,
		ccs []*entity.CourseContact,
		cns []*entity.CourseNotify,
		courseTimings []*entity.CourseTiming,
	) error
	DeleteCourse(id entity.ID) error
	GetCount(id entity.ID) int
}
