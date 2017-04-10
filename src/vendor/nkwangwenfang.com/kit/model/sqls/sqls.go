package sqls

import (
	"errors"
	"fmt"
	"strings"
)

func SelectSQL(tableName string, selects Fields, conds Parser) (string, []interface{}, error) {
	sql := fmt.Sprintf("SELECT ")

	if len(selects) == 0 {
		sql += "*"
	} else {
		sql += strings.Join(selects, ", ")
	}
	sql += fmt.Sprintf(" FROM %s", tableName)

	whereSql, whereArgs, err := conds.Parse()
	if err != nil {
		return "", nil, err
	}
	if len(whereArgs) == 0 { // 不允许没有WHERE的SELECT
		return "", nil, errors.New("no where condition set")
	}
	sql += " WHERE " + whereSql

	return sql, whereArgs, nil
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

func UpdateSQL(tableName string, updates FieldsValues, conds Parser) (string, []interface{}, error) {
	sql := fmt.Sprintf("UPDATE %s SET", tableName)
	for idx, field := range updates.Fields {
		if idx == 0 {
			sql += fmt.Sprintf(" %s=?", field)
		} else {
			sql += fmt.Sprintf(", %s=?", field)
		}
	}

	whereSql, whereArgs, err := conds.Parse()
	if err != nil {
		return "", nil, err
	}
	if len(whereArgs) == 0 { // 不允许没有WHERE的UPDATE
		return "", nil, errors.New("no where condition set")
	}
	sql += " WHERE " + whereSql
	args := append(updates.Values, whereArgs...)

	return sql, args, nil
}

func DeleteSQL(tableName string, conds Parser) (string, []interface{}, error) {
	sql := fmt.Sprintf("DELETE FROM %s", tableName)

	whereSql, whereArgs, err := conds.Parse()
	if err != nil {
		return "", nil, err
	}
	if len(whereArgs) == 0 { // 不允许没有WHERE的DELETE
		return "", nil, errors.New("no where condition set")
	}
	sql += " WHERE " + whereSql

	return sql, whereArgs, nil
}
