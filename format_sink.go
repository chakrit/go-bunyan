package bunyan

import "io"
import "fmt"
import "time"
import "bytes"
import "strconv"

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

	noError := func(e error) {
		if e != nil {
			panic(e)
		}
	}

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

	// converters
	// TODO: Better handle these things
	levelify := func(v interface{}) interface{} {
		switch v.(type) {
		case string:
			return v.(string)
		case []byte:
			return string(v.([]byte))
		case Level:
			return v.(Level)
		case int:
			return Level(v.(int))
		case float64: // json decode result
			return Level(int(v.(float64)))
		default:
			return v
		}
	}

	intify := func(v interface{}) interface{} {
		switch v.(type) {
		case string:
			result, e := strconv.Atoi(v.(string))
			noError(e)
			return result
		case []byte:
			result, e := strconv.Atoi(string(v.([]byte)))
			noError(e)
			return result
		case Level:
			return v.(Level)
		case float64:
			return int(v.(float64))
		default:
			return v
		}
	}

	timify := func(v interface{}) interface{} {
		switch v.(type) {
		case string:
			return v.(string)
		case []byte:
			return string(v.([]byte))
		case time.Time:
			return v.(time.Time).Format(time.RFC3339)
		default:
			return v
		}
	}

	// standard pre-defined keys
	write("time", "[%s]", timify)
	write("level", "%6s: ", levelify)
	write("name", "%s", nil)
	write("pid", "/%d", intify)
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
