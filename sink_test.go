package bunyan

import "io"
import "io/ioutil"
import "os"
import "fmt"
import "testing"
import a "github.com/stretchr/testify/assert"

func TestSinkFunc(t *testing.T) {
	var given Record
	var err error = fmt.Errorf("test error.")

	sink := SinkFunc(func(record Record) error {
		given = record
		return err
	})
	a.NotNil(t, sink, "cannot create a sink from a function.")

	record := NewSimpleRecord("hello", "world")
	e := sink.Write(record)
	a.Equal(t, err, e, "returns error from sink function.")
	a.Equal(t, record, given, "sink function were given wrong record.")
}

func ExampleStdoutSink() {
	sink := StdoutSink()
	sink.Write(NewSimpleRecord("the quick", "brown fox"))
	// Output:
	// {"the quick":"brown fox"}
}

func ExampleFileSink() {
	tmp, e := ioutil.TempFile("", "bunyan")
	if e != nil {
		panic(e)
	}

	path := tmp.Name()
	sink := FileSink(path)
	sink.Write(NewSimpleRecord("the brown", "quick fox"))
	tmp.Close()

	tmp, e = os.Open(path)
	defer tmp.Close()
	if e != nil {
		panic(e)
	}

  _, e = io.Copy(os.Stdout, tmp)
	if e != nil {
		panic(e)
	}

	// Output:
	// {"the brown":"quick fox"}
}

