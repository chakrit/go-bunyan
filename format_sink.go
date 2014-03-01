package bunyan

import "io"
import "fmt"
import "time"
import "bytes"

type FormatSink struct {
	output io.Writer
}

func NewFormatSink(output io.Writer) *FormatSink {
	return &FormatSink{output}
}

// TODO: Colors
func (sink *FormatSink) Write(record Record) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}() // never panics

	line := &bytes.Buffer{}
	printed := map[string]bool{}

	write := func(key string, format string, filter func(interface{}) interface{}) {
		value, ok := record[key]
		if !ok {
			return
		}

		printed[key] = true
		if filter != nil {
			value = filter(value)
		}

		_, e = fmt.Fprintf(line, format, value)
		if e != nil {
			panic(e)
		}
	}

	// standard pre-defined keys
	write("time", "[%s]", func(v interface{}) interface{} {
		return v.(time.Time).Format(time.RFC3339)
	})

	write("level", "%6s: ", nil)
	write("name", "%s", nil)
	write("pid", "/%d", nil)
	write("hostname", " on %s", nil)
	write("msg", ": %s", nil)
	printed["v"] = true // do not print

	// extras:
	for key, value := range record {
		if printed[key] {
			continue
		}

		_, e := fmt.Fprintf(line, "\n  %s: %#v", key, value)
		if e != nil {
			panic(e)
		}
	}

	_, e = fmt.Fprintln(line)
	if e != nil {
		return e
	}

	_, e = sink.output.Write(line.Bytes())
	return
}
