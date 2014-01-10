package bunyan

import "testing"
import a "github.com/stretchr/testify/assert"

var _ Log = &Logger{}

func TestNewEmptyLogger(t *testing.T) {
	result := NewEmptyLogger().(*Logger)
	a.NotNil(t, result, "cannot create empty logger.")
	a.Empty(t, result.sinks, "empty logger should have no sinks.")
	a.Empty(t, result.infos, "empty logger should have no infos.")
}

func TestNewLogger(t *testing.T) {
	sink := StdoutSink()
	result := NewLogger("test-logger", sink).(*Logger)

	a.NotNil(t, result, "cannot create new logger.")
	a.Equal(t, result.sinks[0], sink, "logger does not saves given sink.")
	a.NotEmpty(t, result.infos, "logger should starts with some basic info.")
}

func TestChildLogger(t *testing.T) {
	sink := StdoutSink()
	parent := NewLogger("test-logger", sink).(*Logger)
	parent.template = NewSimpleRecord("parent", "key")

	template := NewSimpleRecord("extra", "value")
	result := NewChildLogger(parent, template).(*Logger)

	a.NotNil(t, result, "cannot create child logger.")
	a.Equal(t, result.sinks[0], parent, "child should should use parent as sink.")
	a.Empty(t, result.infos, "child should start with empty infos.")
}
