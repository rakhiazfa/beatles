package logger

import (
	"io"
)

type Logger interface {
	GetLevel() Level
	SetLevel(level Level)

	GetFormat() Format
	SetFormat(format Format)

	SetOutput(out io.Writer)

	WithField(key string, value any) Logger
	WithFields(fields map[string]any) Logger

	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	Panic(args ...any)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Panicf(format string, args ...any)
}
