package sqls

import (
	"fmt"
)

func SelectSQL(tableName string, selects Fields, conds Conds) (string, []interface{}) {
	sql := fmt.Sprintf("SELECT")

	if len(selects) == 0 {
		sql += " *"
	} else {
		for idx, field := range selects {
			if idx == 0 {
				sql += fmt.Sprintf(" %s", field)
			} else {
				sql += fmt.Sprintf(", %s", field)
			}
		}
	}
	sql += fmt.Sprintf(" FROM %s", tableName)

	whereSql, whereArgs := conds.parse()
	if len(whereArgs) == 0 { // 不允许没有WHERE的SELECT
		return "", nil
	}
	sql += " " + whereSql

	return sql, whereArgs
}

func InsertSQL(tableName string, inserts FieldsValues) (string, []interface{}) {
	sql := fmt.Sprintf("INSERT INTO %s SET", tableName)
	for idx, field := range inserts.Fields {
		if idx == 0 {
			sql += fmt.Sprintf(" %s=?", field)
		} else {
			sql += fmt.Sprintf(", %s=?", field)
		}
	}
	return sql, inserts.Values
}

func InsertUpdateSQL(tableName string, inserts, updates FieldsValues) (string, []interface{}) {
	sql, args := InsertSQL(tableName, inserts)
	sql += " ON DUPLICATE KEY UPDATE"
	for idx, field := range updates.Fields {
		if idx == 0 {
			sql += fmt.Sprintf(" %s=?", field)
		} else {
			sql += fmt.Sprintf(", %s=?", field)
		}
	}
	args = append(args, updates.Values...)
	return sql, args
}

func UpdateSQL(tableName string, updates FieldsValues, conds Conds) (string, []interface{}) {
	sql := fmt.Sprintf("UPDATE %s SET", tableName)
	for idx, field := range updates.Fields {
		if idx == 0 {
			sql += fmt.Sprintf(" %s=?", field)
		} else {
			sql += fmt.Sprintf(", %s=?", field)
		}
	}

	whereSql, whereArgs := conds.parse()
	if len(whereArgs) == 0 { // 不允许没有WHERE的UPDATE
		return "", nil
	}
	sql += " " + whereSql
	args := append(updates.Values, whereArgs...)
	return sql, args
}

func DeleteSQL(tableName string, conds Conds) (string, []interface{}) {
	sql := fmt.Sprintf("DELETE FROM %s", tableName)

	whereSql, whereArgs := conds.parse()
	if len(whereArgs) == 0 { // 不允许没有WHERE的DELETE
		return "", nil
	}
	sql += " " + whereSql

	return sql, whereArgs
}
