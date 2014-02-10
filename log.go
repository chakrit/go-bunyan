package bunyan

// Log interface is the core interface you will interacting with while using go-bunyan.
// This interface contains methods for writing standard log level messages to configured
// sinks as well as methods for recording extra information and for creating child
// loggers.
type Log interface {
	Sink

	Include(info Info) Log
	Record(key string, value interface{}) Log
	Recordf(key, value string, args ...interface{}) Log
	Child() Log // create child logger with what's recorded so far as template

	Tracef(msg string, args ...interface{}) // logs with what's recorded so far included
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
}

// NewEmptyLogger() creates a new, empty logger attached to the given Sink. This is not
// recommended unless you have a very specific use case or have fine-grained customization
// needs. NewStdLogger() is recommended over this method.
func NewEmptyLogger(output Sink) Log {
	return NewLogger(output)
}

// NewStdLogger() creates the standard bunyan logger with the given name and attached to
// the given Sink. Standard information includes the version key, PID, machine's hostname
// and time.
//
// Using the standard logger ensures that records will be compatible with node-bunyan CLI
// tool.
func NewStdLogger(name string, output Sink) Log {
	// allow changing sink in tests
	log := NewLogger(output).
		Record("name", name).
		Include(LogVersionInfo(0)). // should follows node-bunyan
		Include(PidInfo()).
		Include(HostnameInfo()).
		Include(TimeInfo())

	return log
}
