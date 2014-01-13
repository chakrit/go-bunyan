package bunyan

import "fmt"

type RecordBuilder struct{
	sink Sink
	record Record
}

func NewRecordBuilder(target Sink) *RecordBuilder {
	return &RecordBuilder{target, NewRecord()}
}

func (b *RecordBuilder) Write(record Record) error {
	record.TemplateMerge(b.record)
	return b.sink.Write(record)
}

func (b *RecordBuilder) Include(info Info) Log {
	return NewRecordBuilder(InfoSink(b.sink, info))
}

func (b *RecordBuilder) Record(key string, value interface{}) Log {
	b.record[key] = value
	return b
}

func (b *RecordBuilder) Recordf(key, value string, args...interface{}) Log {
	return b.Record(key, fmt.Sprintf(value, args...))
}

func (b *RecordBuilder) Child() Log {
	return NewRecordBuilder(b)
}

func (b *RecordBuilder) Tracef(msg string, args...interface{}) {
	b.writef(TRACE, msg, args...)
}

func (b *RecordBuilder) Debugf(msg string, args...interface{}) {
	b.writef(DEBUG, msg, args...)
}

func (b *RecordBuilder) Infof(msg string, args...interface{}) {
	b.writef(INFO, msg, args...)
}

func (b *RecordBuilder) Warnf(msg string, args...interface{}) {
	b.writef(WARN, msg, args...)
}

func (b *RecordBuilder) Errorf(msg string, args...interface{}) {
	b.writef(ERROR, msg, args...)
}

func (b *RecordBuilder) Fatalf(msg string, args...interface{}) {
	b.writef(FATAL, msg, args...)
}

func (b *RecordBuilder) writef(level Level, msg string, args...interface{}) {
	record := NewRecord()
	record.SetMessagef(level, msg, args...)
	e := b.Write(record)

	// TODO: Do not panic. Recover gracefully, maybe write something to stderr.
	if e != nil {
		panic(e)
	}
}
