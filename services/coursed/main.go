/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"ac9/glad/repository"

	// Uber zap logging
	"ac9/glad/pkg/logger"

	"ac9/glad/usecase/account"
	"ac9/glad/usecase/center"
	"ac9/glad/usecase/course"
	"ac9/glad/usecase/product"
	"ac9/glad/usecase/tenant"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"ac9/glad/config"
	"ac9/glad/pkg/metric"
	"ac9/glad/pkg/middleware"
	"ac9/glad/pkg/util"
	"ac9/glad/services/coursed/handler"

	_ "github.com/go-sql-driver/mysql"
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

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		util.GetStrEnvOrConfig("DB_USER", config.DB_USER),
		util.GetStrEnvOrConfig("DB_PASSWORD", config.DB_PASSWORD),
		util.GetStrEnvOrConfig("DB_HOST", config.DB_HOST),
		util.GetStrEnvOrConfig("DB_DATABASE", config.DB_DATABASE),
		util.GetStrEnvOrConfig("DB_SSLMODE", config.DB_SSLMODE))
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		Log.Fatalf("Unable to initialize database: %v", err.Error())
	}
	defer db.Close()

	accountRepo := repository.NewAccountPGSQL(db)
	accountService := account.NewService(accountRepo)

	centerRepo := repository.NewCenterPGSQL(db)
	centerService := center.NewService(centerRepo)

	courseRepo := repository.NewCoursePGSQL(db)
	courseTimingRepo := repository.NewCourseTimingPGSQL(db)
	courseService := course.NewService(courseRepo, courseTimingRepo)

	productRepo := repository.NewProductPGSQL(db)
	productService := product.NewService(productRepo)

	tenantRepo := repository.NewTenantPGSQL(db)
	tenantService := tenant.NewService(tenantRepo)

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
		negroni.HandlerFunc(middleware.AddDefaultAccount),
		negroni.NewLogger(),
	)
	n.Use(&middleware.APILogging{Log: Log})

	// log handler
	logger.MakeLogHandlers(r, *n, "coursed", Log)

	// account
	handler.MakeAccountHandlers(r, *n, accountService)

	// center
	handler.MakeCenterHandlers(r, *n, centerService)

	// course
	handler.MakeCourseHandlers(r, *n, courseService, accountService)

	// participant
	handler.MakeParticipantHandlers(r, *n, accountService)

	// product
	handler.MakeProductHandlers(r, *n, productService)

	// tenant
	handler.MakeTenantHandlers(r, *n, tenantService)

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
