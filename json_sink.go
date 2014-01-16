package bunyan

import "io"
import "encoding/json"

// JsonSink is the core Sink implementation that marshals records into its JSON format and
// writes it to the given io.Writer.
type JsonSink struct {
	*json.Encoder
}

// NewJsonSink() creates a JsonSink implementation attached to the given io.Writer.
func NewJsonSink(output io.Writer) *JsonSink {
	return &JsonSink{json.NewEncoder(output)}
}

func (sink *JsonSink) Write(record Record) error {
	return sink.Encoder.Encode(record)
}
