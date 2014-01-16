package bunyan

import "fmt"
import "encoding/json"

// Record represents an individual log record. The underlying type is simply a map with
// string key (since JSON does not supports non-string key). You can builds the Record
// manually and pass to the logger's Write() method to writes it or use one of the
// logger's Record() or Include() method to modify.
type Record map[string]interface{}

const InitialRecordCapacity = 16

// NewRecord() creates a new, empty record with initial capacity sets to the
// InitialRecordCapacity constant.
func NewRecord() Record {
	m := make(map[string]interface{}, InitialRecordCapacity)
	return Record(m)
}

// NewSimpleRecord() creates a new record and immediately adds the given key and value
// pair to it.
func NewSimpleRecord(key string, value interface{}) Record {
	record := NewRecord()
	record[key] = value
	return record
}

// UnmarshalRecord() re-creates a Record from a JSON string.
func UnmarshalRecord(bytes []byte) (Record, error) {
	result := make(map[string]interface{})
	e := json.Unmarshal(bytes, &result)
	if e != nil {
		return nil, e
	}

	return Record(result), nil
}

// Marshal() translates the Record into its JSON representation.
func (r Record) Marshal() ([]byte, error) {
	result, e := json.Marshal(r)
	return result, e
}

// SetIfNot() sets the given value to the given key if and only if the key does not
// already exists in the record. Useful for providing record templates that maybe
// overridden.
func (r Record) SetIfNot(key string, value interface{}) Record {
	if _, ok := r[key]; !ok {
		r[key] = value
	}

	return r
}

// TemplateMerge() merges the given record into the receiver record. Existing keys in the
// receiver record will not be overridden.
//
// SetIfNot() is used internally to transfer key and values.
func (r Record) TemplateMerge(template Record) Record {
	for k, v := range template {
		r.SetIfNot(k, v)
	}

	return r
}

// SetMessagef() is a convenience method for setting the standard level and msg keys that
// should be present on all records.
func (r Record) SetMessagef(level Level, msg string, args ...interface{}) {
	r["level"], r["msg"] = level, fmt.Sprintf(msg, args...)
}
