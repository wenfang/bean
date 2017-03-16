package module

import (
	"testing"
)

func TestSQL(t *testing.T) {
	selectFields := []string{"name", "age"}
	insertFields := []string{"id", "pop_id", "store_id"}
	updateFields := []string{"pop_id", "store_id"}
	updateValues := []interface{}{10, 100}
	whereFields := []string{"id", "pop_id"}
	whereValues := map[string][]interface{}{
		"id":       {10, 20},
		"pop_id":   {30},
		"store_id": {},
	}

	t.Log(SelectSQL("mytable", selectFields, whereFields, whereValues))
	t.Log(InsertSQL("mytable", insertFields))
	t.Log(InsertUpdateSQL("mytable", insertFields, updateFields))
	t.Log(UpdateSQL("mytable", updateFields, updateValues, whereFields, whereValues))
	t.Log(DeleteSQL("mytable", whereFields, whereValues))
}
