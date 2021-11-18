package logger

import (
	"github.com/rs/zerolog"
	"io"
	"strings"
)

type Options struct {
	Level      string
	TimeFormat string
	Writer     io.Writer
}

func Create(options Options) zerolog.Logger {
	var level zerolog.Level
	switch strings.ToUpper(options.Level) {
	case "ERROR":
		level = zerolog.ErrorLevel
	case "WARN":
		level = zerolog.WarnLevel
	case "INFO":
		level = zerolog.InfoLevel
	case "DEBUG", "TRACE":
		level = zerolog.DebugLevel
	default:
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	format := options.TimeFormat
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	zerolog.TimeFieldFormat = format

	ctx := zerolog.
		New(options.Writer).
		With().
		Timestamp().
		Stack()
	if level == zerolog.DebugLevel {
		ctx = ctx.Caller()
	}

	return ctx.Logger()
}
