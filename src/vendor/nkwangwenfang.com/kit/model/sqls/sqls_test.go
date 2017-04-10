package sqls

import (
	"testing"
)

func TestSQL(t *testing.T) {
	selects := Fields{"name", "age"}
	inserts := FieldsValues{
		Fields: Fields{"id", "pop_id", "store_id"},
		Values: []interface{}{10, 20, 30},
	}
	updates := FieldsValues{
		Fields: Fields{"pop_id", "store_id"},
		Values: []interface{}{10, 100},
	}
	cond1 := Cond{Field: "id", Value: 10, Typ: TypEQ}
	cond2 := Cond{Field: "pop_id", Value: 20, Typ: TypNE}
	cond3 := Cond{Field: "store_id", Value: 30, Typ: TypEQ}
	conds := LogicAND(cond1, cond2, cond3)

	t.Log(SelectSQL("mytable", selects, conds))
	t.Log(InsertSQL("mytable", inserts))
	t.Log(InsertUpdateSQL("mytable", inserts, updates))
	t.Log(UpdateSQL("mytable", updates, conds))
	t.Log(DeleteSQL("mytable", conds))
}

func TestSQL2(t *testing.T) {
	cond1 := Cond{Field: "id", Value: 10, Typ: TypEQ}
	cond2 := Cond{Field: "pop_id", Value: 20, Typ: TypNE}
	cond3 := Cond{Field: "store_id", Value: 30, Typ: TypEQ}
	sql, err := LogicAND(LogicOR(cond1, cond2), cond3).ParseSQL()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
