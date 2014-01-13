package bunyan

import "fmt"
import "bytes"
import "testing"
import a "github.com/stretchr/testify/assert"

var _ Log = &RecordBuilder{}

func TestNewRecordBuilder(t *testing.T) {
	sink := StdoutSink()
	builder := NewRecordBuilder(sink)
	a.NotNil(t, builder, "cannot create new builder")
	a.NotNil(t, builder.record, "builder does not initialize new record.")
	a.Equal(t, builder.sink, sink, "builder does not saves given sink.")
}

func TestRecordBuilder_Write(t *testing.T) {
	json := "{\"hello\":\"world\"}"
	excercise(t, json, func(log Log) error {
		return log.Write(helloWorld())
	})
}

func TestRecordBuilder_Include(t *testing.T) {
	json := "{\"info\":\"included\",\"hello\":\"world\"}"
	excercise(t, json, func(log Log) error {
		return log.Include(SimpleInfo("info", "included")).Write(helloWorld())
	})
}

func TestRecordBuilder_Record(t *testing.T) {
	json := "{\"extra\":\"info\",\"hello\":\"world\"}"
	excercise(t, json, func(log Log) error {
		return log.Record("extra", "info").Write(helloWorld())
	})
}

func TestRecordBuilder_Recordf(t *testing.T) {
	json := "{\"extra\":\"info formatted\",\"hello\":\"world\"}"
	excercise(t, json, func(log Log) error {
		return log.Recordf("extra", "info %s", "formatted").Write(helloWorld())
	})
}

func TestRecordBuilder_Child(t *testing.T) {
	json := "{\"child info\":\"via parent\",\"hello\":\"world\"}"
	excercise(t, json, func(log Log) error {
		child := log.Record("child info", "via parent").Child()
		return child.Write(helloWorld())
	})
}

func TestLogMethods(t *testing.T) {
	type logFunc func(msg string, args...interface{})
	funcFor := func(lvl Level, log Log) logFunc {
		switch lvl {
		case TRACE:
			return log.Tracef
		case DEBUG:
			return log.Debugf
		case INFO:
			return log.Infof
		case WARN:
			return log.Warnf
		case ERROR:
			return log.Errorf
		case FATAL:
			return log.Fatalf
		default:
			panic(fmt.Errorf("no log method for level: %s", lvl))
		}
	}

	for _, lvl := range AllLevels() {
		json := fmt.Sprintf("{\"level\":%d,\"msg\":\"hello loggables\"}", lvl)
		excercise(t, json, func(log Log) error {
			logf := funcFor(lvl, log)
			logf("hello %s", "loggables")
			return nil
		})
	}
}

func helloWorld() Record {
	return NewSimpleRecord("hello", "world")
}

func excercise(t *testing.T, expectedJson string, ex func(log Log) error) {
	buffer := &bytes.Buffer{}
	log := NewRecordBuilder(NewJsonSink(buffer))

	e := ex(log)
	a.NoError(t, e)

	expected, e := UnmarshalRecord([]byte(expectedJson))
	a.NoError(t, e)

	result, e := UnmarshalRecord(buffer.Bytes())
	a.NoError(t, e)
	a.Equal(t, len(expected), len(result), "result has wrongg number of keys.")
	if len(result) != len(expected) {
		str, _ := result.Marshal()
		panic(string(str))
	}

	for k, v := range expected {
		result, ok := result[k]
		a.True(t, ok, "expected key `%s` missing in result.", k)
		a.Equal(t, v, result, "key `%s` has wrong value in result.", k)
	}
}
