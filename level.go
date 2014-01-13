package bunyan

import "strings"

type Level byte

// REF: https://github.com/trentm/node-bunyan#levels
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

func ParseLevel(str string) Level {
	str = strings.ToUpper(str)
	switch str {
	case "EVEYRTHING":
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
