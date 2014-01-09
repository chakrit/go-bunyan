package bunyan

import "bytes"
import "testing"
import "encoding/json"
import a "github.com/stretchr/testify/assert"

var _ Sink = &Logger{}

type DummyFilter struct{}

func (d *DummyFilter) Apply(record Record) bool {
	return true
}

func TestNewEmptyLogger(t *testing.T) {
	logger := NewEmptyLogger()

	a.NotNil(t, logger, "cannot create empty logger.")
	a.NotNil(t, logger.sinks, "initialize sinks array.")
	a.NotNil(t, logger.filters, "initialize filters array.")
	a.Equal(t, cap(logger.sinks), InitialSinksCapacity, "sinks array capacity wrong.")
	a.Equal(t, cap(logger.filters), InitialFiltersCapacity, "filters array capacity wrong.")
	a.Empty(t, logger.sinks, "sinks array should be initially empty.")
	a.Empty(t, logger.filters, "filters array should be initially empty.")
}

func TestAddSink(t *testing.T) {
	logger, sink := NewEmptyLogger(), StdoutSink()
	logger.AddSink(INFO, sink)

	a.IsType(t, &filteredSink{}, logger.sinks[0], "sink added without filter.")
	sink_ := logger.sinks[0].(*filteredSink)
	a.Equal(t, sink_.sink, sink, "filter sink outlet is not the given sink.")

	logger, sink = NewEmptyLogger(), StdoutSink()
	logger.AddSink(EVERYTHING, sink)
	a.IsType(t, &JsonSink{}, logger.sinks[0], "`everything` sink should not have a filter.")
	a.Equal(t, logger.sinks[0], sink, "wrong sink is added.")
}

func TestAddFilter(t *testing.T) {
	logger, filter := NewEmptyLogger(), &DummyFilter{}
	logger.AddFilter(filter)
	a.Equal(t, logger.filters[0], filter, "filter not added.")
}

func TestClearFilters(t *testing.T) {
	logger, sink, filter := NewEmptyLogger(), StdoutSink(), &DummyFilter{}
	logger.AddSink(EVERYTHING, sink)
	logger.AddFilter(filter)
	logger.Clear()
	a.Empty(t, logger.filters, "filters array not cleared.")
	a.Empty(t, logger.sinks, "sinks array not cleared.")
}

func TestWrite(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := NewEmptyLogger()
	logger.AddSink(INFO, NewJsonSink(buffer))

	record := NewRecord()
	record["level"] = INFO
	record["msg"] = "test message."
	e := logger.Write(record)
	a.NoError(t, e)

	result := make(map[string]interface{})
	e = json.Unmarshal(buffer.Bytes(), &result)
	a.NoError(t, e)

	a.True(t, len(result) > 0, "output json reading failed!")
	a.Equal(t, result["level"], INFO, "wrong level written.")
	a.Equal(t, result["msg"], "test message.", "wrong message written.")

	buffer.Reset()
	record["level"] = TRACE
	record["msg"] = "trace message hidden."

	e = logger.Write(record)
	a.NoError(t, e)
	a.Equal(t, buffer.Len(), 0, "trace message should not be written.")
}
