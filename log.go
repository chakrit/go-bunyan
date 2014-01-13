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
