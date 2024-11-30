/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package logger

import (
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global variable of type Logger (interface)
var Log Logger

func InitLogger() error {

	// Read log level from environment variable
	logLevelEnv := os.Getenv("LOG_LEVEL") // e.g., "debug", "info", "warn", "error"
	if logLevelEnv == "" {
		logLevelEnv = "INFO" // Default log level
	}

	// Convert environment variable to zapcore.Level
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(strings.ToLower(logLevelEnv))); err != nil {
		log.Printf("Invalid log level in environment variable, valid are [DEBUG, INFO, WARN, ERROR]: " + logLevelEnv)
		return err
	}

	// Create a new production zap config
	cfg := zap.NewProductionConfig()
	// Override the default config of zap
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.000")
	cfg.Level.SetLevel(logLevel)

	// Create the logger instance with the custom config
	zapLogger, err := cfg.Build()
	if err != nil {
		log.Printf("unable to create zap logger: %v", err)
		return err
	}

	// Set the global Log to a ZapLoggerImpl instance
	Log = &ZapLoggerImpl{
		logger: zapLogger,
	}

	return nil
}
