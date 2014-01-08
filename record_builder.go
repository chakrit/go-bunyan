package bunyan

type RecordBuilder struct{
	template Record
	record Record
	sink Sink
}

func NewRecordBuilder(target Sink) *RecordBuilder {
	return nil
}
