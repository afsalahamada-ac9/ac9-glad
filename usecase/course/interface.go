/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"
)

// CourseReader course reader
type CourseReader interface {
	Get(id id.ID) (*entity.Course, error)
	Search(tenantID id.ID, query string, page, limit int) ([]*entity.Course, error)
	List(tenantID id.ID, page, limit int) ([]*entity.Course, error)
	GetCount(id id.ID) (int, error)
	GetByAccount(tenantID id.ID,
		accountID id.ID,
		page, limit int,
	) (int, []*entity.Course, error)
}

// CourseWriter course writer
type CourseWriter interface {
	Create(e *entity.Course) (id.ID, error)
	Update(e *entity.Course) error
	Delete(id id.ID) error
}

// CourseOrganizerWriter course organizer writer
type CourseOrganizerWriter interface {
	InsertCourseOrganizer(id.ID, []*entity.CourseOrganizer) error
	UpdateCourseOrganizer(id.ID, []*entity.CourseOrganizer) error
	DeleteCourseOrganizer(id.ID, []*entity.CourseOrganizer) error
	DeleteCourseOrganizerByCourse(id.ID) error
}

// CourseOrganizerReader course organizer reader
type CourseOrganizerReader interface {
	GetCourseOrganizer(id.ID) ([]*entity.CourseOrganizer, error)
	MultiGetCourseOrganizer([]id.ID) ([][]*entity.CourseOrganizer, error)
}

// CourseTeacherWriter course teacher writer
type CourseTeacherWriter interface {
	InsertCourseTeacher(id.ID, []*entity.CourseTeacher) error
	UpdateCourseTeacher(id.ID, []*entity.CourseTeacher) error
	DeleteCourseTeacher(id.ID, []*entity.CourseTeacher) error
	DeleteCourseTeacherByCourse(id.ID) error
}

// CourseTeacherReader course teacher reader
type CourseTeacherReader interface {
	GetCourseTeacher(id.ID) ([]*entity.CourseTeacher, error)
	MultiGetCourseTeacher([]id.ID) ([][]*entity.CourseTeacher, error)
}

// CourseContactWriter course contact writer
type CourseContactWriter interface {
	InsertCourseContact(id.ID, []*entity.CourseContact) error
	UpdateCourseContact(id.ID, []*entity.CourseContact) error
	DeleteCourseContact(id.ID, []*entity.CourseContact) error
	DeleteCourseContactByCourse(id.ID) error
}

// CourseContactReader course contact reader
type CourseContactReader interface {
	GetCourseContact(id.ID) ([]*entity.CourseContact, error)
	MultiGetCourseContact([]id.ID) ([][]*entity.CourseContact, error)
}

// CourseNotifyWriter course notify writer
type CourseNotifyWriter interface {
	InsertCourseNotify(id.ID, []*entity.CourseNotify) error
	UpdateCourseNotify(id.ID, []*entity.CourseNotify) error
	DeleteCourseNotify(id.ID, []*entity.CourseNotify) error
	DeleteCourseNotifyByCourse(id.ID) error
}

// CourseNotifyReader course notify reader
type CourseNotifyReader interface {
	GetCourseNotify(id.ID) ([]*entity.CourseNotify, error)
	MultiGetCourseNotify([]id.ID) ([][]*entity.CourseNotify, error)
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
	Get(id id.ID) (*entity.CourseTiming, error)
	GetByCourse(courseID id.ID) ([]*entity.CourseTiming, error)
	GetCount() (int, error)
	MultiGetCourseTiming(courseIDList []id.ID) ([][]*entity.CourseTiming, error)
}

// CourseTimingWriter course timing writer
type CourseTimingWriter interface {
	Create(e *entity.CourseTiming) (id.ID, error)
	Update(e *entity.CourseTiming) error
	Delete(id id.ID) error
}

// Course timing repository interface
type CourseTimingRepository interface {
	CourseTimingReader
	CourseTimingWriter
}

// UseCase interface
type UseCase interface {
	GetCourse(id id.ID) (*entity.CourseFull, error)
	SearchCourses(tenantID id.ID, query string, page, limit int) ([]*entity.Course, error)
	ListCourses(tenantID id.ID, page, limit int) ([]*entity.Course, error)
	CreateCourse(
		course entity.Course,
		cos []*entity.CourseOrganizer,
		cts []*entity.CourseTeacher,
		ccs []*entity.CourseContact,
		cns []*entity.CourseNotify,
		courseTimings []*entity.CourseTiming,
	) (id.ID, []id.ID, error)
	UpdateCourse(
		course entity.Course,
		cos []*entity.CourseOrganizer,
		cts []*entity.CourseTeacher,
		ccs []*entity.CourseContact,
		cns []*entity.CourseNotify,
		courseTimings []*entity.CourseTiming,
	) error
	GetCourseByAccount(
		tenantID id.ID,
		accountID id.ID,
		page, limit int,
	) (int, []*entity.CourseFull, error)
	DeleteCourse(id id.ID) error
	GetCount(id id.ID) int
}
