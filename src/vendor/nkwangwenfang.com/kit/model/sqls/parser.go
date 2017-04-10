package sqls

import (
	"fmt"
)

type Parser interface {
	ParseSQL() (string, error)
	Parse() (string, []interface{}, error)
}

type logic struct {
	typ     string
	parsers []Parser
}

func (l *logic) ParseSQL() (string, error) {
	sql := "("
	for idx, cond := range l.parsers {
		if idx != 0 {
			sql += fmt.Sprintf(" %s ", l.typ)
		}
		condSql, err := cond.ParseSQL()
		if err != nil {
			return "", err
		}
		sql += condSql
	}
	sql += ")"
	return sql, nil
}

func (l *logic) Parse() (string, []interface{}, error) {
	var (
		sql  string
		args []interface{}
	)
	sql = "("
	for idx, parser := range l.parsers {
		if idx != 0 {
			sql += fmt.Sprintf(" %s ", l.typ)
		}
		condSql, condArgs, err := parser.Parse()
		if err != nil {
			return "", nil, err
		}
		sql += condSql
		args = append(args, condArgs...)
	}
	sql += ")"
	return sql, args, nil
}

func LogicAND(parsers ...Parser) Parser {
	return &logic{typ: "AND", parsers: parsers}
}

func LogicOR(parsers ...Parser) Parser {
	return &logic{typ: "OR", parsers: parsers}
}
