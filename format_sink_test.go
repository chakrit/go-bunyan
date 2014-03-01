package bunyan_test

import "os"
import "bytes"
import "time"
import "testing"
import . "github.com/chakrit/go-bunyan"
import a "github.com/stretchr/testify/assert"

var _ Sink = NewFormatSink(&bytes.Buffer{})

func TestNewFormatSink(t *testing.T) {
	buffer := &bytes.Buffer{}
	sink := NewFormatSink(buffer)
	a.NotNil(t, sink, "format sink ctor returns null.")
}

func ExampleFormatSink() {
	t, e := time.Parse(time.RFC3339, "2014-03-02T00:33:59+07:00")
	if e != nil {
		panic(e)
	}

	record := Record(map[string]interface{}{
		"v":        0,
		"pid":      6503,
		"hostname": "go-bunyan.local",
		"time":     t,
		"address":  ":8080",
		"user":     "chakrit",
		"level":    INFO,
		"msg":      "starting up...",
		"name":     "bunyan",
	})

	e = NewFormatSink(os.Stdout).Write(record)
	if e != nil {
		panic(e)
	}

	// Output:
	// [2014-03-02T00:33:59+07:00]  INFO: bunyan/6503 on go-bunyan.local: starting up...
	//   address: ":8080"
	//   user: "chakrit"
}
