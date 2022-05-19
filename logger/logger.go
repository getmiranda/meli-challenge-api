package logger

import (
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const (
	envLogLevel        = "LOG_LEVEL"
	envLogEncodeOutput = "LOG_ENCODE_OUTPUT"
)

var (
	logger zerolog.Logger
)

func init() {
	zerolog.SetGlobalLevel(getLevel())
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	logger = zerolog.New(getEncode()).With().Timestamp().Logger().With().Caller().Logger()
	zerolog.DefaultContextLogger = &logger

	logger.Info().Msg("Logger initialized")
}

func GetLogger() zerolog.Logger {
	return logger
}

func getLevel() zerolog.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}

func getEncode() io.Writer {
	output := strings.TrimSpace(os.Getenv(envLogEncodeOutput))
	switch output {
	case "json":
		return os.Stdout
	case "console":
		return zerolog.ConsoleWriter{Out: os.Stderr}
	default:
		return zerolog.ConsoleWriter{Out: os.Stderr}
	}
}
