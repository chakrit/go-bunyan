package bunyan

type FilteredSink struct{
	sink Sink
	filter Filter
}

func (sink *FilteredSink) Write(record Record) error {
	result := sink.filter.Apply(record)
	if !result {
		return nil
	}

	return sink.sink.Write(record)
}

func FilterSink(sink Sink, filters...Filter) Sink {
	return &FilteredSink{sink, MultiFilter(filters...)}
}
