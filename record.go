package bunyan

import "fmt"
import "encoding/json"

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

func UnmarshalRecord(bytes []byte) (Record, error) {
	result := make(map[string]interface{})
	e := json.Unmarshal(bytes, &result)
	if e != nil {
		return nil, e
	}

	return Record(result), nil
}

func (r Record) Marshal() ([]byte, error) {
	result, e := json.Marshal(r)
	return result, e
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

func (r Record) SetMessagef(level Level, msg string, args...interface{}) {
	r["level"], r["msg"] = level, fmt.Sprintf(msg, args...)
}
