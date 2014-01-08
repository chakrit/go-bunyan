package bunyan

import "io"
import "encoding/json"

type JsonSink struct{
	*json.Encoder
}

func NewJsonSink(output io.Writer) *JsonSink {
	return &JsonSink{json.NewEncoder(output)}
}

func (sink *JsonSink) Write(record Record) error {
	return sink.Encoder.Encode(record)
}
