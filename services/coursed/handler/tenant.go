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
	"ac9/glad/usecase/tenant"

	"ac9/glad/services/coursed/presenter"

	"ac9/glad/entity"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// TODO: Add pagination support
func listTenants(service tenant.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading tenants"
		var data []*entity.Tenant
		var err error

		page, limit, err := common.HttpGetPageParams(w, r)
		if err != nil {
			return
		}

		data, err = service.ListTenants(page, limit)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount()
		w.Header().Set(common.HttpHeaderTotalCount, strconv.Itoa(total))

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Tenant
		for _, d := range data {
			toJ = append(toJ, &presenter.Tenant{
				ID:        d.ID,
				Name:      d.Name,
				Country:   d.Country,
				AuthToken: d.AuthToken,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode tenant"))
		}
	})
}

func createTenant(service tenant.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding tenant"
		var input struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		tenantID, err := service.CreateTenant(input.Name, input.Country)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.Tenant{
			ID:      tenantID,
			Name:    input.Name,
			Country: input.Country,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func getTenant(service tenant.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading tenant"
		vars := mux.Vars(r)
		tenantID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetTenant(tenantID)
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

		toJ := &presenter.Tenant{
			ID:        data.ID,
			Name:      data.Name,
			Country:   data.Country,
			AuthToken: data.AuthToken,
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode tenant"))
		}
	})
}

func deleteTenant(service tenant.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing tenant"
		vars := mux.Vars(r)
		tenantID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteTenant(tenantID)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case glad.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Tenant doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func login(service tenant.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading tenant"
		var input struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		data, err := service.Login(input.Name, input.Country)
		switch err {
		case nil:
			break
		// intentionally returning same response for auth failure and not found scenarios
		case glad.ErrAuthFailure, glad.ErrNotFound:
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Invalid login credentials"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Tenant{
			ID:        data.ID,
			Name:      data.Name,
			AuthToken: data.AuthToken,
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode tenant"))
		}
	})
}

func updateTenant(service tenant.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating tenant"

		vars := mux.Vars(r)
		tenantID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input entity.Tenant
		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		err = service.UpdateTenant(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		input.ID = tenantID
		toJ := &presenter.Tenant{
			ID:      input.ID,
			Name:    input.Name,
			Country: input.Country,
			// AuthToken is not returned back
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeTenantHandlers make url handlers
func MakeTenantHandlers(r *mux.Router, n negroni.Negroni, service tenant.UseCase) {
	r.Handle("/v1/tenants", n.With(
		negroni.Wrap(listTenants(service)),
	)).Methods("GET", "OPTIONS").Name("listTenants")

	r.Handle("/v1/tenants", n.With(
		negroni.Wrap(createTenant(service)),
	)).Methods("POST", "OPTIONS").Name("createTenant")

	r.Handle("/v1/tenants/{id}", n.With(
		negroni.Wrap(getTenant(service)),
	)).Methods("GET", "OPTIONS").Name("getTenant")

	r.Handle("/v1/tenants/{id}", n.With(
		negroni.Wrap(deleteTenant(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteTenant")

	r.Handle("/v1/login", n.With(
		negroni.Wrap(login(service)),
	)).Methods("POST", "OPTIONS").Name("login")

	r.Handle("/v1/tenants/{id}", n.With(
		negroni.Wrap(updateTenant(service)),
	)).Methods("PUT", "OPTIONS").Name("updateTenant")
}
