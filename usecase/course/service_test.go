/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"testing"
	"time"

	"ac9/glad/entity"

	"github.com/stretchr/testify/assert"
)

const (
	courseDefault  entity.ID = 13790493495087071234
	tenantAlice    entity.ID = 13790492210917015554
	aliceExtID               = "000aliceExtID"
	aliceCenterID            = 13790493495087075501
	aliceProductID           = 13790493495087076601

	// todo: add multi-tenant support
	// tenantBob    entity.ID = 13790492210917015555
	bobExtID    = "000bobExtID"
	bobCenterID = 13790493495087075502

	aliceOrganizer1ID = 100000001
	aliceOrganizer2ID = 100000002

	aliceTeacher1ID = 200000001
	aliceTeacher2ID = 200000002

	aliceTiming1ID = 300000001
	aliceTiming2ID = 300000002
	aliceTiming3ID = 300000003
)

func newFixtureCourse() *entity.Course {
	extID := aliceExtID

	return &entity.Course{
		ID:        courseDefault,
		TenantID:  tenantAlice,
		ExtID:     &extID,
		CenterID:  aliceCenterID,
		ProductID: aliceProductID,
		Name:      "Course Part 1",
		Notes:     "This is a course notes. It can be multi-line text. The notes can be longer than this.",
		Timezone:  "PST",
		Address: entity.CourseAddress{
			Street1: "1 Street Way",
			Street2: "",
			City:    "CityName",
			State:   "California",
			Country: "USA",
			Zip:     "12345-6789",
		},
		Status:       entity.CourseActive,
		Mode:         entity.CourseInPerson,
		MaxAttendees: 50,
		NumAttendees: 12,
		CreatedAt:    time.Now(),
	}
}

func newCourseOrganizer() []*entity.CourseOrganizer {
	co1 := entity.CourseOrganizer{ID: aliceOrganizer1ID}
	co2 := entity.CourseOrganizer{ID: aliceOrganizer2ID}

	var cos []*entity.CourseOrganizer
	cos = append(cos, &co1)
	cos = append(cos, &co2)

	return cos
}

func newCourseTeacher() []*entity.CourseTeacher {
	t1 := entity.CourseTeacher{ID: aliceTeacher1ID, IsPrimary: true}
	t2 := entity.CourseTeacher{ID: aliceTeacher2ID, IsPrimary: false}

	var cts []*entity.CourseTeacher
	cts = append(cts, &t1)
	cts = append(cts, &t2)

	return cts
}

func newCourseContact() []*entity.CourseContact {
	c1 := entity.CourseContact{ID: aliceTeacher1ID}
	c2 := entity.CourseContact{ID: aliceOrganizer1ID}

	var ccs []*entity.CourseContact
	ccs = append(ccs, &c1)
	ccs = append(ccs, &c2)

	return ccs
}

func newCourseNotify() []*entity.CourseNotify {
	cn1 := entity.CourseNotify{ID: aliceTeacher1ID}
	cn2 := entity.CourseNotify{ID: aliceOrganizer1ID}
	cn3 := entity.CourseNotify{ID: aliceTeacher2ID}

	var cns []*entity.CourseNotify
	cns = append(cns, &cn1)
	cns = append(cns, &cn2)
	cns = append(cns, &cn3)

	return cns
}

// TODO: Need to set proper dates, time and verify the same
func newCourseTiming() []*entity.CourseTiming {
	ct1 := entity.CourseTiming{ID: aliceTiming1ID}
	ct2 := entity.CourseTiming{ID: aliceTiming2ID}
	ct3 := entity.CourseTiming{ID: aliceTiming3ID}

	var cts []*entity.CourseTiming
	cts = append(cts, &ct1)
	cts = append(cts, &ct2)
	cts = append(cts, &ct3)

	return cts
}

