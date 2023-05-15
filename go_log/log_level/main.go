package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logDebug(logger)
	logInfo(logger)
	logWarn(logger)
	logError(logger)
	logFatal(logger)
}

func logDebug(logger *zap.Logger) {
	logger.Debug("Debug level log message")
}

func logInfo(logger *zap.Logger) {
	logger.Info("Info level log message")
}

func logWarn(logger *zap.Logger) {
	logger.Warn("Warn level log message")
}

func logError(logger *zap.Logger) {
	logger.Error("Error level log message")
}

func logFatal(logger *zap.Logger) {
	logger.Fatal("Fatal level log message")
}
