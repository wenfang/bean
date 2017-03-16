package kproducer

import (
	"github.com/Shopify/sarama"
)

// KProducer 结构
type KProducer struct {
	sarama.SyncProducer
	topic string
}

// New 创建KProducer
func New(brokers []string, topic string) (*KProducer, error) {
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		return nil, err
	}
	return &KProducer{SyncProducer: producer, topic: topic}, nil
}

// SyncSendMessage 发送消息
func (kp *KProducer) SyncSendMessage(key, value string) (int32, int64, error) {
	// set producer message
	msg := &sarama.ProducerMessage{
		Topic: kp.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	var (
		partition int32
		offset    int64
		err       error
	)
	// send message
	if partition, offset, err = kp.SendMessage(msg); err != nil {
		return 0, 0, err
	}
	return partition, offset, nil
}
