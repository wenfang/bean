package kproducer

import (
	"fmt"
	"testing"
	"time"
)

func TestDProducer(t *testing.T) {
	brokers := []string{"127.0.0.1:9090", "127.0.0.1:9091", "127.0.0.1:9092"}
	topic := "orders_realtime"

	kp, err := New(brokers, topic)
	if err != nil {
		t.Fatal(err)
	}
	defer kp.Close()

	for i := 0; i < 100; i++ {
		_, _, err := kp.SyncSendMessage(fmt.Sprintf("%d", i), fmt.Sprintf("Message_%d_%d", time.Now().Unix(), i))
		if err != nil {
			t.Fatal(err)
		}
	}
}
