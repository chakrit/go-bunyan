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

func (r Record) SetIfNot(key string, value interface{}) Record {
	if _, ok := r[key]; !ok {
		r[key] = value
	}

	return r
}

func (r Record) TemplateMerge(template Record) Record {
	for k, v := range template {
		r.SetIfNot(k, v)
	}

	return r
}

func (r Record) SetMessagef(level int, msg string, args...interface{}) {
	r["level"], r["msg"] = level, fmt.Sprintf(msg, args...)
}
