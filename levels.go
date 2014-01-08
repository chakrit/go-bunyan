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
