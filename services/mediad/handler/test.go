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

	l "ac9/glad/pkg/logger"
	"ac9/glad/services/mediad/presenter"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func getMetadata() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// errorMessage := "Error retrieving metadata"
		// // Prepare a client
		// c := &fasthttp.Client{}

		// fullPath := "http://" +
		// 	util.GetStrEnvOrConfig("COURSED_ADDR", config.COURSED_ADDR) +
		// 	"/v1/products"
		// l.Log.Infof("fullPath=%v", fullPath)
		// statusCode, body, err := c.Get(nil, fullPath)

		// if err != nil {
		// 	l.Log.Errorf("%v. err=%v", errorMessage, err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	_, _ = w.Write([]byte("Get() coursed products error"))
		// 	return
		// }

		// if statusCode != fasthttp.StatusOK {
		// 	w.WriteHeader(statusCode)
		// 	return
		// }

		vars := mux.Vars(r)
		mType := vars["type"]

		var mTypes []string
		if mType == "" {
			mTypes = append(mTypes, "quote")
			mTypes = append(mTypes, "media")
		} else {
			mTypes = append(mTypes, mType)
		}

		var toJ presenter.MetadataResponse

		l.Log.Debugf("mTypes=%v", mTypes)
		for _, mType = range mTypes {
			switch mType {
			case "quote":
				toJ.Quote = &presenter.Metadata{
					Version:     1,
					LastUpdated: "2024-12-03T15:38:43+05:30",
					Total:       rand.IntN(2000),
					URL:         "https://dummy-s3-url.abovecloud9.ai/bucket/quote/quote.json",
				}
			case "media":
				toJ.Media = &presenter.Metadata{
					Version:     2,
					LastUpdated: "2024-12-04T15:38:43+05:30",
					Total:       rand.IntN(200),
					URL:         "https://dummy-s3-url.abovecloud9.ai/bucket/media/mediaMeta.json",
				}
			}
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
		negroni.Wrap(getMetadata()),
	)).Methods("GET", "OPTIONS").Name("getMetadata")
}
