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
	"ac9/glad/pkg/logger"
	"ac9/glad/services/gcd/handler"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"ac9/glad/config"
	"ac9/glad/pkg/metric"
	"ac9/glad/pkg/middleware"
	"ac9/glad/pkg/util"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urfave/negroni"
)

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
	n.Use(&middleware.APILogging{Log: Log})

	// log handler
	logger.MakeLogHandlers(r, *n, "gcd", Log)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// log handler
	logger.MakeLogHandlers(r, *n, "gcd", Log)

	// test
	handler.MakeTestHandlers(r, *n)

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
