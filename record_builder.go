package bunyan

import "fmt"
// NOTE: Meant to be used as one-off record building

type RecordBuilder struct{
	sink Sink
	template Record
	record Record
}

func NewRecordBuilder(target Sink, template Record) *RecordBuilder {
	return &RecordBuilder{target, template, NewRecord()}
}

func (b *RecordBuilder) Write(record Record) error {
	return b.sink.Write(record)
}

func (b *RecordBuilder) Template() Record {
	return b.template
}

func (b *RecordBuilder) Record(key string, value interface{}) Log {
	b.record[key] = value
	return b
}

func (b *RecordBuilder) Recordf(key, value string, args...interface{}) Log {
	return b.Record(key, fmt.Sprintf(value, args...))
}

func (b *RecordBuilder) Child() Log {
	// TODO: Creates new *Logger with record as child template
	panic("not implemented.")
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

func (b *RecordBuilder) writef(level int, msg string, args...interface{}) {
	// TODO: Handle write errors.
	b.record.SetMessagef(level, msg, args...)
	b.sink.Write(b.record.TemplateMerge(b.template))
}
