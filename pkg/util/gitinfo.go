/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package util

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	buildTime string
	gitHash   string
	version   string
)

// ServiceInfo data
type ServiceInfo struct {
	BuildTime string `json:"buildTime,omitempty"`
	GitHash   string `json:"gitHash,omitempty"`
	Name      string `json:"name"`
	Version   string `json:"version"`
}

func getInfo(serviceName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Unable to get service information"

		jsonResponse := &ServiceInfo{
			BuildTime: buildTime,
			GitHash:   gitHash,
			Name:      serviceName,
			Version:   version,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeInfoHandlers make info handlers
func MakeInfoHandlers(r *mux.Router, n negroni.Negroni, serviceName string) {
	r.Handle("/v1/"+serviceName+"/info", n.With(
		negroni.Wrap(getInfo(serviceName)),
	)).Methods(http.MethodGet, http.MethodOptions).Name("getInfo")
}
