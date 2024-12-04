// Package handler implements HTTP handlers for live-darshan endpoints
package handler

import (
	"encoding/json"
	"net/http"

	"ac9/glad/services/ldsd/presenter"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func getLiveDarshanToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signature := presenter.ZoomSignature{
			Signature: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZGtLZXkiOiJhYmMxMjMiLCJtbiI6IjEyMzQ1Njc4OSIsInJvbGUiOjAsImlhdCI6MTY0NjkzNzU1MywiZXhwIjoxNjQ2OTQ0NzUzLCJhcHBLZXkiOiJhYmMxMjMiLCJ0b2tlbkV4cCI6MTY0Njk0NDc1M30.UcWxbWY-y22wFarBBc9i3lGQuZAsuUpl8GRR8wUah2M",
			FirstName: "AboveCloud9",
			LastName:  "AI",
		}

		metadataWrap := presenter.MetadataWrapper{
			Metadata: presenter.Metadata{Zoom: signature},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(metadataWrap); err != nil {
			http.Error(w, "Unable to encode metadata", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func getLiveDarshanDetails() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zoomDetails := presenter.ZoomDetails{
			ID:         10000000,
			Date:       "2024-12-04",
			StartTime:  "15:04:00", 
			MeetingID:  "1234567890",
			Password:   "test-password",
			MeetingURL: "https://zoom.us/j/5551112222",
		}

		zoomWrap := []presenter.ZoomWrapper{{Zoom: zoomDetails}}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(zoomWrap); err != nil {
			http.Error(w, "Unable to encode details", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

// MakeTestHandlers sets up URL handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/live-darshan", n.With(
		negroni.Wrap(getLiveDarshanDetails()),
	)).Methods(http.MethodGet).Name("live-darshan")

	r.Handle("/v1/live-darshan/token", n.With(
		negroni.Wrap(getLiveDarshanToken()),
	)).Methods(http.MethodGet).Name("live-darshan-token")
}
