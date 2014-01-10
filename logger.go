package bunyan

type Logger struct {
	template Record
	sinks    []Sink
	infos    []Info
}

func NewEmptyLogger() Log {
	return &Logger{nil, []Sink{}, []Info{}}
}

func NewLogger(name string, sinks ...Sink) Log {
	infos := []Info{
		SimpleInfo("name", name),
		PidInfo(),
		HostnameInfo(),
		TimeInfo(),
	}

	return &Logger{nil, sinks[:], infos}
}

func NewChildLogger(parent Sink, template Record) Log {
	return &Logger{template, []Sink{parent}, []Info{}}
}

func (logger *Logger) Write(record Record) error {
	record.TemplateMerge(logger.template)
	for _, info := range logger.infos {
		record.SetIfNot(info.Key(), info.Value())
	}

	for _, sink := range logger.sinks {
		sink.Write(record)
	}
	return nil
}

func (logger *Logger) Template() Record {
	return logger.template
}

func (logger *Logger) Record(key string, value interface{}) Log {
	return NewRecordBuilder(logger, nil).Record(key, value)
}

func (logger *Logger) Recordf(key, value string, args ...interface{}) Log {
	return nil
}

func (logger *Logger) Child() Log {
	return nil
}

func (logger *Logger) Tracef(msg string, args ...interface{}) {
}

func (logger *Logger) Debugf(msg string, args ...interface{}) {
}

func (logger *Logger) Infof(msg string, args ...interface{}) {
}

func (logger *Logger) Warnf(msg string, args ...interface{}) {
}

func (logger *Logger) Errorf(msg string, args ...interface{}) {
}

func (logger *Logger) Fatalf(msg string, args ...interface{}) {
}
