package bunyan_test

import . "github.com/chakrit/go-bunyan"

func ExampleFilterSink() {
	sink := FilterSink(INFO, StdoutSink())

	records := map[string]Level{
		"first message":  TRACE,
		"second message": DEBUG,
		"third message":  INFO,
		"fourth message": WARN,
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
