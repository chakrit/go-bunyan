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

func TestFormatSink_Write_Time(t *testing.T) {
	input := "2014-03-02T00:33:59+07:00"
	time_, e := time.Parse(time.RFC3339, input)
	a.NoError(t, e)

	expected := "[2014-03-02T00:33:59+07:00]\n"

	checkFormat(t, NewSimpleRecord("time", input), expected)
	checkFormat(t, NewSimpleRecord("time", time_), expected)
}

func TestFormatSink_Write_Level(t *testing.T) {
	expected := "  INFO: \n"

	checkFormat(t, NewSimpleRecord("level", float64(30)), expected)
	checkFormat(t, NewSimpleRecord("level", int(30)), expected)
	checkFormat(t, NewSimpleRecord("level", INFO), expected)
}

func TestFormatSink_Write_Pid(t *testing.T) {
	expected := "/35330\n"

	checkFormat(t, NewSimpleRecord("pid", float64(35330)), expected)
	checkFormat(t, NewSimpleRecord("pid", int(35330)), expected)
	checkFormat(t, NewSimpleRecord("pid", "35330"), expected)
	checkFormat(t, NewSimpleRecord("pid", 35330), expected)
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
}

func checkFormat(t *testing.T, record Record, expected string) {
	buffer := &bytes.Buffer{}
	sink := NewFormatSink(buffer)

	e := sink.Write(record)
	a.NoError(t, e)
	a.Equal(t, expected, string(buffer.Bytes()), "format sink output is wrong.")
}
