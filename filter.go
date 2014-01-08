package bunyan

import "os"
import "time"

type Filter interface{
	Apply(record Record) bool
}

type filterImpl struct{
	apply func(record Record) bool
}

func (f *filterImpl) Apply(record Record) bool {
	return f.apply(record)
}

func FilterFunc(f func(record Record) bool) Filter {
	return &filterImpl{f}
}

func SimpleFilter(key string, value interface{}) Filter {
	return FilterFunc(func(r Record) bool {
		r[key] = value
		return true
	})
}

func PidFilter() Filter {
	return SimpleFilter("pid", os.Getpid())
}

func HostnameFilter() Filter {
	hostname, e := os.Hostname()
	if e != nil {
		panic(e)
	}

	return SimpleFilter("hostname", hostname)
}

func TimeFilter() Filter {
	return FilterFunc(func(record Record) bool {
		record["time"] = time.Now().Format(time.RFC3339)
		return true
	})
}

func LevelFilter(level int) Filter {
	return FilterFunc(func(record Record) bool {
		lvl, ok := record["level"].(int)
		return ok && lvl >= level
	})
}

func MultiFilter(filters...Filter) Filter {
	return FilterFunc(func(record Record) bool {
		result := true
		for _, filter := range filters {
			result_ := filter.Apply(record) // don't short-circuit
			result = result && result_
		}
		return result
	})
}

func StdFilter(name string) Filter {
	return MultiFilter(
		SimpleFilter("v", 0),
		SimpleFilter("name", name),
		PidFilter(),
		HostnameFilter(),
		TimeFilter(),
	)
}
