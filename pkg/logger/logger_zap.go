/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLoggerImpl is a concrete implementation of the Logger interface using zap.
type ZapLoggerImpl struct {
	cfg    zap.Config
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func NewLoggerZap() *ZapLoggerImpl {

	// Read log level from environment variable
	// e.g., "debug", "info", "warn", "error"
	logLevelEnv := os.Getenv("LOG_LEVEL")
	if logLevelEnv == "" {
		// Default log level
		logLevelEnv = "INFO"
	}

	// Convert environment variable to zapcore.Level
	zapLevel, err := zapcore.ParseLevel(logLevelEnv)
	if err != nil {
		log.Printf("Invalid log level %s. err=%v", logLevelEnv, err)
		return nil
	}

	// Create a new production zap config
	cfg := zap.NewProductionConfig()
	// Override the default config of zap
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.000")
	cfg.Level.SetLevel(zapLevel)

	// Create the logger instance with the custom config
	zapLogger, err := cfg.Build()
	if err != nil {
		log.Printf("Unable to create zap logger: %v", err)
		return nil
	}

	// Set the global Log to a ZapLoggerImpl instance
	logger := &ZapLoggerImpl{
		cfg:    cfg,
		logger: zapLogger,
		sugar:  zapLogger.Sugar(),
	}

	return logger
}

func (z *ZapLoggerImpl) SetLevel(level string) error {
	if zapLevel, err := zapcore.ParseLevel(level); err != nil {
		return err
	} else {
		z.cfg.Level.SetLevel(zapLevel)
	}

	return nil
}

func (z *ZapLoggerImpl) GetLevel() string {
	return z.cfg.Level.String()
}

func (z *ZapLoggerImpl) Debugf(template string, args ...interface{}) {
	z.sugar.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
}

func (z *ZapLoggerImpl) Infof(template string, args ...interface{}) {
	z.sugar.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
}

func (z *ZapLoggerImpl) Warnf(template string, args ...interface{}) {
	z.sugar.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
}

func (z *ZapLoggerImpl) Errorf(template string, args ...interface{}) {
	z.sugar.WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
}

func (z *ZapLoggerImpl) Panicf(template string, args ...interface{}) {
	z.sugar.WithOptions(zap.AddCallerSkip(1)).DPanicf(template, args...)
}

func (z *ZapLoggerImpl) Fatalf(template string, args ...interface{}) {
	z.sugar.WithOptions(zap.AddCallerSkip(1)).Fatalf(template, args...)
}

func (z *ZapLoggerImpl) Debug(msg string, fields ...zap.Field) {
	z.logger.Debug(msg, fields...)
}

func (z *ZapLoggerImpl) Info(msg string, fields ...zap.Field) {
	z.logger.Info(msg, fields...)
}

func (z *ZapLoggerImpl) Warn(msg string, fields ...zap.Field) {
	z.logger.Warn(msg, fields...)
}

func (z *ZapLoggerImpl) Error(msg string, err error, fields ...zap.Field) {

	if err != nil {
		fields = append(fields, zap.Error(err)) // Attach the error as a field
	}
	z.logger.Error(msg, fields...)
}

func (z *ZapLoggerImpl) Panic(msg string, fields ...zap.Field) {
	z.logger.Panic(msg, fields...)
}

func (z *ZapLoggerImpl) Fatal(msg string, fields ...zap.Field) {
	z.logger.Fatal(msg, fields...)
}

// Implement the Sync method
func (z *ZapLoggerImpl) Sync() error {
	// Call zap's Sync() to flush any buffered logs
	return z.logger.Sync()
}
