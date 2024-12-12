/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ac9/glad/entity"
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"ac9/glad/services/coursed/presenter"

	amock "ac9/glad/usecase/account/mock"
	mock "ac9/glad/usecase/course/mock"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
)

func Test_listCourses(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	aservice := amock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service, aservice)
	path, err := r.GetRoute("listCourses").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses", path)
	course := &entity.Course{
		ID:       id.New(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(course.TenantID).Return(1)
	service.EXPECT().
		ListCourses(course.TenantID, gomock.Any(), gomock.Any()).
		Return([]*entity.Course{course}, nil)
	ts := httptest.NewServer(listCourses(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listCourses_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listCourses(service))
	defer ts.Close()
	tenantID := tenantAlice
	service.EXPECT().GetCount(tenantID).Return(0)
	service.EXPECT().
		SearchCourses(tenantID, "non-existent", gomock.Any(), gomock.Any()).
		Return(nil, glad.ErrNotFound)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet,
		ts.URL+"?q=non-existent",
		nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listCourses_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	course := &entity.Course{
		ID:       id.New(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(course.TenantID).Return(1)
	service.EXPECT().
		SearchCourses(course.TenantID, "default", gomock.Any(), gomock.Any()).
		Return([]*entity.Course{course}, nil)
	ts := httptest.NewServer(listCourses(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet,
		ts.URL+"?q=default",
		nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createCourse(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	aservice := amock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service, aservice)
	path, err := r.GetRoute("createCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses", path)

	courseID := id.New()
	// TODO: Check the length of the courseTimings and same number of
	// courseTimingIDs to be returned
	service.EXPECT().
		CreateCourse(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).
		Return(courseID, nil, nil)
	h := createCourse(service)

	ts := httptest.NewServer(h)
	defer ts.Close()

	payload := struct {
		// TenantID id.ID         `json:"tenant_id"`
		ExtID string            `json:"extID"`
		Name  string            `json:"name"`
		Mode  entity.CourseMode `json:"mode"`
		// CenterID id.ID         `json:"center_id"`
	}{
		// TenantID: tenantAlice,
		ExtID: aliceExtID,
		Name:  "default-0",
		Mode:  (entity.CourseInPerson),
		// CenterID: aliceCenterID,
	}
	payloadBytes, err := json.Marshal(payload)
	assert.Nil(t, err)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost,
		ts.URL+"/v1/courses",
		bytes.NewReader(payloadBytes))
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var course *presenter.Course
	json.NewDecoder(res.Body).Decode(&course)
	assert.Equal(t, courseID, course.ID)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_getCourse(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	aservice := amock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service, aservice)
	path, err := r.GetRoute("getCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses/{id}", path)
	extID := aliceExtID
	course := &entity.Course{
		ID:       id.New(),
		TenantID: tenantAlice,
		ExtID:    &extID,
		Name:     "default-0",
		Mode:     entity.CourseInPerson,
	}
	courseFull := &entity.CourseFull{
		Course: course,
	}
	service.EXPECT().
		GetCourse(course.ID).
		Return(courseFull, nil)
	handler := getCourse(service)
	r.Handle("/v1/courses/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/courses/" + course.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// presenter.Course is returned by the api (http) server
	var d *presenter.Course
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)

	assert.Equal(t, course.ID, d.ID)
	assert.Equal(t, course.Name, d.Name)
	assert.Equal(t, course.Mode, d.Mode)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_deleteCourse(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	aservice := amock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service, aservice)
	path, err := r.GetRoute("deleteCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses/{id}", path)

	id := id.New()
	service.EXPECT().DeleteCourse(id).Return(nil)
	handler := deleteCourse(service)
	req, _ := http.NewRequest("DELETE", "/v1/courses/"+id.String(), nil)
	r.Handle("/v1/courses/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_deleteCourseNonExistent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	aservice := amock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service, aservice)
	path, err := r.GetRoute("deleteCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses/{id}", path)

	id := id.New()
	service.EXPECT().DeleteCourse(id).Return(glad.ErrNotFound)
	handler := deleteCourse(service)
	req, _ := http.NewRequest("DELETE", "/v1/courses/"+id.String(), nil)
	r.Handle("/v1/courses/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// TODO: Test case for updating course
