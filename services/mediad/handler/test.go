/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"

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
			Total:       rand.IntN(2000),
			URL:         "https://dummy-s3-url.abovecloud9.ai/bucket/quote/quote.json",
		}
		toJ.Media = &presenter.Metadata{
			Version:     2,
			LastUpdated: "2024-12-04T15:38:43+05:30",
			Total:       rand.IntN(200),
			URL:         "https://dummy-s3-url.abovecloud9.ai/bucket/media/mediaMeta.json",
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

		var toJ presenter.Metadata

		switch mType {
		case "quote":
			toJ = presenter.Metadata{
				Version:     1,
				LastUpdated: "2024-12-03T15:38:43+05:30",
				Total:       rand.IntN(2000),
				URL:         "https://dummy-s3-url.abovecloud9.ai/bucket/quote/quote.json",
			}
		case "media":
			toJ = presenter.Metadata{
				Version:     2,
				LastUpdated: "2024-12-04T15:38:43+05:30",
				Total:       rand.IntN(200),
				URL:         "https://dummy-s3-url.abovecloud9.ai/bucket/media/mediaMeta.json",
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

// MakeTestHandlers make url handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/media/metadata", n.With(
		negroni.Wrap(getMetadata()),
	)).Methods("GET", "OPTIONS").Name("getMetadata")

	r.Handle("/v1/media/metadata/{type}", n.With(
		negroni.Wrap(getMetadataByType()),
	)).Methods("GET", "OPTIONS").Name("getMetadataByType")
}
