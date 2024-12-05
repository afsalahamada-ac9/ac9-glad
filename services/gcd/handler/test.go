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
		auth := presenter.Auth{
			ClientID:     "abcd567efghijkl",
			ClientSecret: "abcd567efghijklabcd567efghijkl",
			Domain:       "http://auth.ac9ai.com",
			Region:       "us-east-2",
			UserPoolID:   "JMzj123s",
			URL:          "http://ac9.ai.com/", // url not defined in spec
		}

		config := presenter.Config{
			Version:  1,
			Timezone: []string{"EST", "Americas/New York", "CST", "MST", "PST", "Americas/Los Angeles", "HST", "AST"},
			Auth:     auth,
		}

		if err := json.NewEncoder(w).Encode(config); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode glad configuration"))
		}

		w.WriteHeader(http.StatusOK)
	})
}

// MakeTestHandlers make glad handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	// Note: Deprecated. This path was mentioned in the MSC that's frozen.
	// To avoid work at client, adding this duplicate path support that should be
	// removed once client supports config.
	r.Handle("/v1/glad/info", n.With(
		negroni.Wrap(getConfig()),
	)).Methods("GET").Name("getConfig")

	r.Handle("/v1/glad/config", n.With(
		negroni.Wrap(getConfig()),
	)).Methods("GET").Name("getConfig")
}
