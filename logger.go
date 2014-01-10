package bunyan

type Logger struct {
	template Record
	sinks    []Sink
	infos    []Info
}

func NewLogger(sinks ...Sink) Log {
	return &Logger{nil, sinks[:], nil}
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
