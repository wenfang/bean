package rdskoffset

import (
	"testing"

	"nkwangwenfang.com/rds/codis"
)

func TestPartitionOffset(t *testing.T) {
	config := codis.Config{
		Addrs:          []string{"127.0.0.1:6379"},
		MaxIdle:        20,
		MaxActive:      50,
		ConnectTimeout: 3000,
		ReadTimeout:    3000,
		WriteTimeout:   3000,
		IdleTimeout:    30000,
		Wait:           true,
	}
	r := codis.New(config)
	defer r.Close()

	app := "testApp"
	ro := New(app, r)

	tryCase(t, ro, testCase{partition: 21, offset: 0})
	tryCase(t, ro, testCase{partition: 21, offset: 23423})
}
