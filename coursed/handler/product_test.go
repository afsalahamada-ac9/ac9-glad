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

	"sudhagar/glad/coursed/presenter"
	"sudhagar/glad/entity"
	"sudhagar/glad/pkg/common"

	mock "sudhagar/glad/usecase/product/mock"

	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// TODO: Add test cases to test page and limit functionality
func Test_listProducts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("listProducts").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/products", path)

	tmpl := &entity.Product{
		ID:           entity.NewID(),
		TenantID:     tenantAlice,
		ExtID:        aliceExtID,
		ExtName:      "product-1",
		Title:        "Product One",
		CType:        "TYPE-1",
		DurationDays: 7,
		Visibility:   entity.ProductVisibilityPublic,
		MaxAttendees: 100,
		Format:       entity.ProductFormatInPerson,
	}
	service.EXPECT().GetCount(tmpl.TenantID).Return(1)
	service.EXPECT().
		ListProducts(tmpl.TenantID, gomock.Any(), gomock.Any()).
		Return([]*entity.Product{tmpl}, nil)

	ts := httptest.NewServer(listProducts(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listProducts_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listProducts(service))
	defer ts.Close()
	tenantID := tenantAlice
	service.EXPECT().GetCount(tenantID).Return(0)
	service.EXPECT().
		SearchProducts(tenantID, "non-existent", gomock.Any(), gomock.Any()).
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

func Test_listProducts_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	tmpl := &entity.Product{
		ID:           entity.NewID(),
		TenantID:     tenantAlice,
		ExtName:      "product-1",
		Title:        "Product One",
		CType:        "TYPE-1",
		DurationDays: 7,
		Visibility:   entity.ProductVisibilityPublic,
		MaxAttendees: 100,
		Format:       entity.ProductFormatInPerson,
	}
	service.EXPECT().GetCount(tmpl.TenantID).Return(1)
	service.EXPECT().
		SearchProducts(tmpl.TenantID, "product", gomock.Any(), gomock.Any()).
		Return([]*entity.Product{tmpl}, nil)
	ts := httptest.NewServer(listProducts(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet,
		ts.URL+"?q=product",
		nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("createProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/products", path)

	id := entity.NewID()
	service.EXPECT().
		CreateProduct(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).
		Return(id, nil)
	h := createProduct(service)

	ts := httptest.NewServer(h)
	defer ts.Close()

	payload := struct {
		ExtID            string                   `json:"extId"`
		ExtName          string                   `json:"Extname"`
		Title            string                   `json:"title"`
		CType            string                   `json:"ctype"`
		BaseProductExtID string                   `json:"baseProductExtId"`
		DurationDays     int32                    `json:"durationDays"`
		Visibility       entity.ProductVisibility `json:"visibility"`
		MaxAttendees     int32                    `json:"maxAttendees"`
		Format           entity.ProductFormat     `json:"format"`
	}{
		ExtID:            aliceExtID,
		ExtName:          "product-1",
		Title:            "Product One",
		CType:            "TYPE-1",
		BaseProductExtID: "BASE-1",
		DurationDays:     7,
		Visibility:       entity.ProductVisibilityPublic,
		MaxAttendees:     100,
		Format:           entity.ProductFormatInPerson,
	}
	payloadBytes, err := json.Marshal(payload)
	assert.Nil(t, err)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost,
		ts.URL+"/v1/products",
		bytes.NewReader(payloadBytes))
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var tmpl *presenter.Product
	json.NewDecoder(res.Body).Decode(&tmpl)
	assert.Equal(t, id, tmpl.ID)
	assert.Equal(t, payload.ExtName, tmpl.ExtName)
	assert.Equal(t, payload.Title, tmpl.Title)
	assert.Equal(t, payload.CType, tmpl.CType)
	assert.Equal(t, payload.BaseProductExtID, tmpl.BaseProductExtID)
	assert.Equal(t, payload.DurationDays, tmpl.DurationDays)
	assert.Equal(t, payload.Visibility, tmpl.Visibility)
	assert.Equal(t, payload.MaxAttendees, tmpl.MaxAttendees)
	assert.Equal(t, payload.Format, tmpl.Format)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_getProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("getProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/products/{id}", path)

	tmpl := &entity.Product{
		ID:           entity.NewID(),
		TenantID:     tenantAlice,
		ExtID:        aliceExtID,
		ExtName:      "product-1",
		Title:        "Product One",
		CType:        "TYPE-1",
		DurationDays: 7,
		Visibility:   entity.ProductVisibilityPublic,
		MaxAttendees: 100,
		Format:       entity.ProductFormatInPerson,
	}
	service.EXPECT().
		GetProduct(tmpl.ID).
		Return(tmpl, nil)
	handler := getProduct(service)
	r.Handle("/v1/products/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/products/" + tmpl.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var d *presenter.Product
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, tmpl.ID, d.ID)
	assert.Equal(t, tmpl.ExtName, d.ExtName)
	assert.Equal(t, tmpl.Title, d.Title)
	assert.Equal(t, tmpl.CType, d.CType)
	assert.Equal(t, tmpl.DurationDays, d.DurationDays)
	assert.Equal(t, tmpl.Visibility, d.Visibility)
	assert.Equal(t, tmpl.MaxAttendees, d.MaxAttendees)
	assert.Equal(t, tmpl.Format, d.Format)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_deleteProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("deleteProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/products/{id}", path)

	id := entity.NewID()
	service.EXPECT().DeleteProduct(id).Return(nil)
	handler := deleteProduct(service)
	req, _ := http.NewRequest("DELETE", "/v1/products/"+id.String(), nil)
	r.Handle("/v1/products/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_deleteProductNonExistent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("deleteProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/products/{id}", path)

	id := entity.NewID()
	service.EXPECT().DeleteProduct(id).Return(entity.ErrNotFound)
	handler := deleteProduct(service)
	req, _ := http.NewRequest("DELETE", "/v1/products/"+id.String(), nil)
	r.Handle("/v1/products/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func Test_updateProduct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)
	path, err := r.GetRoute("updateProduct").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/products/{id}", path)

	id := entity.NewID()
	updatePayload := &entity.Product{
		ID:           id,
		TenantID:     tenantAlice,
		ExtID:        aliceExtID,
		ExtName:      "updated-product",
		Title:        "Updated Product",
		CType:        "TYPE-2",
		DurationDays: 14,
		Visibility:   entity.ProductVisibilityUnlisted,
		MaxAttendees: 200,
		Format:       entity.ProductFormatOnline,
	}

	service.EXPECT().
		UpdateProduct(gomock.Any()).
		Return(nil)

	handler := updateProduct(service)
	payloadBytes, err := json.Marshal(updatePayload)
	assert.Nil(t, err)

	req, _ := http.NewRequest("PUT",
		"/v1/products/"+id.String(),
		bytes.NewReader(payloadBytes))
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	req.Header.Set("Content-Type", "application/json")

	r.Handle("/v1/products/{id}", handler).Methods("PUT", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response presenter.Product
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.Nil(t, err)

	// Verify the response contains the updated values
	assert.Equal(t, id, response.ID)
	assert.Equal(t, updatePayload.ExtName, response.ExtName)
	assert.Equal(t, updatePayload.Title, response.Title)
	assert.Equal(t, updatePayload.CType, response.CType)
	assert.Equal(t, updatePayload.DurationDays, response.DurationDays)
	assert.Equal(t, updatePayload.Visibility, response.Visibility)
	assert.Equal(t, updatePayload.MaxAttendees, response.MaxAttendees)
	assert.Equal(t, updatePayload.Format, response.Format)
	assert.Equal(t, tenantAlice.String(), rr.Header().Get(common.HttpHeaderTenantID))
}

func Test_updateProduct_BadRequest(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)

	id := entity.NewID()
	handler := updateProduct(service)

	// Test with invalid JSON payload
	req, _ := http.NewRequest("PUT",
		"/v1/products/"+id.String(),
		bytes.NewReader([]byte("invalid json")))
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	req.Header.Set("Content-Type", "application/json")

	r.Handle("/v1/products/{id}", handler).Methods("PUT", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_updateProduct_MissingTenant(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeProductHandlers(r, *n, service)

	id := entity.NewID()
	updatePayload := &entity.Product{
		ID:      id,
		ExtName: "updated-product",
	}

	handler := updateProduct(service)
	payloadBytes, err := json.Marshal(updatePayload)
	assert.Nil(t, err)

	req, _ := http.NewRequest("PUT",
		"/v1/products/"+id.String(),
		bytes.NewReader(payloadBytes))
	// Intentionally not setting tenant ID header
	req.Header.Set("Content-Type", "application/json")

	r.Handle("/v1/products/{id}", handler).Methods("PUT", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
