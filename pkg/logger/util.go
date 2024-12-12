/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package logger

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func setLogLevel(logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Unable to set log level"
		vars := mux.Vars(r)

		err := logger.SetLevel(vars["level"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// MakeLogHandlers make log handlers
func MakeLogHandlers(r *mux.Router, n negroni.Negroni, serviceName string, logger Logger) {
	r.Handle("/v1/"+serviceName+"/log/{level}", n.With(
		negroni.Wrap(setLogLevel(logger)),
	)).Methods("POST", "PUT").Name("setLogLevel")
}
