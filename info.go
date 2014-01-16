package bunyan

import "os"
import "time"

// Info interface is for automatically including dynamically-generated information to log
// records.
type Info interface {
	Key() string
	Value() interface{}
}

type funcInfo struct {
	key       string
	valueFunc func() interface{}
}

type simpleInfo struct {
	key   string
	value interface{}
}

func (i *funcInfo) Key() string {
	return i.key
}

func (i *funcInfo) Value() interface{} {
	return i.valueFunc()
}

func (i *simpleInfo) Key() string {
	return i.key
}

func (i *simpleInfo) Value() interface{} {
	return i.value
}

// InfoFunc() returns an Info implementation that calls the given function to obtain
// the value to add to the record.
func InfoFunc(key string, valueFunc func() interface{}) Info {
	return &funcInfo{key, valueFunc}
}

// SimpleInfo() returns an Info implementation that attaches the given value to all
// records.
func SimpleInfo(key string, value interface{}) Info {
	return &simpleInfo{key, value}
}

// LogVersionInfo() returns an Info implementation that adds a `v` key with the given
// version value. Useful for compatibility with the node-bunyan tool which expects a `v`
// key with compatible value.
func LogVersionInfo(v int) Info {
	return SimpleInfo("v", v)
}

// PidInfo() returns an Info implementation that records current process ID number.
func PidInfo() Info {
	return SimpleInfo("pid", os.Getpid())
}

// HostnameInfo() returns an Info implementation that records the current machine
// hostname.
func HostnameInfo() Info {
	hostname, e := os.Hostname()
	if e != nil {
		panic(e)
	}

	return SimpleInfo("hostname", hostname)
}

// TimeInfo() returns an Info implementation that records the current time.
func TimeInfo() Info {
	return InfoFunc("time", func() interface{} {
		return time.Now().Format(time.RFC3339)
	})
}
