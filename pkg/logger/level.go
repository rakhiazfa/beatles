package logger

import (
	"fmt"
	"strings"
)

type Level uint32

const (
	LevelPanic Level = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

func (level Level) Marshal() ([]byte, error) {
	switch level {
	case LevelDebug:
		return []byte("debug"), nil
	case LevelInfo:
		return []byte("info"), nil
	case LevelWarn:
		return []byte("warn"), nil
	case LevelError:
		return []byte("error"), nil
	case LevelFatal:
		return []byte("fatal"), nil
	case LevelPanic:
		return []byte("panic"), nil
	}

	return nil, fmt.Errorf("invalid log level: %s", level.String())
}

func (level *Level) Unmarshal(text []byte) (Level, error) {
	lvl := strings.ToLower(string(text))

	switch lvl {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "fatal":
		return LevelFatal, nil
	case "panic":
		return LevelPanic, nil
	}

	var l Level
	return l, fmt.Errorf("invalid log Level: %s", lvl)
}

func (level Level) String() string {
	if b, err := level.Marshal(); err == nil {
		return string(b)
	}

	return ""
}
