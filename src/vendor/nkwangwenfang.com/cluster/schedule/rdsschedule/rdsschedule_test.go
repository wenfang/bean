package rdsschedule

import (
	"testing"

	"nkwangwenfang.com/cluster/schedule"
	"nkwangwenfang.com/rds/codis"
)

var config = codis.Config{
	Addrs:          []string{"127.0.0.1:6379"},
	MaxIdle:        20,
	MaxActive:      50,
	ConnectTimeout: 3000,
	ReadTimeout:    3000,
	WriteTimeout:   3000,
	IdleTimeout:    30000,
	Wait:           true,
}

func TestRdsInfo(t *testing.T) {
	r := codis.New(config)
	defer r.Close()

	i := New(r, "test")

	info := make(schedule.Info)
	info["node1"] = []int{1, 2}
	info["node2"] = []int{3, 4, 5}
	if err := i.SetInfo(info); err != nil {
		t.Fatal(err)
	}
	if _, err := i.GetInfo(); err != nil {
		t.Fatal(err)
	}
}
