package bunyan

import "fmt"

type Record map[string]interface{}

const InitialRecordCapacity = 16

func NewRecord() Record {
	m := make(map[string]interface{}, InitialRecordCapacity)
	return Record(m)
}

func NewSimpleRecord(key string, value interface{}) Record {
	record := NewRecord()
	record[key] = value
	return record
}

func (r Record) TemplateMerge(template Record) Record {
	for k, v := range template {
		if _, ok := r[k]; !ok {
			r[k] = v
		}
	}

	return r
}

func (r Record) SetMessagef(level int, msg string, args...interface{}) {
	r["level"], r["msg"] = level, fmt.Sprintf(msg, args...)
}
