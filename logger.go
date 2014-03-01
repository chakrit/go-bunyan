package bunyan

import "fmt"

// Logger is the default Log implementation provided by go-bunyan.
type Logger struct {
	sink   Sink
	record Record
}

// NewLogger() creates an empty Logger attached to a given sink.
func NewLogger(target Sink) *Logger {
	return &Logger{target, NewRecord()}
}

// Write() writes the given record to the attached sink. If the receiver is a child
// logger, then additional information from parent loggers will also be included.
func (l *Logger) Write(record Record) error {
	record.TemplateMerge(l.record)
	return l.sink.Write(record)
}

// Include() returns a new logger instance that automatically records information from the
// given Info instance to all its records.
func (l *Logger) Include(info Info) Log {
	return NewLogger(InfoSink(l, info))
}

// Record() adds the given key and value to the logger and returns itself.
// TODO: Wrong doc
func (l *Logger) Record(key string, value interface{}) Log {
	// TODO: Optimize. Don't New too damn much. Probably can get rid of TemplateMerge
	// altogether.
	builder := NewLogger(l)
	builder.record[key] = value
	return builder
}

// Recordf() provides formatting convenience that simply calls Record() with the formatted
// value.
func (l *Logger) Recordf(key, value string, args ...interface{}) Log {
	return l.Record(key, fmt.Sprintf(value, args...))
}

// Child() creates a child logger from the receiver logger. Any records written to the
// child ogger will also contains information recorded in the parent logger.
func (l *Logger) Child() Log {
	return NewLogger(l)
}

func (l *Logger) Tracef(msg string, args ...interface{}) { l.send(TRACE, msg, args...) }
func (l *Logger) Debugf(msg string, args ...interface{}) { l.send(DEBUG, msg, args...) }
func (l *Logger) Infof(msg string, args ...interface{})  { l.send(INFO, msg, args...) }
func (l *Logger) Warnf(msg string, args ...interface{})  { l.send(WARN, msg, args...) }
func (l *Logger) Errorf(msg string, args ...interface{}) { l.send(ERROR, msg, args...) }
func (l *Logger) Fatalf(msg string, args ...interface{}) { l.send(FATAL, msg, args...) }

func (l *Logger) send(level Level, msg string, args ...interface{}) {
	record := NewRecord()
	record.SetMessagef(level, msg, args...)
	e := l.Write(record)

	// TODO: Do not panic. Recover gracefully, maybe write something to stderr.
	if e != nil {
		panic(e)
	}
}
