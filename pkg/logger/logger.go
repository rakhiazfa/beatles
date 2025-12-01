package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rakhiazfa/beatles/pkg/colorutils"
	"github.com/rakhiazfa/beatles/pkg/stringutils"
)

type logger struct {
	mu sync.Mutex

	Out    io.Writer
	Level  Level
	Format Format

	fields map[string]any
}

func New() Logger {
	return &logger{
		Out:    os.Stderr,
		Level:  LevelInfo,
		Format: FormatText,
		fields: make(map[string]any),
	}
}

func (l *logger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.Level
}

func (l *logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Level = level
}

func (l *logger) GetFormat() Format {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.Format
}

func (l *logger) SetFormat(format Format) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Format = format
}

func (l *logger) SetOutput(out io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Out = out
}

func (l *logger) WithField(key string, value any) Logger {
	return l.cloneWithFields(map[string]any{key: value})
}

func (l *logger) WithFields(fields map[string]any) Logger {
	return l.cloneWithFields(fields)
}

func (l *logger) Log(level Level, args ...any) {
	if level > l.GetLevel() {
		return
	}

	l.log(level, fmt.Sprint(args...))
}

func (l *logger) Logf(level Level, format string, args ...any) {
	if level > l.GetLevel() {
		return
	}

	l.log(level, fmt.Sprintf(format, args...))
}

func (l *logger) Debug(args ...any) {
	l.Log(LevelDebug, args...)
}

func (l *logger) Info(args ...any) {
	l.Log(LevelInfo, args...)
}

func (l *logger) Warn(args ...any) {
	l.Log(LevelWarn, args...)
}

func (l *logger) Error(args ...any) {
	l.Log(LevelError, args...)
}

func (l *logger) Fatal(args ...any) {
	l.Log(LevelFatal, args...)
}

func (l *logger) Panic(args ...any) {
	l.Log(LevelPanic, args...)
}

func (l *logger) Debugf(format string, args ...any) {
	l.Logf(LevelDebug, format, args...)
}

func (l *logger) Infof(format string, args ...any) {
	l.Logf(LevelInfo, format, args...)
}

func (l *logger) Warnf(format string, args ...any) {
	l.Logf(LevelWarn, format, args...)
}

func (l *logger) Errorf(format string, args ...any) {
	l.Logf(LevelError, format, args...)
}

func (l *logger) Fatalf(format string, args ...any) {
	l.Logf(LevelFatal, format, args...)
}

func (l *logger) Panicf(format string, args ...any) {
	l.Logf(LevelPanic, format, args...)
}

func (l *logger) cloneWithFields(fields map[string]any) Logger {
	newFields := make(map[string]any, len(l.fields)+len(fields))

	maps.Copy(newFields, l.fields)
	maps.Copy(newFields, fields)

	return &logger{
		Out:    l.Out,
		Level:  l.Level,
		Format: l.Format,
		fields: newFields,
	}
}

func (l *logger) getLogLevelColor(level Level) string {
	switch level {
	case LevelDebug:
		return colorutils.RGB(211, 212, 216)
	case LevelInfo:
		return colorutils.RGB(181, 195, 235)
	case LevelWarn:
		return colorutils.RGB(232, 220, 150)
	case LevelError:
		return colorutils.RGB(186, 152, 126)
	case LevelFatal:
		return colorutils.RGB(171, 130, 95)
	case LevelPanic:
		return colorutils.RGB(154, 112, 74)
	default:
		return colorutils.Reset
	}
}

func (l *logger) log(level Level, message string) {
	format := l.GetFormat()

	message = strings.TrimSuffix(message, "\n")

	switch format {
	case FormatJSON:
		l.logJSON(level, message)
	default:
		l.logText(level, message)
	}
}

func (l *logger) logText(level Level, message string) {
	fieldStr := ""
	for key, value := range l.fields {
		coloredKey := colorutils.Colorize(key, colorutils.RGB(117, 140, 221))
		fieldStr += fmt.Sprintf("%s=%v ", coloredKey, value)
	}

	coloredLevel := colorutils.Colorize(strings.ToUpper(level.String()), l.getLogLevelColor(level))
	formattedLevel := stringutils.PadRight(coloredLevel, 6)
	formattedMessage := stringutils.PadRight(message, 60)

	fmt.Fprintf(l.Out, "%s %s %s\n", formattedLevel, formattedMessage, fieldStr)
}

func (l *logger) logJSON(level Level, message string) {
	data := map[string]any{
		"level":   level.String(),
		"message": message,
		"time":    time.Now().Format(time.RFC3339),
	}

	maps.Copy(data, l.fields)

	b, _ := json.Marshal(data)

	l.Out.Write(b)
	l.Out.Write([]byte("\n"))
}
