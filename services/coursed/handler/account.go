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
	"strconv"

	"ac9/glad/pkg/common"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"ac9/glad/usecase/account"

	"ac9/glad/services/coursed/presenter"

	"ac9/glad/entity"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func listAccounts(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading accounts"
		var data []*entity.Account
		var err error
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		at := r.URL.Query().Get(common.HttpParamType)

		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		l.Log.Debugf("Test")

		search := r.URL.Query().Get(common.HttpParamQuery)
		page, limit, err := common.HttpGetPageParams(w, r)
		if err != nil {
			return
		}

		switch {
		case search == "":
			data, err = service.ListAccounts(tenantID, page, limit, entity.AccountType(at))
		default:
			data, err = service.SearchAccounts(tenantID, search, page, limit, entity.AccountType(at))
		}

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		// TODO: For search, this count should be equal to the number of records
		// that match the given search query
		total := service.GetCount(tenantID)
		w.Header().Set(common.HttpHeaderTotalCount, strconv.Itoa(total))
		w.Header().Set(common.HttpHeaderTenantID, tenant)

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Account
		for _, d := range data {
			toJ = append(toJ, &presenter.Account{
				ID:        d.ID,
				Username:  d.Username,
				FirstName: d.FirstName,
				LastName:  d.LastName,
				Phone:     d.Phone,
				Email:     d.Email,
				Type:      d.Type,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode account"))
		}
	})
}

// func createAccount(service account.UseCase) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		errorMessage := "Error adding account"
// 		var input struct {
// 			Name    string             `json:"name"`
// 			Type    entity.AccountType `json:"type"`
// 			Content string             `json:"content"`
// 		}

// 		tenant := r.Header.Get(common.HttpHeaderTenantID)
// 		tenantID, err := id.FromString(tenant)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			_, _ = w.Write([]byte("Missing tenant ID"))
// 			return
// 		}

// 		err = json.NewDecoder(r.Body).Decode(&input)
// 		if err != nil {
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusBadRequest)
// 			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
// 			return
// 		}

// 		id, err := service.CreateAccount(
// 			tenantID,
// 			input.Name,
// 			input.Type,
// 			input.Content)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
// 			return
// 		}
// 		toJ := &presenter.Account{
// 			ID:       id,
// 			Name:     input.Name,
// 			Type:     entity.AccountText,
// 			Content:  input.Content,
// 			TenantID: tenantID,
// 		}

// 		w.Header().Set(common.HttpHeaderTenantID, tenant)
// 		w.WriteHeader(http.StatusCreated)
// 		if err := json.NewEncoder(w).Encode(toJ); err != nil {
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			_, _ = w.Write([]byte(errorMessage))
// 			return
// 		}
// 	})
// }

func getAccount(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenant := r.Header.Get(common.HttpHeaderTenantID)

		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		errorMessage := "Error reading account"
		vars := mux.Vars(r)
		username := vars["username"]
		data, err := service.GetAccountByName(tenantID, username)
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Empty data returned"))
			return
		}

		toJ := &presenter.Account{}
		toJ.FromAccountEntity(data)

		w.Header().Set(common.HttpHeaderTenantID, r.Header.Get(common.HttpHeaderTenantID))
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode account"))
		}
	})
}

func deleteAccount(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenant := r.Header.Get(common.HttpHeaderTenantID)

		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		errorMessage := "Error removing account"
		vars := mux.Vars(r)
		username := vars["username"]
		err = service.DeleteAccountByName(tenantID, username)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case glad.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Account doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateAccount(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating account"

		vars := mux.Vars(r)
		username := vars["username"]

		var input entity.Account
		// tenant := r.Header.Get(common.HttpHeaderTenantID)
		// tenantID, err := id.FromString(tenant)
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	_, _ = w.Write([]byte("Missing tenant ID"))
		// 	return
		// }

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		input.Username = username
		err = service.UpdateAccount(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Account{
			ID:        input.ID,
			Username:  input.Username,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Phone:     input.Phone,
			Email:     input.Email,
			Type:      input.Type,
		}

		w.Header().Set(common.HttpHeaderTenantID, "")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeAccountHandlers make url handlers
func MakeAccountHandlers(r *mux.Router, n negroni.Negroni, service account.UseCase) {
	r.Handle("/v1/accounts", n.With(
		negroni.Wrap(listAccounts(service)),
	)).Methods("GET", "OPTIONS").Name("listAccounts")

	// r.Handle("/v1/accounts", n.With(
	// 	negroni.Wrap(createAccount(service)),
	// )).Methods("POST", "OPTIONS").Name("createAccount")

	r.Handle("/v1/accounts/{username}", n.With(
		negroni.Wrap(getAccount(service)),
	)).Methods("GET", "OPTIONS").Name("getAccount")

	r.Handle("/v1/accounts/{username}", n.With(
		negroni.Wrap(deleteAccount(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteAccount")

	r.Handle("/v1/accounts/{username}", n.With(
		negroni.Wrap(updateAccount(service)),
	)).Methods("PUT", "OPTIONS").Name("updateAccount")
}
