/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

// CourseFull complete course and related data such as organizer, teacher, contact, notify and timings
type CourseFull struct {
	Course       *Course
	Cos          []*CourseOrganizer
	Cts          []*CourseTeacher
	Ccs          []*CourseContact
	Cns          []*CourseNotify
	CourseTiming []*CourseTiming
}

// NewCourseFull creates a new course full data
func NewCourseFull(
	course Course,
	cos []*CourseOrganizer,
	cts []*CourseTeacher,
	ccs []*CourseContact,
	cns []*CourseNotify,
	courseTiming []*CourseTiming,
) *CourseFull {
	return &CourseFull{
		Course:       &course,
		Cos:          cos,
		Cts:          cts,
		Ccs:          ccs,
		Cns:          cns,
		CourseTiming: courseTiming,
	}
}
