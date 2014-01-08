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
