package bunyan

// import "fmt"
// import "bytes"
// import "testing"
// import a "github.com/stretchr/testify/assert"

// type LogFactory func(*bytes.Buffer) Log

// func Log_Write(t *testing.T, factory LogFactory) {
// 	logger, buffer := newTestLogger(factory)

// 	record := NewSimpleRecord("hello", "world")
// 	e := logger.Write(record)
// 	a.NoError(t, e)

// 	result, e := UnmarshalRecord(buffer.Bytes())
// 	a.NoError(t, e)
// 	checkRecordEquals(t, record, result)
// }

// func Log_Record(t *testing.T, factory LogFactory) {
// 	logger, buffer := newTestLogger(factory)
// 	logger.Record("first", "value").
// 		Record("second", "value").
// 		Infof("write!")

// 	json := "{\"first\":\"value\",\"second\":\"value\",\"level\":%d,\"msg\":\"write!\"}"
// 	expected, e := UnmarshalRecord([]byte(fmt.Sprintf(json, INFO)))
// 	a.NoError(t, e)

// 	result, e := UnmarshalRecord(buffer.Bytes())
// 	a.NoError(t, e)
// 	checkRecordEquals(t, expected, result)
// }

// func Log_Recordf(t *testing.T, factory LogFactory) {
// 	logger, buffer := newTestLogger(factory)
// 	logger.Recordf("formatted", "values: %s %s", "hello", "world").Infof("write!")

// 	json := "{\"formatted\":\"values: hello world\",\"level\":%d,\"msg\":\"write!\"}"
// 	expected, e := UnmarshalRecord([]byte(fmt.Sprintf(json, INFO)))
// 	a.NoError(t, e)

// 	result, e := UnmarshalRecord(buffer.Bytes())
// 	a.NoError(t, e)
// 	checkRecordEquals(t, expected, result)
// }

// func Log_Child(t *testing.T, factory LogFactory) {
// 	logger, buffer := newTestLogger(factory)
// 	child := logger.Record("parent", "value").Child()
// 	a.NotNil(t, child, "failed to create child logger.")

// 	child.Infof("child", "value")

// 	json := "{\"parent\":\"value\",\"child\":\"value\",\"level\":%d,\"msg\":\"write!\"}"
// 	expected, e := UnmarshalRecord([]byte(fmt.Sprintf(json, INFO)))
// 	a.NoError(t, e)

// 	result, e := UnmarshalRecord(buffer.Bytes())
// 	a.NoError(t, e)
// 	checkRecordEquals(t, expected, result)
// }

// func Log_Methods(t *testing.T, factory LogFactory) {
// 	logger, buffer := newTestLogger(factory)

// 	type logFunc func(string, ...interface{})
// 	test := func(f logFunc, level Level, msg string, args...interface{}) {
// 		buffer.Reset()
// 		f(msg, args...)

// 		result, e := UnmarshalRecord(buffer.Bytes())
// 		a.NoError(t, e)
// 		a.Equal(t, result["level"], int(level), "result has wrong level value.")
// 		a.Equal(t, result["msg"], fmt.Sprintf(msg, args...), "result has wrong message.")
// 	}

// 	mappings := map[Level]logFunc{
// 		TRACE: logger.Tracef,
// 		DEBUG: logger.Debugf,
// 		INFO: logger.Infof,
// 		WARN: logger.Warnf,
// 		ERROR: logger.Errorf,
// 		FATAL: logger.Fatalf,
// 	}

// 	for lvl, f := range mappings {
// 		test(f, lvl, "hello %d %s", lvl, "world")
// 	}
// }

// func newTestLogger(factory LogFactory) (Log, *bytes.Buffer) {
// 	buffer := &bytes.Buffer{}
// 	logger := factory(buffer)
// 	return logger, buffer
// }

// func checkRecordEquals(t *testing.T, expected, result Record) {
// 	a.Equal(t, len(expected), len(result), "result has wrong number of keys.")

// 	for k, v := range expected {
// 		result, ok := result[k]
// 		a.True(t, ok, "key `%s` missing in result", k)
// 		a.Equal(t, v, result, "key `%s` has wrong value in result", k)
// 	}
// }
