package bunyan

import "os"
import "time"
import "bytes"
import "testing"
import a "github.com/stretchr/testify/assert"

func TestNewLogger(t *testing.T) {
	logger := NewLogger(NilSink())
	a.NotNil(t, logger, "cannot create new logger.")
}

func TestNewStdLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewStdLogger("log_test", NewJsonSink(buffer))
	a.NotNil(t, logger, "cannot create standard logger.")

	logger.Tracef("hello %s", "world")

	result, e := UnmarshalRecord(buffer.Bytes())
	a.NoError(t, e, "while unmarshaling result record.")

	hostname, e := os.Hostname()
	a.NoError(t, e, "while determining hostname.")

	a.Equal(t, result["name"], "log_test", "log name not written.")
	a.Equal(t, result["v"], 0, "output record does not include bunyan record version.")
	a.Equal(t, result["pid"], os.Getpid(), "PID not included.")
	a.Equal(t, result["hostname"], hostname, "Hostname not included.")
	a.IsType(t, "str", result["time"], "time value not included (or not a string.)")

	ref, e := time.Parse(time.RFC3339, result["time"].(string))
	a.NoError(t, e, "while parsing time field in output.")

	diff := time.Now().Sub(ref)
	a.True(t, diff.Seconds() < 1, "time recorded is not the correct time.")
}

