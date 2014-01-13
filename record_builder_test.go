package bunyan

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

var recordBuilder = LogFactory(func(buffer *bytes.Buffer) Log {
	return NewRecordBuilder(NewJsonSink(buffer))
})

func TestRecordBuilder_Write(t *testing.T) {
	Log_Write(t, recordBuilder)
}

func TestRecordBuilder_Record(t *testing.T) {
	Log_Record(t, recordBuilder)
}

func TestRecordBuilder_Recordf(t *testing.T) {
	Log_Recordf(t, recordBuilder)
}

func TestRecordBuilder_Child(t *testing.T) {
	// Log_Child(t, recordBuilder)
}

func TestLogMethods(t *testing.T) {
	Log_Methods(t, recordBuilder)
}
