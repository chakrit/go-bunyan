package bunyan

const InitialSinksCapacity = 4

type Logger struct{
	sinks []Sink
}

func NewEmptyLogger() *Logger {
	logger := &Logger{}
	logger.Clear()
	return logger
}

func NewLogger(name string) *Logger {
	logger := NewEmptyLogger()
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

func (l *Logger) Clear() {
	l.sinks = make([]Sink, 0, InitialSinksCapacity)
}

func (l *Logger) Write(record Record) error {
	// TODO: Handle Write() errors, never panic
	for _, sink := range l.sinks {
		e := sink.Write(record)
		if e != nil {
			panic(e)
		}
	}

	return nil
}

