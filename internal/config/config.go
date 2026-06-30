package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServiceName     string
	HTTPAddr        string
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	LogLevel        slog.Level
}

func Load() Config {
	return Config{
		ServiceName:     getenv("SERVICE_NAME", "go-service-starter"),
		HTTPAddr:        getenv("HTTP_ADDR", ":8080"),
		ShutdownTimeout: getDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
		ReadTimeout:     getDuration("HTTP_READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getDuration("HTTP_WRITE_TIMEOUT", 10*time.Second),
		LogLevel:        getLogLevel(getenv("LOG_LEVEL", "info")),
	}
}

func getenv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return fallback
	}
	return d
}

func getLogLevel(v string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "info", "":
		return slog.LevelInfo
	default:
		if n, err := strconv.Atoi(v); err == nil {
			return slog.Level(n)
		}
		return slog.LevelInfo
	}
}
