package kconsumer

import (
	"errors"

	"github.com/Shopify/sarama"
	"golang.org/x/net/context"

	"nkwangwenfang.com/kafka/kmessage"
	"nkwangwenfang.com/kafka/koffset"
)

// KConsumer 结构
type KConsumer struct {
	consumer   sarama.Consumer
	topic      string
	partitions []int32
	offsetter  koffset.KOffsetter
}

// New 创建KConsumer
func New(brokers []string, topic string, offsetter koffset.KOffsetter) (*KConsumer, error) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}
	// 获取Partition ID
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		consumer.Close()
		return nil, err
	}
	return &KConsumer{consumer: consumer, topic: topic, partitions: partitions, offsetter: offsetter}, nil
}

// Partitions 返回Partition ID列表
func (kc *KConsumer) Partitions() []int32 {
	return kc.partitions
}

// Close 关闭KConsumer
func (kc *KConsumer) Close() {
	kc.consumer.Close()
}

// PartitionConsumer 创建KPartitionConsumer
func (kc *KConsumer) PartitionConsumer(partitionID int32) (*KPartitionConsumer, error) {
	// 获取当前偏移
	offset, err := kc.offsetter.Get(partitionID)
	if err != nil {
		return nil, err
	}
	// 校正offset
	if offset <= 0 {
		offset = sarama.OffsetOldest
	} else {
		offset = offset + 1
	}
	// 创建partitionConsumer
	partitionConsumer, err := kc.consumer.ConsumePartition(kc.topic, partitionID, offset)
	if err != nil {
		// 如果发生错误，重新打开最老的Offset
		if partitionConsumer, err = kc.consumer.ConsumePartition(kc.topic, partitionID, sarama.OffsetOldest); err != nil {
			return nil, err
		}
	}
	return &KPartitionConsumer{KOffsetter: kc.offsetter, partitionConsumer: partitionConsumer}, nil
}

var (
	// ErrNilMessage 空消息
	ErrNilMessage = errors.New("nil message")
	// ErrDone 执行GetMessage完成
	ErrDone = errors.New("get message done")
)

// KPartitionConsumer 结构
type KPartitionConsumer struct {
	koffset.KOffsetter
	partitionConsumer sarama.PartitionConsumer
}

// GetMessage 获取一个kafka消息
func (kpc *KPartitionConsumer) GetMessage(ctx context.Context) (*kmessage.KMessage, error) {
	select {
	case msg, ok := <-kpc.partitionConsumer.Messages():
		if ok {
			return kmessage.New(msg, kpc.KOffsetter), nil
		}
		return nil, ErrNilMessage
	case <-ctx.Done():
		return nil, ErrDone
	}
}

// Close 关闭KPartitionConsumer
func (kpc *KPartitionConsumer) Close() {
	kpc.partitionConsumer.Close()
}
