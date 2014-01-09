package bunyan

import "os"
import "time"
import "testing"
import a "github.com/stretchr/testify/assert"

func TestInfoFunc(t *testing.T) {
	info := InfoFunc("hello", func() interface{} {
		return "return value"
	})

	checkInfo(t, info, "hello", "return value")
}

func TestSimpleInfo(t *testing.T) {
	checkInfo(t, SimpleInfo("key", "value"), "key", "value")
}

func TestPidInfo(t *testing.T) {
	checkInfo(t, PidInfo(), "pid", os.Getpid())
}

func TestHostnameInfo(t *testing.T) {
	hostname, e := os.Hostname()
	a.NoError(t, e)
	checkInfo(t, HostnameInfo(), "hostname", hostname)
}

func TestTimeInfo(t *testing.T) {
	now := time.Now()
	info := TimeInfo()

	a.NotNil(t, info, "failed to create time info.")
	a.Equal(t, info.Key(), "time", "time info has wrong key.")

	raw := info.Value()
	a.IsType(t, "str", raw, "time info should returns string.")

	value, e := time.Parse(time.RFC3339, raw.(string))
	a.NoError(t, e)
	a.True(t, value.Sub(now) < time.Second, "wrong time returned from time info.")
}

func checkInfo(t *testing.T, info Info, expectedKey string, expectedValue interface{}) {
	a.NotNil(t, info, "cannot create info.")
	a.Equal(t, expectedKey, info.Key(), "wrong key returned.")
	a.Equal(t, expectedValue, info.Value(), "wrong value returned.")
}
