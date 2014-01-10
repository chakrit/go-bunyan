package bunyan

import "testing"
import a "github.com/stretchr/testify/assert"

func TestNewLogger(t *testing.T) {
	result := NewLogger(StdoutSink())
	a.NotNil(t, result, "cannot create new logger.")
}

