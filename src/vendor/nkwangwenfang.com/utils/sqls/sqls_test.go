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
	conds := Conds{
		Fields: Fields{"id", "pop_id"},
		Values: map[string][]interface{}{
			"id":       {10, 20},
			"pop_id":   {30},
			"store_id": {},
		},
	}

	t.Log(SelectSQL("mytable", selects, conds))
	t.Log(InsertSQL("mytable", inserts))
	t.Log(InsertUpdateSQL("mytable", inserts, updates))
	t.Log(UpdateSQL("mytable", updates, conds))
	t.Log(DeleteSQL("mytable", conds))
}
