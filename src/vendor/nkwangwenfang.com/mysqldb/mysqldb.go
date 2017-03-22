package mysqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB 对mysql db的封装
type DB struct {
	*sql.DB
}

// Open 打开一个mysql db
func Open(config Config) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&timeout=5s&readTimeout=30s",
		config.User,
		config.Passwd,
		config.Addr,
		config.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
