package bunyan

import "testing"
import a "github.com/stretchr/testify/assert"

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

	a.IsType(t, &FilteredSink{}, logger.sinks[0], "sink added without filter.")
	sink_ := logger.sinks[0].(*FilteredSink)
	a.Equal(t, sink_.sink, sink, "inner sink is not the given sink.")

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
