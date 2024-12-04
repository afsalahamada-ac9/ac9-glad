/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */
// Package handler implements HTTP handlers for live-darshan endpoints
package handler

import (
	"encoding/json"
	"net/http"

	"ac9/glad/services/ldsd/presenter"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

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

func listLiveDarshan() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ld := presenter.LiveDarshan{
			ID:         10000000,
			Date:       "2024-12-04",
			StartTime:  "15:04:00",
			MeetingID:  "1234567890",
			Password:   "test-password",
			MeetingURL: "https://zoom.us/j/5551112222",
		}

		var ldList []presenter.LiveDarshan
		ldList = append(ldList, ld)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ldList); err != nil {
			http.Error(w, "Unable to encode live darshan details", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

// MakeTestHandlers sets up live darshan handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/live-darshan", n.With(
		negroni.Wrap(listLiveDarshan()),
	)).Methods(http.MethodGet).Name("listLiveDarshan")

	r.Handle("/v1/live-darshan/config", n.With(
		negroni.Wrap(getLiveDarshanConfig()),
	)).Methods(http.MethodGet).Name("getLiveDarshanConfig")
}
