/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"ac9/glad/entity"
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/id"
	"ac9/glad/pkg/logger"
	"ac9/glad/services/coursed/presenter"

	mock "ac9/glad/usecase/account/mock"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
)

const (
	accountIDPrimary   id.ID = 13790492210917010000
	accountIDSecondary id.ID = 13790492210917010002

	accountUsernamePrimary   string = "12345550001"
	accountUsernameSecondary string = "12345550002"
)

func TestMain(m *testing.M) {
	// Initialize logger
	Log := logger.NewLoggerZap()
	if Log == nil {
		log.Fatalf("Failed to initialize logger")
	}

	// Run tests
	code := m.Run()

	// Exit with the result code
	os.Exit(code)
}

func Test_listAccounts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAccountHandlers(r, *n, service)
	path, err := r.GetRoute("listAccounts").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/accounts", path)
	account := &entity.Account{
		ID:       accountIDPrimary,
		ExtID:    aliceExtID,
		Username: accountUsernamePrimary,
		Type:     entity.AccountTeacher,
	}
	service.EXPECT().GetCount(tenantAlice).Return(1)
	service.EXPECT().
		ListAccounts(tenantAlice, gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]*entity.Account{account}, nil)
	ts := httptest.NewServer(listAccounts(service))
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

// func Test_createAccount(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	service := mock.NewMockUseCase(controller)
// 	r := mux.NewRouter()
// 	n := negroni.New()
// 	MakeAccountHandlers(r, *n, service)
// 	path, err := r.GetRoute("createAccount").GetPathTemplate()
// 	assert.Nil(t, err)
// 	assert.Equal(t, "/v1/accounts", path)

// 	id := id.New()
// 	service.EXPECT().
// 		CreateAccount(gomock.Any(),
// 			gomock.Any(),
// 			gomock.Any(),
// 			gomock.Any()).
// 		Return(id, nil)
// 	h := createAccount(service)

// 	ts := httptest.NewServer(h)
// 	defer ts.Close()

// 	payload := struct {
// 		TenantID id.ID          `json:"tenant_id"`
// 		Username     string             `json:"Username"`
// 		Type     entity.AccountType `json:"type"`
// 		Content  string             `json:"content"`
// 	}{TenantID: tenantAlice,
// 		Username:    "default-0",
// 		Type:    (entity.AccountText),
// 		Content: "This is a default message"}
// 	payloadBytes, err := json.Marshal(payload)
// 	assert.Nil(t, err)

// 	client := &http.Client{}
// 	req, _ := http.NewRequest(http.MethodPost,
// 		ts.URL+"/v1/accounts",
// 		bytes.NewReader(payloadBytes))
// 	req.Header.Set(common.HttpHeaderTenantID, tenantAlice)
// 	req.Header.Set("Content-Type", "application/json")
// 	res, err := client.Do(req)

// 	assert.Nil(t, err)
// 	assert.Equal(t, http.StatusCreated, res.StatusCode)

// 	var account *presenter.Account
// 	json.NewDecoder(res.Body).Decode(&account)
// 	assert.Equal(t, id, account.ID)
// 	assert.Equal(t, payload.Content, account.Content)
// 	assert.Equal(t, payload.Username, account.Username)
// 	assert.Equal(t, payload.Type, account.Type)
// 	assert.Equal(t, tenantAlice, res.Header.Get(common.HttpHeaderTenantID))
// }

func Test_getAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAccountHandlers(r, *n, service)
	path, err := r.GetRoute("getAccount").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/accounts/{username}", path)
	account := &entity.Account{
		ID:       accountIDPrimary,
		TenantID: tenantAlice,
		ExtID:    aliceExtID,
		Username: accountUsernamePrimary,
		Type:     entity.AccountTeacher,
	}
	service.EXPECT().
		GetAccountByName(account.TenantID, account.Username).
		Return(account, nil)

	handler := getAccount(service)
	r.Handle("/v1/accounts/{username}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/v1/accounts/"+account.Username, nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	res, err := client.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// presenter.Account is returned by the api (http) server
	var d *presenter.Account
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)

	assert.Equal(t, account.ID, d.ID)
	assert.Equal(t, account.Username, d.Username)
	assert.Equal(t, account.Type, d.Type)
	assert.Equal(t, tenantAlice.String(), res.Header.Get(common.HttpHeaderTenantID))
}

func Test_deleteAccount(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAccountHandlers(r, *n, service)
	path, err := r.GetRoute("deleteAccount").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/accounts/{username}", path)

	username := accountUsernamePrimary
	service.EXPECT().DeleteAccountByName(tenantAlice, username).Return(nil)
	handler := deleteAccount(service)
	req, _ := http.NewRequest("DELETE", "/v1/accounts/"+username, nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	r.Handle("/v1/accounts/{username}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_deleteAccountNonExistent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeAccountHandlers(r, *n, service)
	path, err := r.GetRoute("deleteAccount").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/accounts/{username}", path)

	username := accountUsernamePrimary
	service.EXPECT().DeleteAccountByName(tenantAlice, username).Return(entity.ErrNotFound)
	handler := deleteAccount(service)
	req, _ := http.NewRequest("DELETE", "/v1/accounts/"+username, nil)
	req.Header.Set(common.HttpHeaderTenantID, tenantAlice.String())
	r.Handle("/v1/accounts/{username}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// TODO: Test case for updating account
