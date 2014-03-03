package bunyan

import "io"
import "encoding/json"

type JsonPump struct {
	*json.Decoder
	Records chan Record
}

func NewJsonPump(input io.Reader) *JsonPump {
	pump := &JsonPump{
		Decoder: json.NewDecoder(input),
		Records: make(chan Record),
	}

	go pump.pump()
	return pump
}

func (pump *JsonPump) pump() {
	for {
		record := NewRecord()
		e := pump.Decoder.Decode(&record)

		// TODO: Gracefully stop synchronously (mutex?)
		if e == nil {
			pump.Records <- record

		} else if e == io.EOF {
			close(pump.Records)
			break

		} else {
			panic(e)
		}
	}
}
