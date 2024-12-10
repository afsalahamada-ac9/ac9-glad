//go:build testing

/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package config

const (
	// Database
	DB_USER     = "glad_user"
	DB_PASSWORD = "glad1234"
	DB_DATABASE = "glad"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = 5432 /* 3306 for MySQL */
	DB_SSLMODE  = "require"

	// Defaults
	DEFAULT_TENANT  = "5306526529902621696"
	DEFAULT_ACCOUNT = "100016472"

	// API port
	API_PORT = 8080

	// Metrics
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"

	// SFSYNCD specific consts
	COURSED_ADDR = "localhost:8080"

	// Firebase
	FCM_JSON     = ""
	PUSH_DRY_RUN = true
)
