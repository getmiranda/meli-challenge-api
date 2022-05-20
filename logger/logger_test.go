package logger

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGetLevel(t *testing.T) {
	t.Setenv("LOG_LEVEL", "debug")
	level := getLevel()
	assert.Equal(t, zerolog.DebugLevel, level)

	t.Setenv("LOG_LEVEL", "info")
	level = getLevel()
	assert.Equal(t, zerolog.InfoLevel, level)

	t.Setenv("LOG_LEVEL", "error")
	level = getLevel()
	assert.Equal(t, zerolog.ErrorLevel, level)
}

func TestGetEncode(t *testing.T) {
	t.Setenv("LOG_ENCODE_OUTPUT", "json")
	output := getEncode()
	assert.Equal(t, os.Stdout, output)

	t.Setenv("LOG_ENCODE_OUTPUT", "console")
	output = getEncode()
	assert.Equal(t, zerolog.ConsoleWriter{Out: os.Stderr}, output)
}

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	assert.NotNil(t, logger)
}
