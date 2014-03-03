package bunyan_test

import "bytes"
import "testing"
import . "github.com/chakrit/go-bunyan"
import a "github.com/stretchr/testify/assert"

func TestNewJsonPump(t *testing.T) {
	input := []byte(`{"hello":"world"}`)
	buffer := bytes.NewBuffer(input)
	pump := NewJsonPump(buffer)

	a.NotNil(t, pump, "cannot create json pump.")
	a.NotNil(t, pump.Decoder, "pump does not initialize decoder.")
	a.NotNil(t, pump.Records, "pump does not initialize output channel.")
}

func TestJsonPump_Records(t *testing.T) {
	input := `
	{"hello":"world"}
	{"second":"item"}
	{"third":"item"}
	`

	buffer := bytes.NewBuffer([]byte(input))
	pump := NewJsonPump(buffer)

	first, second, third := <-pump.Records, <-pump.Records, <-pump.Records
	a.NotNil(t, first, "first record read is not nil.")
	a.NotNil(t, second, "second record read is not nil.")
	a.NotNil(t, third, "third record read is not nil.")

	a.Equal(t, first["hello"], "world", "first record has incorrect values.")
	a.Equal(t, second["second"], "item", "second record has incorrect values.")
	a.Equal(t, third["third"], "item", "third record has incorrect values.")
}
