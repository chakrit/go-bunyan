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
		sink = FilterSink(level, sink)
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

func (l *Logger) Write(record Record) error {
	// TODO: Handle Write() errors, never panic

	valid := true
	for _, filter := range l.filters {
		valid_ := filter.Apply(record)
		valid = valid && valid_
	}
	if !valid {
		return nil
	}

	for _, sink := range l.sinks {
		e := sink.Write(record)
		if e != nil {
			panic(e)
		}
	}

	return nil
}

