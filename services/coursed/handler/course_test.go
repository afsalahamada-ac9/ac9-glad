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
	"ac9/glad/services/coursed/presenter"

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
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service)
	path, err := r.GetRoute("listCourses").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses", path)
	tmpl := &entity.Course{
		ID:       entity.NewID(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(tmpl.TenantID).Return(1)
	service.EXPECT().
		ListCourses(tmpl.TenantID, gomock.Any(), gomock.Any()).
		Return([]*entity.Course{tmpl}, nil)
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
		Return(nil, entity.ErrNotFound)

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
	tmpl := &entity.Course{
		ID:       entity.NewID(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(tmpl.TenantID).Return(1)
	service.EXPECT().
		SearchCourses(tmpl.TenantID, "default", gomock.Any(), gomock.Any()).
		Return([]*entity.Course{tmpl}, nil)
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
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service)
	path, err := r.GetRoute("createCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses", path)

	id := entity.NewID()
	// TODO: Check the length of the courseTimings and same number of
	// courseTimingIDs to be returned
	service.EXPECT().
		CreateCourse(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).
		Return(id, nil, nil)
	h := createCourse(service)

	ts := httptest.NewServer(h)
	defer ts.Close()

	payload := struct {
		// TenantID entity.ID         `json:"tenant_id"`
		ExtID string            `json:"extID"`
		Name  string            `json:"name"`
		Mode  entity.CourseMode `json:"mode"`
		// CenterID entity.ID         `json:"center_id"`
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

	var tmpl *presenter.Course
	json.NewDecoder(res.Body).Decode(&tmpl)
	assert.Equal(t, id, tmpl.ID)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_getCourse(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service)
	path, err := r.GetRoute("getCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses/{id}", path)
	extID := aliceExtID
	tmpl := &entity.Course{
		ID:       entity.NewID(),
		TenantID: tenantAlice,
		ExtID:    &extID,
		Name:     "default-0",
		Mode:     entity.CourseInPerson,
	}
	service.EXPECT().
		GetCourse(tmpl.ID).
		Return(tmpl, nil)
	handler := getCourse(service)
	r.Handle("/v1/courses/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/courses/" + tmpl.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// presenter.Course is returned by the api (http) server
	var d *presenter.Course
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)

	assert.Equal(t, tmpl.ID, d.ID)
	assert.Equal(t, tmpl.Name, *d.Name)
	assert.Equal(t, tmpl.Mode, *d.Mode)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_deleteCourse(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service)
	path, err := r.GetRoute("deleteCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses/{id}", path)

	id := entity.NewID()
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
	r := mux.NewRouter()
	n := negroni.New()
	MakeCourseHandlers(r, *n, service)
	path, err := r.GetRoute("deleteCourse").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/courses/{id}", path)

	id := entity.NewID()
	service.EXPECT().DeleteCourse(id).Return(entity.ErrNotFound)
	handler := deleteCourse(service)
	req, _ := http.NewRequest("DELETE", "/v1/courses/"+id.String(), nil)
	r.Handle("/v1/courses/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// TODO: Test case for updating course
