package bunyan

import "fmt"

// RecordBuilder is the default Log implementation provided by go-bunyan.
type RecordBuilder struct {
	sink   Sink
	record Record
}

// NewRecordBuilder() creates an empty RecordBuilder attached to a given sink.
func NewRecordBuilder(target Sink) *RecordBuilder {
	return &RecordBuilder{target, NewRecord()}
}

// Write() writes the given record to the attached sink. If the receiver is a child
// logger, then additional information from parent loggers will also be included.
func (b *RecordBuilder) Write(record Record) error {
	record.TemplateMerge(b.record)
	return b.sink.Write(record)
}

// Include() returns a new logger instance that automatically records information from the
// given Info instance to all its records.
func (b *RecordBuilder) Include(info Info) Log {
	return NewRecordBuilder(InfoSink(b, info))
}

// Record() adds the given key and value to the logger and returns itself.
func (b *RecordBuilder) Record(key string, value interface{}) Log {
	// TODO: Optimize. Don't New too damn much. Probably can get rid of TemplateMerge
	// altogether.
	builder := NewRecordBuilder(b)
	builder.record[key] = value
	return builder
}

// Recordf() provides formatting convenience that simply calls Record() with the formatted
// value.
func (b *RecordBuilder) Recordf(key, value string, args ...interface{}) Log {
	// TODO: Needs Child() calls
	return b.Record(key, fmt.Sprintf(value, args...))
}

// Child() creates a child logger from the receiver logger. Any records written to the
// child ogger will also contains information recorded in the parent logger.
func (b *RecordBuilder) Child() Log {
	return NewRecordBuilder(b)
}

func (b *RecordBuilder) Tracef(msg string, args ...interface{}) {
	b.writef(TRACE, msg, args...)
}

func (b *RecordBuilder) Debugf(msg string, args ...interface{}) {
	b.writef(DEBUG, msg, args...)
}

func (b *RecordBuilder) Infof(msg string, args ...interface{}) {
	b.writef(INFO, msg, args...)
}

func (b *RecordBuilder) Warnf(msg string, args ...interface{}) {
	b.writef(WARN, msg, args...)
}

func (b *RecordBuilder) Errorf(msg string, args ...interface{}) {
	b.writef(ERROR, msg, args...)
}

func (b *RecordBuilder) Fatalf(msg string, args ...interface{}) {
	b.writef(FATAL, msg, args...)
}

func (b *RecordBuilder) writef(level Level, msg string, args ...interface{}) {
	record := NewRecord()
	record.SetMessagef(level, msg, args...)
	e := b.Write(record)

	// TODO: Do not panic. Recover gracefully, maybe write something to stderr.
	if e != nil {
		panic(e)
	}
}
