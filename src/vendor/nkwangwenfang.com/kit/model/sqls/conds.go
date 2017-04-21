package sqls

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	TypEQ = iota // =
	TypGT        // >
	TypGE        // >=
	TypLT        // <
	TypLE        // <=
	TypNE        // !=
	TypLK        // LIKE
)

// Cond 代表一个SQL查询条件
type Cond struct {
	Field string
	Value interface{}
	Typ   int
}

func (cond Cond) ParseSQL() (string, error) {
	// 处理类型为string的情况
	if reflect.TypeOf(cond.Value).Kind() == reflect.String {
		switch cond.Typ {
		case TypEQ:
			return fmt.Sprintf("%s = '%v'", cond.Field, cond.Value), nil
		case TypLK:
			return fmt.Sprintf("%s LIKE '%v'", cond.Field, cond.Value), nil
		}
		return "", errors.New("condition type error")
	}
	// 处理其他数值情况
	switch cond.Typ {
	case TypEQ:
		return fmt.Sprintf("%s = %v", cond.Field, cond.Value), nil
	case TypGT:
		return fmt.Sprintf("%s > %v", cond.Field, cond.Value), nil
	case TypGE:
		return fmt.Sprintf("%s >= %v", cond.Field, cond.Value), nil
	case TypLT:
		return fmt.Sprintf("%s < %v", cond.Field, cond.Value), nil
	case TypLE:
		return fmt.Sprintf("%s <= %v", cond.Field, cond.Value), nil
	case TypNE:
		return fmt.Sprintf("%s != %v", cond.Field, cond.Value), nil
	}
	return "", errors.New("condition type error")
}

func (cond Cond) Parse() (string, []interface{}, error) {
	switch cond.Typ {
	case TypEQ:
		return fmt.Sprintf("%s = ?", cond.Field), []interface{}{cond.Value}, nil
	case TypGT:
		return fmt.Sprintf("%s > ?", cond.Field), []interface{}{cond.Value}, nil
	case TypGE:
		return fmt.Sprintf("%s >= ?", cond.Field), []interface{}{cond.Value}, nil
	case TypLT:
		return fmt.Sprintf("%s < ?", cond.Field), []interface{}{cond.Value}, nil
	case TypLE:
		return fmt.Sprintf("%s <= ?", cond.Field), []interface{}{cond.Value}, nil
	case TypNE:
		return fmt.Sprintf("%s != ?", cond.Field), []interface{}{cond.Value}, nil
	case TypLK:
		return fmt.Sprintf("%s LIKE ?", cond.Field), []interface{}{cond.Value}, nil
	}
	return "", nil, errors.New("condition type error")
}
