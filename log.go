package bunyan

// REF: https://github.com/trentm/node-bunyan#levels
const(
	EVERYTHING = 10 * iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type Log interface{
	Sink

	Template() Record // gets the record template that'll be merged on Write()

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
