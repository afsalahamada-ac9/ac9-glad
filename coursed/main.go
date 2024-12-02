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

	"sudhagar/glad/repository"

	// Uber zap logging
	"sudhagar/glad/pkg/logger"

	"sudhagar/glad/usecase/account"
	"sudhagar/glad/usecase/center"
	"sudhagar/glad/usecase/course"
	"sudhagar/glad/usecase/product"
	"sudhagar/glad/usecase/tenant"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"sudhagar/glad/config"
	"sudhagar/glad/coursed/handler"
	"sudhagar/glad/pkg/metric"
	"sudhagar/glad/pkg/middleware"
	"sudhagar/glad/pkg/util"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
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

	for _, env := range os.Environ() {
		Log.Infof(env)
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		util.GetStrEnvOrConfig("DB_USER", config.DB_USER),
		util.GetStrEnvOrConfig("DB_PASSWORD", config.DB_PASSWORD),
		util.GetStrEnvOrConfig("DB_HOST", config.DB_HOST),
		// util.GetIntEnvOrConfig("DB_PORT", config.DB_PORT),
		util.GetStrEnvOrConfig("DB_DATABASE", config.DB_DATABASE),
		util.GetStrEnvOrConfig("DB_SSLMODE", config.DB_SSLMODE))
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		Log.Fatalf("Unable to initialize database: %v", err.Error())
	}
	defer db.Close()

	productRepo := repository.NewProductPGSQL(db)
	productService := product.NewService(productRepo)

	centerRepo := repository.NewCenterPGSQL(db)
	centerService := center.NewService(centerRepo)

	tenantRepo := repository.NewTenantPGSQL(db)
	tenantService := tenant.NewService(tenantRepo)

	accountRepo := repository.NewAccountPGSQL(db)
	accountService := account.NewService(accountRepo)

	courseRepo := repository.NewCoursePGSQL(db)
	courseTimingRepo := repository.NewCourseTimingPGSQL(db)
	courseService := course.NewService(courseRepo, courseTimingRepo)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := mux.NewRouter()
	// handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.AddDefaultTenant),
		negroni.NewLogger(),
	)
	// center
	handler.MakeCenterHandlers(r, *n, centerService)

	// tenant
	handler.MakeTenantHandlers(r, *n, tenantService)

	// account
	handler.MakeAccountHandlers(r, *n, accountService)

	// course
	handler.MakeCourseHandlers(r, *n, courseService)

	// product
	handler.MakeProductHandlers(r, *n, productService)

	// log handler
	logger.MakeLogHandlers(r, *n, "coursed", Log)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		Log.Debugf("Health check called")
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
