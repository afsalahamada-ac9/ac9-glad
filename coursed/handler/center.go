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

	"sudhagar/glad/pkg/common"
	"sudhagar/glad/usecase/center"

	"sudhagar/glad/coursed/presenter"

	"sudhagar/glad/entity"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// TODO:
// 	- Implement pagination for center listing/search
// 	- JSON based search and formatting requires some work
// 	- ENUM can be optimized by storing integer value in the mapping
// 	- Support for location and geolocation

func listCenters(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading centers"
		var data []*entity.Center
		var err error
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		search := r.URL.Query().Get(httpParamQuery)
		page, _ := strconv.Atoi(r.URL.Query().Get(httpParamPage))
		limit, _ := strconv.Atoi(r.URL.Query().Get(httpParamLimit))

		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		switch {
		case search == "":
			data, err = service.ListCenters(tenantID, page, limit)
		default:
			// TODO: search need to be reworked; need to add a count
			// for search; also need to see how the caller generates
			// the search query request
			data, err = service.SearchCenters(tenantID, search, page, limit)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(httpHeaderTotalCount, strconv.Itoa(total))

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Center
		for _, d := range data {
			toJ = append(toJ, &presenter.Center{
				ID:      d.ID,
				Name:    d.Name,
				Mode:    d.Mode,
				ExtName: d.ExtName,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.Header().Set(common.HttpHeaderTenantID, tenant)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode center"))
		}
	})
}

func createCenter(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding center"
		var input struct {
			ExtID     string            `json:"extId"`
			ExtName   string            `json:"extName"`
			Name      string            `json:"name"`
			Mode      entity.CenterMode `json:"mode"`
			IsEnabled bool              `json:"isEnabled"`
		}

		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		id, err := service.CreateCenter(
			tenantID,
			input.ExtID,
			input.ExtName,
			input.Name,
			input.Mode,
			input.IsEnabled)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.Center{
			ID:   id,
			Name: input.Name,
			Mode: input.Mode,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func getCenter(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading center"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetCenter(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Empty data returned"))
			return
		}

		toJ := &presenter.Center{
			ID:      data.ID,
			Name:    data.Name,
			Mode:    data.Mode,
			ExtName: data.ExtName,
		}

		w.Header().Set(common.HttpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode center"))
		}
	})
}

func deleteCenter(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing center"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteCenter(id)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case entity.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Center doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateCenter(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating center"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input entity.Center
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		input.ID = id
		input.TenantID = tenantID
		err = service.UpdateCenter(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Center{
			ID:   input.ID,
			Name: input.Name,
			Mode: input.Mode,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeCenterHandlers make url handlers
func MakeCenterHandlers(r *mux.Router, n negroni.Negroni, service center.UseCase) {
	r.Handle("/v1/centers", n.With(
		negroni.Wrap(listCenters(service)),
	)).Methods("GET", "OPTIONS").Name("listCenters")

	r.Handle("/v1/centers", n.With(
		negroni.Wrap(createCenter(service)),
	)).Methods("POST", "OPTIONS").Name("createCenter")

	r.Handle("/v1/centers/{id}", n.With(
		negroni.Wrap(getCenter(service)),
	)).Methods("GET", "OPTIONS").Name("getCenter")

	r.Handle("/v1/centers/{id}", n.With(
		negroni.Wrap(deleteCenter(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteCenter")

	r.Handle("/v1/centers/{id}", n.With(
		negroni.Wrap(updateCenter(service)),
	)).Methods("PUT", "OPTIONS").Name("updateCenter")
}
