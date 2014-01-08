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

func TestMergeRecord(t *testing.T) {
	one := NewSimpleRecord("hello", "world")
	one["argus"] = "was not here"
	two := NewSimpleRecord("argus", "magnimus") // should override

	result := MergeRecord(one, two)
	a.NotNil(t, result, "merged record should not be nil.")
	a.NotNil(t, result["hello"], "keys from first record not merged.")
	a.Equal(t, result["hello"], "world", "keys from first record has wrong value.")
	a.NotNil(t, result["argus"], "keys from second record not merged.")
	a.Equal(t, result["argus"], "magnimus", "keys from second record has wrong value.")
}
