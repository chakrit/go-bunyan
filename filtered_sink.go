package bunyan

type filteredSink struct{
	level Level
	sink Sink
}

func (sink *filteredSink) Write(record Record) error {
	lvl, ok := record["level"].(Level)
	if !ok || lvl < sink.level {
		return nil // filtered
	}

	return sink.sink.Write(record)
}

func FilterSink(level Level, sink Sink) Sink {
	return &filteredSink{level, sink}
}
