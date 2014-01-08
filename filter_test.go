package bunyan

import "os"
import "time"
import "testing"
import a "github.com/stretchr/testify/assert"

func TestFilterFunc(t *testing.T) {
	called := false
	returnValue := false
	f := func(record Record) bool {
		record["called"] = true
		called = true
		return returnValue
	}

	record := NewRecord()

	filter := FilterFunc(f)
	a.NotNil(t, filter, "filter from function returns nil.")

	result := filter.Apply(record)
	a.True(t, called, "filter function not called on Apply.")
	a.Equal(t, result, returnValue, "return value not propagated.")
	a.True(t, record["called"].(bool), "record not modified.")

	returnValue, called = true, false
	result = filter.Apply(record)
	a.True(t, called, "filter function not called on Apply.")
	a.Equal(t, result, returnValue, "return value not propagated.")
}

func TestSimpleFilter(t *testing.T) {
	record, result := runFilter(SimpleFilter("key", "value"))

	a.True(t, result, "simple filter should always returns true.")
	checkRecord(t, record, "key", "value")
}

func TestPidFilter(t *testing.T) {
	pid := os.Getpid()
	record, result := runFilter(PidFilter())

	a.True(t, result, "pid filter should always returns true.")
	checkRecord(t, record, "pid", pid)
}

func TestHostnameFilter(t *testing.T) {
	hostname, e := os.Hostname()
	a.NoError(t, e)

	record, result := runFilter(HostnameFilter())

	a.True(t, result, "hostname filter should always returns true.")
	checkRecord(t, record, "hostname", hostname)
}

func TestTimeFilter(t *testing.T) {
	now := time.Now()

	record, result := runFilter(TimeFilter())

	a.True(t, result, "time filter should always returns true.")
	a.NotNil(t, record["time"], "time not recorded.")
	a.IsType(t, "", record["time"], "time key is of wrong type.")

	value, e := time.Parse(time.RFC3339, record["time"].(string))
	a.NoError(t, e)
	a.True(t, value.Sub(now) < time.Second, "wrong recorded time.")
}

func TestLevelFilter(t *testing.T) {
	record := NewRecord()
	record["level"] = 10

	filter := LevelFilter(30)
	result := filter.Apply(record)
	a.False(t, result, "level filter does not discard lower level logs.")

	record["level"] = 30
	result = filter.Apply(record)
	a.True(t, result, "level filter does not accept logs with same level.")

	record["level"] = 50
	result = filter.Apply(record)
	a.True(t, result, "level filter does not accept higher level logs.")

	delete(record, "level")
	result = filter.Apply(record)
	a.False(t, result, "level filter should not accept record without `level` key.")
}

func TestMultiFilter(t *testing.T) {
	one := SimpleFilter("one", "there can only be one.")
	two := SimpleFilter("two", "there maybe two.")
	three := FilterFunc(func(r Record) bool {
		r["three"] = "third time's a charm"
		return false
	})

	filter := MultiFilter(one, two)
	record, result := runFilter(filter)
	a.True(t, result, "multi filter result should be true if no filters return false.")
	checkRecord(t, record, "one", "there can only be one.")
	checkRecord(t, record, "two", "there maybe two.")

	filter = MultiFilter(one, two, three)
	record, result = runFilter(filter)
	a.False(t, result, "multi filter result should be false if one filter returns false.")
	checkRecord(t, record, "one", "there can only be one.")
	checkRecord(t, record, "two", "there maybe two.")
	checkRecord(t, record, "three", "third time's a charm")
}

func TestStdFilter(t *testing.T) {
	record, result := runFilter(StdFilter("test-name"))
	hostname, e := os.Hostname()
	a.NoError(t, e)

	a.True(t, result, "standard filter should always returns true.")
	checkRecord(t, record, "v", 0)
	checkRecord(t, record, "name", "test-name")
	checkRecord(t, record, "pid", os.Getpid())
	checkRecord(t, record, "hostname", hostname)
	a.IsType(t, "", record["time"], "time key is missing.")
}

func runFilter(filter Filter) (Record, bool) {
	record := NewRecord()
	result := filter.Apply(record)
	return record, result
}

func checkRecord(t *testing.T, r Record, key string, value interface{}) {
	a.NotNil(t, r, "record is nil.")
	a.NotNil(t, r[key], "record is missing expected key.")
	a.IsType(t, value, r[key], "value for key `%s` has wrong type.", key)
	a.Equal(t, value, r[key], "value for key `%s` is wrong.", key)
}
