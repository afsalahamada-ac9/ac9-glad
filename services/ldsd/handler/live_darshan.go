/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */
// Package handler implements HTTP handlers for live-darshan endpoints
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ac9/glad/pkg/common"
	"ac9/glad/pkg/glad"
	l "ac9/glad/pkg/logger"
	"ac9/glad/services/ldsd/entity"
	"ac9/glad/services/ldsd/presenter"
	"ac9/glad/services/ldsd/usecase/live_darshan"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// TODO: integrate with zoom library to retrieve the information
func getLiveDarshanConfig() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zoomInfo := presenter.ZoomInfo{
			Signature:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZGtLZXkiOiJhYmMxMjMiLCJtbiI6IjEyMzQ1Njc4OSIsInJvbGUiOjAsImlhdCI6MTY0NjkzNzU1MywiZXhwIjoxNjQ2OTQ0NzUzLCJhcHBLZXkiOiJhYmMxMjMiLCJ0b2tlbkV4cCI6MTY0Njk0NDc1M30.UcWxbWY-y22wFarBBc9i3lGQuZAsuUpl8GRR8wUah2M",
			DisplayName: "AboveCloud9 AI",
		}

		config := presenter.LiveDarshanConfig{
			Zoom: zoomInfo,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(config); err != nil {
			http.Error(w, "Unable to encode live darshan config", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func createLiveDarshan(service live_darshan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding live darshan"
		var req presenter.LiveDarshanReq

		tenantID, err := common.HttpGetTenantID(w, r)
		if err != nil {
			l.Log.Debugf("Tenant id is missing")
			return
		}

		accountID, err := common.HttpGetAccountID(w, r)
		if err != nil {
			l.Log.Debugf("Account id is missing")
			return
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		// validation checks
		if req.Date == "" || req.StartTime == "" || req.MeetingID == "" || req.MeetingURL == "" {
			l.Log.Warnf("[%v] Mandatory fields are missing. %#v", tenantID, req)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Mandatory fields missing"))
			return
		}

		ld, err := service.CreateLiveDarshan(
			tenantID,
			req.Date,
			req.StartTime,
			req.MeetingURL,
			accountID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.LiveDarshanResponse{
			ID: ld.ID,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenantID.String())
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			l.Log.Errorf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func listLiveDarshan(service live_darshan.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error retrieving live darshan"
		var lds []*entity.LiveDarshan
		var err error

		tenantID, err := common.HttpGetTenantID(w, r)
		if err != nil {
			l.Log.Debugf("Tenant id is missing")
			return
		}

		page, limit, err := common.HttpGetPageParams(w, r)
		if err != nil {
			return
		}

		lds, err = service.ListLiveDarshan(tenantID, page, limit)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(common.HttpHeaderTotalCount, strconv.Itoa(total))

		if lds == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		var ldList []*presenter.LiveDarshan
		for _, ld := range lds {
			liveDarshan := &presenter.LiveDarshan{}
			liveDarshan.FromEntityLiveDarshan(ld)
			ldList = append(ldList, liveDarshan)
		}

		w.Header().Set(common.HttpHeaderTenantID, tenantID.String())
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ldList); err != nil {
			http.Error(w, "Unable to encode live darshan details", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func updateLiveDarshan() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// ldId := vars["id"]

		w.WriteHeader(http.StatusOK)
	})
}

func deleteLiveDarshan() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)

		w.WriteHeader(http.StatusOK)
	})
}

// MakeLiveDarshanHandlers sets up live darshan handlers
func MakeLiveDarshanHandlers(r *mux.Router, n negroni.Negroni, service live_darshan.UseCase) {
	r.Handle("/v1/live-darshan/config", n.With(
		negroni.Wrap(getLiveDarshanConfig()),
	)).Methods(http.MethodGet, http.MethodOptions).Name("getLiveDarshanConfig")

	r.Handle("/v1/live-darshan", n.With(
		negroni.Wrap(createLiveDarshan(service)),
	)).Methods(http.MethodPost, http.MethodOptions).Name("createLiveDarshan")

	r.Handle("/v1/live-darshan", n.With(
		negroni.Wrap(listLiveDarshan(service)),
	)).Methods(http.MethodGet, http.MethodOptions).Name("listLiveDarshan")

	r.Handle("/v1/live-darshan/{id}", n.With(
		negroni.Wrap(updateLiveDarshan()),
	)).Methods(http.MethodPut, http.MethodOptions).Name("updateLiveDarshan")

	r.Handle("/v1/live-darshan/{id}", n.With(
		negroni.Wrap(deleteLiveDarshan()),
	)).Methods(http.MethodDelete, http.MethodOptions).Name("deleteLiveDarshan")
}
