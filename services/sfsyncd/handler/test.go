/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"net/http"

	"ac9/glad/config"
	l "ac9/glad/pkg/logger"
	"ac9/glad/pkg/util"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/valyala/fasthttp"
)

func getProducts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error retrieving products"

		// Prepare a client
		c := &fasthttp.Client{}

		fullPath := "http://" +
			util.GetStrEnvOrConfig("COURSED_ADDR", config.COURSED_ADDR) +
			"/v1/products"
		l.Log.Infof("fullPath=%v", fullPath)
		statusCode, body, err := c.Get(nil, fullPath)

		if err != nil {
			l.Log.Errorf("%v. err=%v", errorMessage, err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Get() coursed products error"))
			return
		}

		if statusCode != fasthttp.StatusOK {
			w.WriteHeader(statusCode)
			return
		}

		// var toJ []*presenter.Tenant
		// for _, d := range data {
		// 	toJ = append(toJ, &presenter.Tenant{
		// 		ID:        d.ID,
		// 		Name:      d.Name,
		// 		Country:   d.Country,
		// 		AuthToken: d.AuthToken,
		// 	})
		// }
		// if err := json.NewEncoder(w).Encode(toJ); err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	_, _ = w.Write([]byte("Unable to encode products"))
		// }

		w.Write(body)
		w.WriteHeader(statusCode)
	})
}

// MakeTestHandlers make url handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/products", n.With(
		negroni.Wrap(getProducts()),
	)).Methods("GET", "OPTIONS").Name("getProducts")
}
