package logging_test

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hadroncorp/geck/observability/logging"
)

func TestStdLogger_Level(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	logger := logging.NewStdLoggerAdapter(log.New(buf, "", 0))
	logger.Debug().
		WithField("str", "string").
		WithField("int", 100).
		WithField("float64", float32(99.99)).
		Write("some message")
	out := buf.String()
	assert.NotEmpty(t, out)
	assert.True(t, strings.HasPrefix(out, "DEBUG"))
	assert.True(t, strings.HasSuffix(out, "\n"))
	assert.Equal(t, len("DEBUG str:\"string\" int:100 float64:99.99 message:\"some message\"\n"), len(out))

	buf.Reset()
	logger.Module("app.context").Info().WithField("field", true).Write("some message")
	out = buf.String()
	assert.True(t, strings.HasPrefix(out, "INFO"))
	assert.True(t, strings.HasSuffix(out, "\n"))
	assert.Equal(t, len("INFO module:\"app.context\" field:true message:\"some message\"\n"), len(out))

	buf.Reset()
	logger.Error().WithField("field", true).Write("some message")
	out = buf.String()
	assert.True(t, strings.HasPrefix(out, "ERROR"))
	assert.True(t, strings.HasSuffix(out, "\n"))
	assert.Equal(t, len("ERROR field:true message:\"some message\"\n"), len(out))

	buf.Reset()
	err := errors.New("some error")
	logger.WithError(err).WithField("field", true).Write("some message")
	out = buf.String()
	assert.Equal(t, len("ERROR error:\"some error\" field:true message:\"some message\"\n"), len(out))
}
