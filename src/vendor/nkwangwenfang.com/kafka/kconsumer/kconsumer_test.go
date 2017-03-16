package kconsumer

import (
	"sync"
	"testing"

	"golang.org/x/net/context"

	"nkwangwenfang.com/kafka/koffset/rdskoffset"
	"nkwangwenfang.com/rds/codis"
)

var wg sync.WaitGroup

func getMessage(t *testing.T, kpc *KPartitionConsumer) {
	defer wg.Done()

	ctx, _ := context.WithTimeout(context.Background(), 3e9)
	for {
		msg, err := kpc.GetMessage(ctx)
		if err == ErrDone {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		msg.Commit()
	}
}

func TestCons(t *testing.T) {
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

	app := "test"
	off := rdskoffset.New(app, r)

	brokers := []string{"wanliu01.jx.diditaxi.com:9092", "wanliu02.jx.diditaxi.com:9092", "wanliu03.jx.diditaxi.com:9092"}
	topic := "wanliu_order_basic"

	kc, err := New(brokers, topic, off)
	if err != nil {
		t.Fatal(err)
	}
	defer kc.Close()
	t.Fatal("new broker")

	for _, partitionID := range kc.Partitions() {
		kpc, err := kc.PartitionConsumer(partitionID)
		if err != nil {
			t.Fatal(err)
		}
		wg.Add(1)
		go getMessage(t, kpc)
	}
	wg.Wait()
}
