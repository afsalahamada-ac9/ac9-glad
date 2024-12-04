/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"encoding/json"
	"net/http"

	"ac9/glad/services/gcd/presenter"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func getConfig() http.Handler {
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

		auth := presenter.Auth{
			ClientId:     "abcd567efghijkl",
			ClientSecret: "abcd567efghijklabcd567efghijkl",
			Domain:       "http://auth.ac9ai.com",
			Region:       "us-east-2",
			UserPoolId:   "JMzj123s",
			Url:          "http://ac9.ai.com/", // url not defined in spec
		}

		config := presenter.Config{
			Version:  1,
			Timezone: []string{"EST"},
			Auth:     auth,
		}

		configWrap := []presenter.Config{config}

		if err := json.NewEncoder(w).Encode(configWrap); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode metadata"))
		}

		w.WriteHeader(http.StatusOK)
	})
}

// MakeTestHandlers make url handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/glad/config", n.With(
		negroni.Wrap(getConfig()),
	)).Methods("GET").Name("glad-config")
}
