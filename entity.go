package pgxsqlbuilder

type Field struct {
	name  string
	value interface{}
}

func NewFiled(name string, value interface{}) *Field {
	return &Field{
		name:  name,
		value: value,
	}
}

type Query struct {
	sql    string
	values []interface{}
}
