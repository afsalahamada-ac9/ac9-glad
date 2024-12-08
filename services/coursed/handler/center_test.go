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
	"ac9/glad/pkg/id"
	"ac9/glad/services/coursed/presenter"

	mock "ac9/glad/usecase/center/mock"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
)

const (
	tenantAlice       id.ID = 7264348473653242881
	tenantNonExistent id.ID = 0
	aliceExtID              = "000aliceExtID"
)

func Test_listCenters(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCenterHandlers(r, *n, service)
	path, err := r.GetRoute("listCenters").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/centers", path)
	center := &entity.Center{
		ID:       id.New(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(center.TenantID).Return(1)
	service.EXPECT().
		ListCenters(center.TenantID, gomock.Any(), gomock.Any()).
		Return([]*entity.Center{center}, nil)
	ts := httptest.NewServer(listCenters(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listCenters_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listCenters(service))
	defer ts.Close()
	tenantID := tenantAlice
	service.EXPECT().GetCount(tenantID).Return(0)
	service.EXPECT().
		SearchCenters(tenantID, "non-existent", gomock.Any(), gomock.Any()).
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

func Test_listCenters_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	center := &entity.Center{
		ID:       id.New(),
		TenantID: tenantAlice,
		Name:     "default-0",
	}
	service.EXPECT().GetCount(center.TenantID).Return(1)
	service.EXPECT().
		SearchCenters(center.TenantID, "default", gomock.Any(), gomock.Any()).
		Return([]*entity.Center{center}, nil)
	ts := httptest.NewServer(listCenters(service))
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

func Test_createCenter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCenterHandlers(r, *n, service)
	path, err := r.GetRoute("createCenter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/centers", path)

	centerID := id.New()
	service.EXPECT().
		CreateCenter(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).
		Return(centerID, nil)
	h := createCenter(service)

	ts := httptest.NewServer(h)
	defer ts.Close()

	payload := struct {
		TenantID id.ID             `json:"tenant_id"`
		ExtID    string            `json:"extId"`
		Name     string            `json:"name"`
		Mode     entity.CenterMode `json:"mode"`
		Content  string            `json:"content"`
	}{TenantID: tenantAlice,
		ExtID: aliceExtID,
		Name:  "default-0",
		Mode:  (entity.CenterInPerson)}
	payloadBytes, err := json.Marshal(payload)
	assert.Nil(t, err)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost,
		ts.URL+"/v1/centers",
		bytes.NewReader(payloadBytes))
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var center *presenter.Center
	json.NewDecoder(res.Body).Decode(&center)
	assert.Equal(t, centerID, center.ID)
	assert.Equal(t, payload.Name, center.Name)
	assert.Equal(t, payload.Mode, center.Mode)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_getCenter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCenterHandlers(r, *n, service)
	path, err := r.GetRoute("getCenter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/centers/{id}", path)
	center := &entity.Center{
		ID:       id.New(),
		TenantID: tenantAlice,
		ExtID:    aliceExtID,
		Name:     "default-0",
		Mode:     entity.CenterInPerson,
	}
	service.EXPECT().
		GetCenter(center.ID).
		Return(center, nil)
	handler := getCenter(service)
	r.Handle("/v1/centers/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/centers/" + center.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// presenter.Center is returned by the api (http) server
	var d *presenter.Center
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)

	assert.Equal(t, center.ID, d.ID)
	assert.Equal(t, center.Name, d.Name)
	assert.Equal(t, center.Mode, d.Mode)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_deleteCenter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCenterHandlers(r, *n, service)
	path, err := r.GetRoute("deleteCenter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/centers/{id}", path)

	id := id.New()
	service.EXPECT().DeleteCenter(id).Return(nil)
	handler := deleteCenter(service)
	req, _ := http.NewRequest("DELETE", "/v1/centers/"+id.String(), nil)
	r.Handle("/v1/centers/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_deleteCenterNonExistent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeCenterHandlers(r, *n, service)
	path, err := r.GetRoute("deleteCenter").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/centers/{id}", path)

	id := id.New()
	service.EXPECT().DeleteCenter(id).Return(entity.ErrNotFound)
	handler := deleteCenter(service)
	req, _ := http.NewRequest("DELETE", "/v1/centers/"+id.String(), nil)
	r.Handle("/v1/centers/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// TODO: Test case for updating center
