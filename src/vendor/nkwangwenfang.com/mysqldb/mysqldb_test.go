package mysqldb

import (
	"testing"
)

var config = Config{
	Addr:   "127.0.0.1:3306",
	User:   "root",
	DBName: "core",
}

func TestDB(t *testing.T) {
	db, err := Open(config)
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	db.Close()
}
