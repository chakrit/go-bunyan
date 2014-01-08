package bunyan

import "os"

type Sink interface{
	Write(record Record) error
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
