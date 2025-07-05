package helpers

import (
	"log"
	"log/slog"
	"os"
)

func InitializeStandardLogger(logLevel string) error {
	loggerConfig := &slog.HandlerOptions{
        Level: slog.LevelDebug,
	}

	switch logLevel {
	case "debug":
		loggerConfig.Level = slog.LevelDebug
	case "info":
		loggerConfig.Level = slog.LevelInfo
	case "warn":
		loggerConfig.Level = slog.LevelWarn
	case "error":
		loggerConfig.Level = slog.LevelError
	default:
		log.Println("Configuration log level is invalid, defaulting to debug")
		loggerConfig.Level = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stdout, loggerConfig)
	logger := slog.New(handler)
	
	slog.SetDefault(logger)
	log.SetOutput(slog.NewLogLogger(handler, slog.LevelInfo).Writer())
	
	return nil
}