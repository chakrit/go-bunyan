package bunyan

type Log interface{
	Sink

	Include(info Info) Log
	Record(key string, value interface{}) Log
	Recordf(key, value string, args...interface{}) Log
	Child() Log // create child logger with what's recorded so far as template

	Tracef(msg string, args...interface{}) // logs with what's recorded so far included
	Debugf(msg string, args...interface{})
	Infof(msg string, args...interface{})
	Warnf(msg string, args...interface{})
	Errorf(msg string, args...interface{})
	Fatalf(msg string, args...interface{})
}

func NewLogger(output Sink) Log {
	return NewRecordBuilder(output)
}

func NewStdLogger(name string) Log {
	return newStdLogger(name, StdoutSink())
}

func newStdLogger(name string, sink Sink) Log {
	// allow changing sink in tests
	log := NewRecordBuilder(sink).
		Record("name", name).
		Include(PidInfo()).
		Include(HostnameInfo()).
		Include(TimeInfo())

	return log
}
