package bunyan

import "os"

type Sink interface{
	Write(record Record) error
}

type funcSink struct{
	write func(record Record) error
}

func (sink *funcSink) Write(record Record) error {
	return sink.write(record)
}

func SinkFunc(write func(record Record) error) Sink {
	return &funcSink{write}
}

func InfoSink(target Sink, info Info) Sink {
	return SinkFunc(func(record Record) error {
		record.SetIfNot(info.Key(), info.Value())
		return target.Write(record)
	})
}

func StdoutSink() Sink {
	return NewJsonSink(os.Stdout)
}

func FileSink(path string) Sink {
	const flags = os.O_CREATE|os.O_APPEND|os.O_WRONLY
	file, e := os.OpenFile(path, flags, 0666)
	if e != nil {
		panic(e)
	}

	return NewJsonSink(file)
}
