package bunyan

const InitialSinksCapacity = 4
const InitialFiltersCapacity = 8

type Logger struct{
	sinks []Sink
	filters []Filter
}

func NewEmptyLogger() *Logger {
	logger := &Logger{}
	logger.Clear()
	return logger
}

func NewLogger(name string) *Logger {
	logger := NewEmptyLogger()
	logger.AddFilter(StdFilter(name))
	logger.AddSink(EVERYTHING, StdoutSink())
	return logger
}

// TODO: NewLogger() with standard configuration

func (l *Logger) AddSink(level int, sink Sink) {
	switch {
	case level == EVERYTHING:
		l.sinks = append(l.sinks, sink)
	default:
		sink = FilterSink(sink, LevelFilter(level))
		l.sinks = append(l.sinks, sink)
	}
}

func (l *Logger) AddFilter(filter Filter) {
	l.filters = append(l.filters, filter)
}

func (l *Logger) Clear() {
	l.sinks = make([]Sink, 0, InitialSinksCapacity)
	l.filters = make([]Filter, 0, InitialFiltersCapacity)
}

