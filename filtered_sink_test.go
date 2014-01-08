package bunyan

var _ Sink = &FilteredSink{}

func ExampleFilteredSink() {
	sink := FilterSink(StdoutSink(), LevelFilter(30))

	records := map[string]int{
		"first message": 10,
		"second message": 20,
		"third message": 30,
		"fourth message": 40,
	}

	for msg, lvl := range records {
		r := NewRecord()
		r["msg"], r["level"] = msg, lvl
		sink.Write(r)
	}

	// Output:
	// {"level":30,"msg":"third message"}
	// {"level":40,"msg":"fourth message"}
}
