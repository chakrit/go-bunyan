package bunyan

import "os"
import "time"

type Info interface{
	Key() string
	Value() interface{}
}

type funcInfo struct{
	key string
	valueFunc func() interface{}
}

type simpleInfo struct{
	key string
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

func InfoFunc(key string, valueFunc func() interface{}) Info {
	return &funcInfo{key, valueFunc}
}

func SimpleInfo(key string, value interface{}) Info {
	return &simpleInfo{key, value}
}

func PidInfo() Info {
	return SimpleInfo("pid", os.Getpid())
}

func HostnameInfo() Info {
	hostname, e := os.Hostname()
	if e != nil {
		panic(e)
	}

	return SimpleInfo("hostname", hostname)
}

func TimeInfo() Info {
	return InfoFunc("time", func() interface{} {
		return time.Now().Format(time.RFC3339)
	})
}
