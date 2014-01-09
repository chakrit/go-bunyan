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

	Record(key string, value interface{}) Log // builds new record
	Child() Log // child logger with what's recorded so far

	Tracef(msg string, args...interface{})
	Debugf(msg string, args...interface{})
	Infof(msg string, args...interface{})
	Warnf(msg string, args...interface{})
	Errorf(msg string, args...interface{})
	Fatal(msg string, args...interface{})
}
