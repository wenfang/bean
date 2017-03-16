package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"nkwangwenfang.com/log"
)

const tableSize = 1

// StatusDB 状态数据库结构
type StatusDB struct {
	db *sql.DB
}

// New 创建StatusDB
func New(config Config) (*StatusDB, error) {
	var (
		s   StatusDB
		err error
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&timeout=5s&readTimeout=30s",
		config.User,
		config.Passwd,
		config.Addr,
		config.DBName)

	if s.db, err = sql.Open("mysql", dsn); err != nil {
		log.Error("mysql open error", "dsn", dsn, "error", err)
		return nil, err
	}
	return &s, nil
}

func getSQL(ID int64, fields []string) string {
	sql := fmt.Sprintf("INSERT INTO es_order_status_%d SET id=?", ID%tableSize)
	for _, field := range fields {
		sql += fmt.Sprintf(", %s=?", field)
	}
	sql += ", _create_time=NOW() ON DUPLICATE KEY UPDATE "
	for idx, field := range fields {
		if idx != 0 {
			sql += fmt.Sprintf(", %s=?", field)
		} else {
			sql += fmt.Sprintf("%s=?", field)
		}
	}
	return sql
}

// Set StatusDB设置状态
func (s *StatusDB) Set(ID int64, fields []string, args []interface{}) error {
	if len(args) == 0 {
		return nil
	}

	sql := getSQL(ID, fields)

	var allArgs []interface{}
	allArgs = append(allArgs, ID)
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, args...)

	if _, err := s.db.Exec(sql, allArgs...); err != nil {
		log.Error("mysql exec error", "error", err)
		return errSetDB
	}
	return nil
}

// Close 关闭StatusDB
func (s *StatusDB) Close() {
	s.db.Close()
}
