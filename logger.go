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

// TODO: NewChildLogger
func NewChildLogger(parent Log, extraTemplate Record) Log {
	template := NewRecord()
	template.TemplateMerge(extraTemplate)
	template.TemplateMerge(parent.Template())

	return &Logger{template, []Sink{parent}, []Info{}}
}

func (logger *Logger) Write(record Record) error {
	return nil
}

func (logger *Logger) Template() Record {
	return nil
}

func (logger *Logger) Record(key string, value interface{}) Log {
	return nil
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