func Test_Create(t *testing.T) {
	repo := newInmemCourse()
	ctRepo := newInmemCourseTiming()
	m := NewService(repo, ctRepo)
	tmpl := newFixtureCourse()
	_, _, err := m.CreateCourse(
		*tmpl,
		newCourseOrganizer(),
		newCourseTeacher(),
		newCourseContact(),
		newCourseNotify(),
		newCourseTiming(),
	)

	assert.Nil(t, err)
	assert.False(t, tmpl.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmemCourse()
	ctRepo := newInmemCourseTiming()
	m := NewService(repo, ctRepo)
	tmpl1 := newFixtureCourse()
	tmpl2 := newFixtureCourse()
	tmpl2.Name = "Course Sahaj Meditation"
	extID := bobExtID
	tmpl2.ExtID = &extID

	tID, _, _ := m.CreateCourse(
		*tmpl1,
		newCourseOrganizer(),
		newCourseTeacher(),
		newCourseContact(),
		newCourseNotify(),
		newCourseTiming(),
	)
	_, _, _ = m.CreateCourse(
		*tmpl2,
		newCourseOrganizer(),
		newCourseTeacher(),
		newCourseContact(),
		newCourseNotify(),
		newCourseTiming(),
	)

	t.Run("search", func(t *testing.T) {
		res, err := m.SearchCourses(tmpl1.TenantID, "Part", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, tmpl1.ExtID, res[0].ExtID)
		assert.Equal(t, tmpl1.CenterID, res[0].CenterID)
		assert.Equal(t, tmpl1.Status, res[0].Status)
		// TODO: checks for other fields to be added

		// 'default' query value matches both the course names
		res, err = m.SearchCourses(tmpl1.TenantID, "Sahaj", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))

		res, err = m.SearchCourses(tmpl1.TenantID, "non-existent", 0, 0)
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, res)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListCourses(tmpl1.TenantID, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetCourse(tID)
		assert.Nil(t, err)
		assert.Equal(t, tmpl1.TenantID, saved.TenantID)
		assert.Equal(t, tmpl1.ExtID, saved.ExtID)
		assert.Equal(t, tmpl1.CenterID, saved.CenterID)
		assert.Equal(t, tmpl1.Status, saved.Status)
		assert.Equal(t, tmpl1.Name, saved.Name)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmemCourse()
	ctRepo := newInmemCourseTiming()
	m := NewService(repo, ctRepo)
	tmpl := newFixtureCourse()
	courseTiming := newCourseTiming()

	id, courseTimingIDs, err := m.CreateCourse(
		*tmpl,
		newCourseOrganizer(),
		newCourseTeacher(),
		newCourseContact(),
		newCourseNotify(),
		courseTiming,
	)
	assert.Nil(t, err)

	saved, _ := m.GetCourse(id)
	saved.Mode = entity.CourseOnline
	// Update the course timing ids as it is generated afresh
	for i := range courseTiming {
		courseTiming[i].ID = courseTimingIDs[i]
	}

	assert.Nil(t, m.UpdateCourse(
		*saved,
		newCourseOrganizer(),
		newCourseTeacher(),
		newCourseContact(),
		newCourseNotify(),
		courseTiming,
	))

	updated, err := m.GetCourse(id)
	assert.Nil(t, err)
	assert.Equal(t, entity.CourseOnline, updated.Mode)
}

func TestDelete(t *testing.T) {
	repo := newInmemCourse()
	ctRepo := newInmemCourseTiming()
	m := NewService(repo, ctRepo)

	tmpl1 := newFixtureCourse()
	tmpl2 := newFixtureCourse()
	extID := bobExtID
	tmpl2.ExtID = &extID
	t2ID, _, _ := m.CreateCourse(
		*tmpl2,
		newCourseOrganizer(),
		newCourseTeacher(),
		newCourseContact(),
		newCourseNotify(),
		newCourseTiming(),
	)

	err := m.DeleteCourse(tmpl1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteCourse(t2ID)
	assert.Nil(t, err)
	_, err = m.GetCourse(t2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
