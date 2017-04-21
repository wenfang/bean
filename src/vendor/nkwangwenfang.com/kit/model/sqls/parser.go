package sqls

// Parser 条件解析接口
type Parser interface {
	ParseSQL() (string, error)
	Parse() (string, []interface{}, error)
}
