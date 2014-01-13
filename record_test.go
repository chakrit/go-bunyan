package bunyan

import "testing"
import a "github.com/stretchr/testify/assert"

func TestNewRecord(t *testing.T) {
	result := NewRecord()
	a.NotNil(t, result, "cannot create new record.")
}

func TestNewSimpleRecord(t *testing.T) {
	const key = "the quick brown fox"
	const value = "jumps over the lazy dog."

	result := NewSimpleRecord(key, value)
	a.NotNil(t, result, "cannot create simple record.")
	a.Equal(t, len(result), 1, "result have incorrect length.")
}

func TestUnmarshalRecord(t *testing.T) {
	str := "{\"hello\":\"world\",\"witch\":\"craft\"}"

	result, e := UnmarshalRecord([]byte(str))
	a.NoError(t, e)
	a.Equal(t, result["hello"], "world", "unmarshalled record wrong.")
	a.Equal(t, result["witch"], "craft", "unmarshalled record wrong.")
}

func TestMarshal(t *testing.T) {
	record := NewSimpleRecord("hello", "world")
	record["witch"] = "craft"

	result, e := record.Marshal()
	expected := "{\"hello\":\"world\",\"witch\":\"craft\"}"

	a.NoError(t, e)
	a.Equal(t, string(result), expected, "marshalled result wrong.")
}

func TestSetIfNot(t *testing.T) {
	record := NewSimpleRecord("hello", "world")

	result := record.SetIfNot("hello", "not world")
	a.Equal(t, result, record, "should returns self.")
	a.Equal(t, result["hello"], "world", "should not overrides value already set.")

	result = result.SetIfNot("second", "key")
	a.Equal(t, result["second"], "key", "should save values not already set.")
}

func TestTemplateMerge(t *testing.T) {
	one := NewSimpleRecord("hello", "world")
	one["argus"] = "was not here"
	two := NewSimpleRecord("argus", "magnimus") // should override

	result := two.TemplateMerge(one)
	a.NotNil(t, result, "result should not be nil.")
	a.Equal(t, result, two, "should returns itself, merged.")
	a.Equal(t, len(result), 2, "merged record has incorrect length.")
	a.NotNil(t, result["hello"], "keys from template not merged.")
	a.Equal(t, result["hello"], "world", "keys from template has wrong value.")
	a.NotNil(t, result["argus"], "existing keys in record missing.")
	a.Equal(t, result["argus"], "magnimus", "keys from template should not overrides.")

	// should supports nil template automatically
	one = NewSimpleRecord("argus", "maximus")
	result = one.TemplateMerge(nil)
	a.Equal(t, result, one, "should returns itself, merged.")
	a.Equal(t, len(result), 1, "merged record has incorrect length.")
	a.Equal(t, result["argus"], "maximus", "existing keys in recrod missing.")
}

func TestSetMessagef(t *testing.T) {
	r := NewSimpleRecord("hello", "world")
	r.SetMessagef(INFO, "test %s message", "info")

	a.Equal(t, r["hello"], "world", "setting message alter existing keys.")
	a.Equal(t, r["level"], INFO, "message level not saved.")
	a.Equal(t, r["msg"], "test info message", "message content not saved.")
}
