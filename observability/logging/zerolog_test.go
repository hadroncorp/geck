package logging_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hadroncorp/geck/application"
	"github.com/hadroncorp/geck/observability/logging"
	"github.com/hadroncorp/geck/versioning"
)

func TestZerologLogger_Level(t *testing.T) {
	cfg := application.Config{
		ApplicationName: "foo-app",
		Environment:     "dev",
		Version:         "v0.0.1-alpha",
		Semver:          versioning.SemanticVersion{},
	}

	buf := bytes.NewBuffer(nil)
	logger := logging.NewZerologLoggerAdapter(logging.NewApplicationZerologLogger(cfg, buf))
	logger.Trace().
		WithField("str", "string").
		WithField("int", 100).
		WithField("float64", float32(99.99)).
		Write("some message")
	out := buf.String()
	assert.Equal(t, "{\"level\":\"trace\",\"application_name\":\"foo-app\",\"application_environment\":\"dev\",\"application_version\":\"v0.0.1-alpha\",\"str\":\"string\",\"int\":100,\"float64\":99.99,\"message\":\"some message\"}\n", out)

	buf.Reset()
	logger.Module("app.context").Warn().WithField("field", true).Write("some message")
	out = buf.String()
	assert.Equal(t, "{\"level\":\"warn\",\"application_name\":\"foo-app\",\"application_environment\":\"dev\",\"application_version\":\"v0.0.1-alpha\",\"module\":\"app.context\",\"field\":true,\"message\":\"some message\"}\n", out)

	buf.Reset()
	logger.Error().WithField("field", true).Write("some message")
	out = buf.String()
	assert.Equal(t, "{\"level\":\"error\",\"application_name\":\"foo-app\",\"application_environment\":\"dev\",\"application_version\":\"v0.0.1-alpha\",\"field\":true,\"message\":\"some message\"}\n", out)

	buf.Reset()
	err := errors.New("some error")
	logger.WithError(err).WithField("field", true).Write("some message")
	out = buf.String()
	assert.Equal(t, "{\"level\":\"error\",\"application_name\":\"foo-app\",\"application_environment\":\"dev\",\"application_version\":\"v0.0.1-alpha\",\"error\":\"some error\",\"field\":true,\"message\":\"some message\"}\n", out)
}
