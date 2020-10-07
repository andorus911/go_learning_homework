package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func NewLogger(logFile, logLevel string) *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	var zlLevel zapcore.Level
	switch logLevel {
	case "error":
		zlLevel = zap.ErrorLevel
	case "warn":
		zlLevel = zap.WarnLevel
	case "info":
		zlLevel = zap.InfoLevel
	case "debug":
		zlLevel = zap.DebugLevel
	}

	cfg.Level.SetLevel(zlLevel)
	cfg.OutputPaths = append(cfg.OutputPaths, logFile)

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Unable to create logger, %v", err)
	}
	defer logger.Sync()

	logger.Info("log is created")

	return logger
}
