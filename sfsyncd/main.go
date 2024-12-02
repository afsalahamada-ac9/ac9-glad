/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	// Uber zap logging
	"sudhagar/glad/pkg/logger"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"sudhagar/glad/config"
	"sudhagar/glad/pkg/metric"
	"sudhagar/glad/pkg/middleware"
	"sudhagar/glad/pkg/util"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Note: Not a best practice to use global variables in general
var Log *logger.Logger

func main() {

	Log := logger.NewLoggerZap()
	if Log == nil {
		log.Fatalf("Failed to initialize logger")
	}

	// Defer the Sync to ensure logs are flushed out before program exits
	defer func() {
		if err := Log.Sync(); err != nil {
			log.Printf("Failed to sync logger: %v", err)
		}
	}()

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		Log.Fatalf("%v", err.Error())
	}
	r := mux.NewRouter()
	// handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.AddDefaultTenant),
		negroni.NewLogger(),
	)

	// log handler
	logger.MakeLogHandlers(r, *n, "sfsyncd", Log)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(util.GetIntEnvOrConfig("API_PORT", config.API_PORT)),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	Log.Infof("Starting server at %v ...", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		Log.Fatalf(err.Error())
	}
}
