package kmessage

import (
	"github.com/Shopify/sarama"

	"nkwangwenfang.com/kafka/koffset"
)

// KMessage 消息结构
type KMessage struct {
	*sarama.ConsumerMessage
	offsetter koffset.KOffsetter
}

// New 创建KMessage
func New(msg *sarama.ConsumerMessage, offsetter koffset.KOffsetter) *KMessage {
	return &KMessage{ConsumerMessage: msg, offsetter: offsetter}
}

// Commit 同步该消息的offset到offsetter
func (km *KMessage) Commit() {
	km.offsetter.Set(km.Partition, km.Offset)
}
