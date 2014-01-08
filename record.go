package bunyan

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

func MergeRecord(records...Record) Record {
	result := NewRecord()
	for _, record := range records {
		for k, v := range record {
			result[k] = v
		}
	}

	return result
}
