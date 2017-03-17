package sqls

import (
	"fmt"
)

// Conds 代表WHERE条件
type Conds struct {
	Fields
	Values map[string][]interface{}
}

func (conds Conds) parse() (string, []interface{}) {
	sql := "WHERE"
	var args []interface{}

	for idx, field := range conds.Fields {
		length := len(conds.Values[field])
		if length == 0 {
			continue
		}
		// 添加参数
		args = append(args, conds.Values[field]...)
		// 只有一个可选值用=
		if length == 1 {
			if idx == 0 {
				sql += fmt.Sprintf(" %s=?", field)
			} else {
				sql += fmt.Sprintf(" AND %s=?", field)
			}
			continue
		}
		// 多于一个可选值用IN
		if idx == 0 {
			sql += fmt.Sprintf(" %s IN (", field)
		} else {
			sql += fmt.Sprintf(" AND %s IN (", field)
		}

		for i := 0; i < length; i++ {
			if i == 0 {
				sql += "?"
			} else {
				sql += ", ?"
			}
		}
		sql += ")"
	}
	return sql, args
}
