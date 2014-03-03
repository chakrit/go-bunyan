package main

import "os"
import . "github.com/chakrit/go-bunyan"

func main() {
	input := NewJsonPump(os.Stdin)
	output := NewFormatSink(os.Stdout)

	for record := range input.Records {
		e := output.Write(record)
		if e != nil {
			panic(e)
		}
	}
}
