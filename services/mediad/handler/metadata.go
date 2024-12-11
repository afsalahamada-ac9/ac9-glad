/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"encoding/json"
	"net/http"

	"ac9/glad/services/mediad/entity"
	"ac9/glad/services/mediad/presenter"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func getMetadata() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var toJ presenter.MetadataResponse

		toJ.Quote = &presenter.Metadata{
			Version:     1,
			LastUpdated: "2024-12-03T15:38:43+05:30",
			Total:       1240,
			URL:         "https://content-management-service.s3.us-east-1.amazonaws.com/quotes/quotes.json",
		}
		toJ.Media = &presenter.Metadata{
			Version:     1,
			LastUpdated: "2024-12-04T15:38:43+05:30",
			Total:       75,
			URL:         "https://content-management-service.s3.us-east-1.amazonaws.com/media/media.json",
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode metadata"))
		}

		w.WriteHeader(http.StatusOK)
	})
}

func getMetadataByType() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		mType := vars["type"]

		if mType == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing metadata type"))
		}

		var toJ presenter.MetadataResponse

		switch entity.ContentType(mType) {
		case entity.MediaQuote:
			toJ.Quote = &presenter.Metadata{
				Version:     1,
				LastUpdated: "2024-12-03T15:38:43+05:30",
				Total:       1240,
				URL:         "https://content-management-service.s3.us-east-1.amazonaws.com/quotes/quotes.json",
			}
		case entity.MediaImage:
			toJ.Media = &presenter.Metadata{
				Version:     1,
				LastUpdated: "2024-12-04T15:38:43+05:30",
				Total:       75,
				URL:         "https://content-management-service.s3.us-east-1.amazonaws.com/media/media.json",
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Invalid media metadata type"))
			return
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode metadata"))
		}

		w.WriteHeader(http.StatusOK)
	})
}

// MakeMetadataHandlers make media handlers
func MakeMetadataHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/media/metadata", n.With(
		negroni.Wrap(getMetadata()),
	)).Methods("GET", "OPTIONS").Name("getMetadata")

	r.Handle("/v1/media/metadata/{type}", n.With(
		negroni.Wrap(getMetadataByType()),
	)).Methods("GET", "OPTIONS").Name("getMetadataByType")
}
