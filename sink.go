package bunyan

import "os"

// Sink interface handles wiring the actual output of log records created from the
// loggers.
type Sink interface {
	Write(record Record) error
}

type funcSink struct {
	write func(record Record) error
}

func (sink *funcSink) Write(record Record) error {
	return sink.write(record)
}

// SinkFunc() creates a Sink with the Write() method implementation that simply calls the
// given function.
func SinkFunc(write func(record Record) error) Sink {
	return &funcSink{write}
}

// NilSink() creates a Sink that doesn't output anything. This can be used to temporary
// disables a logger or for testing in general.
func NilSink() Sink {
	return SinkFunc(func(record Record) error {
		return nil // no-op
	})
}

func InfoSink(target Sink, info Info) Sink {
	return SinkFunc(func(record Record) error {
		record.SetIfNot(info.Key(), info.Value())
		return target.Write(record)
	})
}

// StdoutSink() creats a Sink that writes records to the standard output.
func StdoutSink() Sink {
	return NewJsonSink(os.Stdout)
}

// FileSink() creates a Sink that writes records to a file.
func FileSink(path string) Sink {
	const flags = os.O_CREATE | os.O_APPEND | os.O_WRONLY
	file, e := os.OpenFile(path, flags, 0666)
	if e != nil {
		panic(e)
	}

	return NewJsonSink(file)
}

// TODO: RotatingFile.
