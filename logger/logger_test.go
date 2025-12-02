package logger

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	log := New()
	log.SetLevel(LevelDebug)

	fmt.Println()

	log.WithField("image", "ubuntu:latest").Debug("Example with field")
	log.WithField("image", "ubuntu:latest").Info("Example with field")
	log.WithField("image", "ubuntu:latest").Warn("Example with field")
	log.WithField("image", "ubuntu:latest").Error("Example with field")
	log.WithField("image", "ubuntu:latest").Fatal("Example with field")
	log.WithField("image", "ubuntu:latest").Panic("Example with field")

	log.WithField("image", "ubuntu:latest").Debugf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Infof("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Warnf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Errorf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Fatalf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Panicf("%s with field", "Example")

	log.WithFields(map[string]any{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}).Info("Example multiple fields")

	log.SetFormat(FormatJSON)

	log.WithField("image", "ubuntu:latest").Debug("Example with field")
	log.WithField("image", "ubuntu:latest").Info("Example with field")
	log.WithField("image", "ubuntu:latest").Warn("Example with field")
	log.WithField("image", "ubuntu:latest").Error("Example with field")
	log.WithField("image", "ubuntu:latest").Fatal("Example with field")
	log.WithField("image", "ubuntu:latest").Panic("Example with field")

	log.WithField("image", "ubuntu:latest").Debugf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Infof("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Warnf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Errorf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Fatalf("%s with field", "Example")
	log.WithField("image", "ubuntu:latest").Panicf("%s with field", "Example")

	log.WithFields(map[string]any{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}).Info("Example multiple fields")
}
