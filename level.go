package bunyan

import "strings"

type Level byte

// See node-bunyan description of each log level for more information:
// https://github.com/trentm/node-bunyan#levels
const (
	EVERYTHING Level = 10 * iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	UNKNOWN = 255
)

// AllLevels() returns all valid logging levels excluding the special EVERYTHING and UNKNOWN level.
func AllLevels() []Level {
	return []Level{
		TRACE,
		DEBUG,
		INFO,
		WARN,
		ERROR,
		FATAL,
	}
}

// String() returns the string representation of the level.
func (l Level) String() string {
	switch l {
	case EVERYTHING:
		return "EVERYTHING"
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// ParseLevel() attempts to parse the given string into a proper Level value.
func ParseLevel(str string) Level {
	str = strings.ToUpper(str)
	switch str {
	case "EVERYTHING":
		return EVERYTHING
	case "TRACE":
		return TRACE
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return UNKNOWN
	}
}
