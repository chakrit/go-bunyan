package bunyan

import "io"
import "io/ioutil"
import "os"

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

