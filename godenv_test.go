package godenv_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/youla-dev/godenv"
)

func TestParse(t *testing.T) {
	raw := `
HTTP_ADDR=:80
JAEGER_AGENT_ENDPOINT="jaeger-otlp-agent:6831"
LOG_FORMAT='json'
LOG_LEVEL=debug
WAIT_BEFORE_EXIT=0s`

	expected := map[string]string{
		"HTTP_ADDR":             ":80",
		"JAEGER_AGENT_ENDPOINT": "jaeger-otlp-agent:6831",
		"LOG_FORMAT":            "json",
		"LOG_LEVEL":             "debug",
		"WAIT_BEFORE_EXIT":      "0s",
	}

	input := bytes.NewBufferString(raw)
	values, err := godenv.Parse(input)
	require.NoError(t, err)
	assert.Equal(t, expected, values)
}
