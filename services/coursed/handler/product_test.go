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

	mock "ac9/glad/usecase/product/mock"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
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

	product := &entity.Product{
		ID:           id.New(),
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
	service.EXPECT().GetCount(product.TenantID).Return(1)
	service.EXPECT().
		ListProducts(product.TenantID, gomock.Any(), gomock.Any()).
		Return([]*entity.Product{product}, nil)

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
	product := &entity.Product{
		ID:           id.New(),
		TenantID:     tenantAlice,
		ExtName:      "product-1",
		Title:        "Product One",
		CType:        "TYPE-1",
		DurationDays: 7,
		Visibility:   entity.ProductVisibilityPublic,
		MaxAttendees: 100,
		Format:       entity.ProductFormatInPerson,
	}
	service.EXPECT().GetCount(product.TenantID).Return(1)
	service.EXPECT().
		SearchProducts(product.TenantID, "product", gomock.Any(), gomock.Any()).
		Return([]*entity.Product{product}, nil)
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

	productID := id.New()
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
		Return(productID, nil)
	h := createProduct(service)

	ts := httptest.NewServer(h)
	defer ts.Close()

	payload := struct {
		ExtID            string                   `json:"extID"`
		ExtName          string                   `json:"extName"`
		Title            string                   `json:"title"`
		CType            string                   `json:"ctype"`
		BaseProductExtID string                   `json:"baseProductExtID"`
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

	var product *presenter.Product
	json.NewDecoder(res.Body).Decode(&product)
	assert.Equal(t, productID, product.ID)
	assert.Equal(t, payload.ExtName, product.ExtName)
	assert.Equal(t, payload.Title, product.Title)
	assert.Equal(t, payload.CType, product.CType)
	assert.Equal(t, payload.BaseProductExtID, product.BaseProductExtID)
	assert.Equal(t, payload.DurationDays, product.DurationDays)
	assert.Equal(t, payload.Visibility, product.Visibility)
	assert.Equal(t, payload.MaxAttendees, product.MaxAttendees)
	assert.Equal(t, payload.Format, product.Format)
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

	product := &entity.Product{
		ID:           id.New(),
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
		GetProduct(product.ID).
		Return(product, nil)
	handler := getProduct(service)
	r.Handle("/v1/products/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/products/" + product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var d *presenter.Product
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, product.ID, d.ID)
	assert.Equal(t, product.ExtName, d.ExtName)
	assert.Equal(t, product.Title, d.Title)
	assert.Equal(t, product.CType, d.CType)
	assert.Equal(t, product.DurationDays, d.DurationDays)
	assert.Equal(t, product.Visibility, d.Visibility)
	assert.Equal(t, product.MaxAttendees, d.MaxAttendees)
	assert.Equal(t, product.Format, d.Format)
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

	id := id.New()
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

	id := id.New()
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

	id := id.New()
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

	id := id.New()
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

	id := id.New()
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
