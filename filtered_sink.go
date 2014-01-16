package bunyan

type filteredSink struct {
	level Level
	sink  Sink
}

func (sink *filteredSink) Write(record Record) error {
	lvl, ok := record["level"].(Level)
	if !ok || lvl < sink.level {
		return nil // filtered
	}

	return sink.sink.Write(record)
}

// FilterSink() is a meta-sink that filter incoming records by the given level. Any
// records that has level below the set level will be discarded.
//
// Useful for reducing log volumes to the standard output or capturing only errors to
// a pre-configured target.
func FilterSink(level Level, sink Sink) Sink {
	return &filteredSink{level, sink}
}
