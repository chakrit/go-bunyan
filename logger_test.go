package bunyan

import "os"
import "bytes"
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
	parent, _ := newTestLogger()
	parent.template = NewSimpleRecord("parent", "key")

	template := NewSimpleRecord("extra", "value")
	result := NewChildLogger(parent, template).(*Logger)

	a.NotNil(t, result, "cannot create child logger.")
	a.Equal(t, result.sinks[0], parent, "child should should use parent as sink.")
	a.Empty(t, result.infos, "child should start with empty infos.")
}

func TestLogger_Write(t *testing.T) {
	logger, buffer := newTestLogger()
	logger.template = NewSimpleRecord("test", "template")

	record := NewSimpleRecord("real", "record")
	record["pid"] = "overridden"
	e := logger.Write(record)
	a.NoError(t, e)

	result, e := UnmarshalRecord(buffer.Bytes())
	a.Equal(t, result["test"], "template", "template values not written to output.")
	a.Equal(t, result["real"], "record", "record data not written to output.")

	// test some standard infos overrides
	a.NotEqual(t, result["pid"], os.Getpid(), "info must be overridden by user keys.")
	a.IsType(t, "", result["time"], "registered time info not written to output.")

	// TODO: Test error propagation
}

func TestTemplate(t *testing.T) {
	logger, _ := newTestLogger()
	template := NewSimpleRecord("test", "template")
	logger.template = template

	result := logger.Template()
	a.Equal(t, result, template, "logger template not returned.")
}

func TestLogger_Record(t *testing.T) {
	logger, buffer := newTestLogger()
	builder := logger.Record("test", "value")

	a.NotNil(t, builder, "should returns a new record builder.")
	a.IsType(t, builder, &RecordBuilder{}, "should returns a new record builder.")

	builder.Record("more test", "value").Tracef("hi msg")

	result, e := UnmarshalRecord(buffer.Bytes())
	a.NoError(t, e)
	a.Equal(t, result["test"], "value", "output should contains recorded value.")
	a.Equal(t, result["more test"], "value", "output should contains recorded value.")
	a.Equal(t, result["msg"], "hi msg", "output have incorrect message.")
}

func newTestLogger() (*Logger, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
	return NewLogger("test-logger", NewJsonSink(buffer)).(*Logger), buffer
}

