package bunyan

import "fmt"
import "bytes"
import "testing"
import "encoding/json"
import a "github.com/stretchr/testify/assert"

var _ Log = &RecordBuilder{}

func TestNewRecordBuilder(t *testing.T) {
	sink := StdoutSink()
	template := NewRecord()
	builder := NewRecordBuilder(sink, template)

	a.NotNil(t, builder, "cannot create new builder")
	a.NotNil(t, builder.record, "builder does not initialize new record.")
	a.Equal(t, builder.template, template, "builder does not saves given template.")
	a.Equal(t, builder.sink, sink, "builder does not saves given sink.")

	builder = NewRecordBuilder(sink, nil)
	a.Nil(t, builder.template, "builder does not initialize template unnecessarily.")
}

func TestWrite(t *testing.T) {
	template := NewSimpleRecord("templated", "value")
	builder, buffer := newTestBuilder(template)

	record := NewSimpleRecord("hello", "world")
	e := builder.Write(record)
	a.NoError(t, e)

	result := make(map[string]interface{})
	json.Unmarshal(buffer.Bytes(), &result)

	a.Equal(t, result["hello"], "world", "record not written to sink.")
	a.Equal(t, result["templated"], "value", "template not merged into written record.")
}

func TestRecord(t *testing.T) {
	builder, _ := newTestBuilder(nil)

	result := builder.Record("test", "value")
	a.Equal(t, builder.record["test"], "value", "value not saved.")
	a.Equal(t, result, builder, "does not return self.")
}

func TestRecordf(t *testing.T) {
	builder, _ := newTestBuilder(nil)

	result := builder.Recordf("test", "value %s", "result")
	a.NotNil(t, builder.record["test"], "value not saved.")
	a.Equal(t, builder.record["test"], "value result", "value not properly formatted.")
	a.Equal(t, result, builder, "does not return self.")
}

func TestChild(t *testing.T) {
	builder, _ := newTestBuilder(nil)
	child := builder.Record("childkey", "value").Child()
	a.NotNil(t, child, "cannot create child logger from record builder.")
	a.NotEqual(t, child, builder, "child must not be the builder.")

	template := child.Template()
	a.NotNil(t, template, "child should have non-nil template.")
	a.Equal(t, template["childkey"], "value", "child template has incorrect values.")
}

func TestLogMethods(t *testing.T) {
	template := NewRecord()
	template["one"] = "first template value."
	template["two"] = "second template value."

	builder, buffer := newTestBuilder(template)

	type logFunc func(string, ...interface{})
	test := func(f logFunc, level int, msg string, args...interface{}) {
		buffer.Reset()
		f(msg, args...)

		result := make(map[string]interface{})
		e := json.Unmarshal(buffer.Bytes(), &result)
		a.NoError(t, e)

		a.Equal(t, result["level"], level, "result has wrong level value.")
		a.Equal(t, result["msg"], fmt.Sprintf(msg, args...), "result has wrong message.")

		for k, v := range template {
			a.Equal(t, result[k], v, "template key `%s` not merged into result.", k)
		}
	}

	// TODO: infer via reflection?
	mappings := map[int]logFunc{
		TRACE: builder.Tracef,
		DEBUG: builder.Debugf,
		INFO: builder.Infof,
		WARN: builder.Warnf,
		ERROR: builder.Errorf,
		FATAL: builder.Fatalf,
	}

	for lvl, f := range mappings {
		test(f, lvl, "hello %d %s", lvl, "world")
	}
}

func newTestBuilder(template Record) (*RecordBuilder, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
	builder := NewRecordBuilder(NewJsonSink(buffer), template)
	return builder, buffer
}
