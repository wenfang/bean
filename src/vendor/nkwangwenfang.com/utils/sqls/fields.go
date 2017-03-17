package sqls

// Fields 代表字段列表
type Fields []string

// FieldsValues 代表字段及对应的值的列表
type FieldsValues struct {
	Fields
	Values []interface{}
}
