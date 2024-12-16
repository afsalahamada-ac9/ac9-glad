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
	"ac9/glad/usecase/center"

	"ac9/glad/services/coursed/presenter"

	"ac9/glad/entity"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// TODO:
// 	- JSON based search and formatting requires some work
// 	- Support for location and geolocation

func listCenters(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading centers"
		var centers []*entity.Center
		var err error
		tenant := r.Header.Get(common.HttpHeaderTenantID)

		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		search := r.URL.Query().Get(common.HttpParamQuery)
		page, limit, err := common.HttpGetPageParams(w, r)
		if err != nil {
			return
		}

		switch {
		case search == "":
			centers, err = service.ListCenters(tenantID, page, limit)
		default:
			// TODO: search need to be reworked; need to add a count
			// for search; also need to see how the caller generates
			// the search query request
			centers, err = service.SearchCenters(tenantID, search, page, limit)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(common.HttpHeaderTotalCount, strconv.Itoa(total))

		if centers == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		var responses []*presenter.Center
		for _, center := range centers {
			resp := &presenter.Center{}
			resp.FromEntityCenter(center)
			l.Log.Debugf("Entity center=%v, Center=%v", center, resp)

			responses = append(responses, resp)
		}
		if err := json.NewEncoder(w).Encode(responses); err != nil {
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
			Name      string            `json:"name"`
			Mode      entity.CenterMode `json:"mode"`
			IsEnabled bool              `json:"isEnabled"`
		}

		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
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

		centerID, err := service.CreateCenter(
			tenantID,
			input.Name,
			input.Mode,
			input.IsEnabled)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		resp := &presenter.CenterResponse{
			ID: centerID,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
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
		centerID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		center, err := service.GetCenter(centerID)
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if center == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Center not found"))
			return
		}

		resp := &presenter.Center{}
		resp.FromEntityCenter(center)

		w.Header().Set(common.HttpHeaderTenantID, center.TenantID.String())
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode center"))
		}
	})
}

func deleteCenter(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing center"
		vars := mux.Vars(r)
		id, err := id.FromString(vars["id"])
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
		case glad.ErrNotFound:
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
		centerID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input entity.Center
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
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

		input.ID = centerID
		input.TenantID = tenantID
		err = service.UpdateCenter(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
	})
}

func importCenter(service center.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error importing centers"

		var gCenters []glad.Center
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&gCenters)
		if err != nil {
			l.Log.Warnf("Unable to decode object. err=%v", err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		var response []*presenter.CenterImportResponse
		for _, gCenter := range gCenters {
			center := &entity.Center{}
			presenter.GladCenterToEntity(gCenter, center)
			center.TenantID = tenantID

			// TODO: optimize DB operations by doing multiple inserts simultaneously
			centerID, err := service.UpsertCenter(center)
			if err != nil {
				l.Log.Warnf("Unable to upsert center extID=%v, err=%v", center.ExtID, err)
			}

			response = append(response, &presenter.CenterImportResponse{
				ID:      centerID,
				ExtID:   center.ExtID,
				IsError: err != nil,
			})
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
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

	r.Handle("/v1/centers/import", n.With(
		negroni.Wrap(importCenter(service)),
	)).Methods("POST", "OPTIONS").Name("importCenter")

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
