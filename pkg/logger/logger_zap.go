/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package logger

import (
	"go.uber.org/zap"
)

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
