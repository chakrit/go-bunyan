package bunyan

type Filter interface{
	Apply(record Record)
}
