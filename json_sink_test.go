package bunyan_test

import "os"
import "bytes"
import "testing"
import . "github.com/chakrit/go-bunyan"
import a "github.com/stretchr/testify/assert"

var _ Sink = NewJsonSink(&bytes.Buffer{})

func TestNewJsonSink(t *testing.T) {
	buffer := &bytes.Buffer{}
	sink := NewJsonSink(buffer)

	a.NotNil(t, sink, "cannot create json sink.")
	a.NotNil(t, sink.Encoder, "json sink must initialize encoder.")
}

func ExampleJsonSink() {
	sink := NewJsonSink(os.Stdout)
	sink.Write(NewSimpleRecord("hello", "world"))
	// Output:
	// {"hello":"world"}
}
