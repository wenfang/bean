package rdsmaster

import (
	"math/rand"
	"testing"

	"nkwangwenfang.com/cluster/master"
	"nkwangwenfang.com/rds/codis"
)

var codisConfig = codis.Config{
	Addrs:          []string{"127.0.0.1:6379"},
	MaxIdle:        20,
	MaxActive:      50,
	ConnectTimeout: 3000,
	ReadTimeout:    3000,
	WriteTimeout:   3000,
	IdleTimeout:    30000,
	Wait:           true,
}

func TestRdsMaster(t *testing.T) {
	r := codis.New(codisConfig)
	defer r.Close()

	num := rand.Intn(10)
	masters := make([]master.Master, num)
	for i := 0; i < num; i++ {
		masters[i] = New(r, "test")
	}
	for i := 0; i < num; i++ {
		ok, err := masters[i].CheckMaster()
		if err != nil {
			t.Fatal(err)
		}
		if ok {
			t.Log("get master")
		}
	}
}
