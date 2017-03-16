package rdslock

import (
	"testing"
	"time"

	"nkwangwenfang.com/rds/codis"
)

func TestRdsLock(t *testing.T) {
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

	lock1 := New(r, "locktest", HoldTime(3000))
	lock2 := New(r, "locktest", HoldTime(3000))
	if ok, err := lock1.Lock(); !ok {
		t.Fatal("lock1 lock error", ok, err)
	}

	if ok, err := lock2.Lock(); ok || err != nil {
		t.Fatal("lock2 lock error", ok, err)
	}

	if ok, err := lock1.Lock(); !ok {
		t.Fatal("lock1 update error", ok, err)
	}

	if ok, err := lock1.Unlock(); !ok {
		t.Fatal("lock1 unlock error", ok, err)
	}

	if ok, err := lock2.Lock(); !ok {
		t.Fatal("lock2 lock error", ok, err)
	}

	time.Sleep(5 * time.Second)

	if ok, err := lock1.Lock(); !ok {
		t.Fatal("lock1 lock error", ok, err)
	}

	if ok, err := lock2.Unlock(); ok || err != nil {
		t.Fatal("lock2 unlock error", ok, err)
	}

	if ok, err := lock1.Unlock(); !ok {
		t.Fatal("lock1 unlock error", ok, err)
	}
}
