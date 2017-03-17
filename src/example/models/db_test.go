package db

import (
	"testing"
)

var config = Config{
	Addr:   "127.0.0.1:3306",
	User:   "root",
	Passwd: "",
	DBName: "test",
}

func TestStatusDB(t *testing.T) {
	statusDB, err := New(config)
	if err != nil {
		t.Fatal(err)
	}
	defer statusDB.Close()

	err = statusDB.Set(123456, []string{"mytest"}, []interface{}{1})
	if err != nil {
		t.Fatal(err)
	}
}
