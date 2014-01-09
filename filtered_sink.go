package bunyan

type filteredSink struct{
	level int
	sink Sink
}

func (sink *filteredSink) Write(record Record) error {
	lvl, ok := record["level"].(int)
	if !ok || lvl < sink.level {
		return nil // filtered
	}

	return sink.sink.Write(record)
}

func FilterSink(level int, sink Sink) Sink {
	return &filteredSink{level, sink}
}
