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

// Logger is an interface that defines the basic logging methods for all log levels.
// And sync method to flush the logs in buffer before the main terminates.
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, err error, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
}

// ZapLoggerImpl is a concrete implementation of the Logger interface using zap.
type ZapLoggerImpl struct {
	logger *zap.Logger
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
