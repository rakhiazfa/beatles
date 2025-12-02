package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestLogger() (Logger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}

	log := New()
	log.SetLevel(LevelDebug)
	log.SetOutput(buffer)

	return log, buffer
}

func TestLogger_Panic(t *testing.T) {
	log, buffer := newTestLogger()

	assert.Panics(t, func() {
		log.Panic("Invalid parameter")
	})
	assert.Contains(t, buffer.String(), "Invalid parameter")
}

func TestLogger_Panicf(t *testing.T) {
	log, buffer := newTestLogger()

	assert.Panics(t, func() {
		log.Panicf("Invalid parameter %s", "email")
	})
	assert.Contains(t, buffer.String(), "Invalid parameter email")
}

func TestLogger_Error(t *testing.T) {
	log, buffer := newTestLogger()

	log.Error("Invalid parameter")

	assert.Contains(t, buffer.String(), "Invalid parameter")
}

func TestLogger_Errorf(t *testing.T) {
	log, buffer := newTestLogger()

	log.Errorf("Invalid parameter %s", "email")

	assert.Contains(t, buffer.String(), "Invalid parameter email")
}

func TestLogger_Warn(t *testing.T) {
	log, buffer := newTestLogger()

	log.Warn("Skipping user email john.doe@example.com")

	assert.Contains(t, buffer.String(), "Skipping user email john.doe@example.com")
}

func TestLogger_Warnf(t *testing.T) {
	log, buffer := newTestLogger()

	log.Warnf("Skipping user email %s", "john.doe@example.com")

	assert.Contains(t, buffer.String(), "Skipping user email john.doe@example.com")
}

func TestLogger_Info(t *testing.T) {
	log, buffer := newTestLogger()

	log.Info("Hello, World")

	assert.Contains(t, buffer.String(), "Hello, World")
}

func TestLogger_Infof(t *testing.T) {
	log, buffer := newTestLogger()

	log.Infof("Hello, %s", "John")

	assert.Contains(t, buffer.String(), "Hello, John")
}

func TestLogger_Debug(t *testing.T) {
	log, buffer := newTestLogger()

	log.Debug("Hello, World")

	assert.Contains(t, buffer.String(), "Hello, World")
}

func TestLogger_Debugf(t *testing.T) {
	log, buffer := newTestLogger()

	log.Debugf("Hello, %s", "John")

	assert.Contains(t, buffer.String(), "Hello, John")
}
