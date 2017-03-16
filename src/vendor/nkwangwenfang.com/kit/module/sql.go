package module

import (
	"fmt"
)

func whereCons(fields []string, values map[string][]interface{}) (string, []interface{}) {
	sql := "WHERE "
	var args []interface{}

	for idx, field := range fields {
		length := len(values[field])

		if length == 0 {
			continue
		}

		args = append(args, values[field]...)

		if length == 1 {
			if idx == 0 {
				sql += fmt.Sprintf("%s=?", field)
			} else {
				sql += fmt.Sprintf(" AND %s=?", field)
			}
			continue
		}

		if idx == 0 {
			sql += fmt.Sprintf("%s IN (", field)
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

func SelectSQL(tableName string, selectFields []string, whereFields []string, whereValues map[string][]interface{}) (string, []interface{}) {
	sql := fmt.Sprintf("SELECT ")
	var args []interface{}

	if len(selectFields) == 0 {
		sql += "*"
	} else {
		for idx, selectField := range selectFields {
			if idx == 0 {
				sql += fmt.Sprintf("%s", selectField)
			} else {
				sql += fmt.Sprintf(", %s", selectField)
			}
		}
	}
	sql += fmt.Sprintf(" FROM %s", tableName)

	whereSql, whereArgs := whereCons(whereFields, whereValues)
	if len(whereArgs) > 0 {
		sql += " " + whereSql + " AND _deleted=0"
		args = append(args, whereArgs...)
	}

	return sql, args
}

func InsertSQL(tableName string, insertFields []string) string {
	sql := fmt.Sprintf("INSERT INTO %s SET ", tableName)
	for idx, insertField := range insertFields {
		if idx == 0 {
			sql += fmt.Sprintf("%s=?", insertField)
		} else {
			sql += fmt.Sprintf(", %s=?", insertField)
		}
	}
	sql += ", _create_time=NOW()"
	return sql
}

func InsertUpdateSQL(tablename string, insertFields []string, updateFields []string) string {
	sql := InsertSQL(tablename, insertFields)
	sql += " ON DUPLICATE KEY UPDATE "
	for idx, updateField := range updateFields {
		if idx == 0 {
			sql += fmt.Sprintf("%s=?", updateField)
		} else {
			sql += fmt.Sprintf(", %s=?", updateField)
		}
	}
	return sql
}

func UpdateSQL(tableName string, updateFields []string, updateValues []interface{}, whereFields []string, whereValues map[string][]interface{}) (string, []interface{}) {
	args := updateValues

	sql := fmt.Sprintf("UPDATE %s SET ", tableName)
	for idx, updateField := range updateFields {
		if idx == 0 {
			sql += fmt.Sprintf("%s=?", updateField)
		} else {
			sql += fmt.Sprintf(", %s=?", updateField)
		}
	}

	whereSql, whereArgs := whereCons(whereFields, whereValues)
	if len(whereArgs) == 0 { // 不允许没有WHERE的UPDATE
		return "", nil
	}
	sql += " " + whereSql
	args = append(args, whereArgs...)
	return sql, args
}

func DeleteSQL(tableName string, whereFields []string, whereValues map[string][]interface{}) (string, []interface{}) {
	var args []interface{}
	sql := fmt.Sprintf("UPDATE %s SET _deleted=1", tableName)

	whereSql, whereArgs := whereCons(whereFields, whereValues)
	if len(whereArgs) == 0 { // 不允许没有WHERE的UPDATE
		return "", nil
	}
	sql += " " + whereSql
	args = append(args, whereArgs...)

	return sql, args
}
