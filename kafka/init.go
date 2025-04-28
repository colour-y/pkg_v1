package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

type msgConsumerGroup struct {
}

func (m msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error { return nil }
func (m msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (m msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}

var ConsumerGroup msgConsumerGroup

func InitSynProducer(brokers []string) sarama.SyncProducer {
	conf := sarama.NewConfig()
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = time.Second * 1
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Metadata.Full = true
	conf.Version = sarama.V0_10_2_0
	conf.Metadata.Full = true

	producer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		panic(err.Error())
	}
	return producer
}

func InitConsumerGroup(brokers []string, groupId string) sarama.ConsumerGroup {
	conf := sarama.NewConfig()
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = time.Second * 1
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Metadata.Full = true
	conf.Version = sarama.V0_10_2_0
	// conf.ClientID = "sasl_scram_client"
	conf.Metadata.Full = true
	// conf.Net.SASL.Enable = true
	// conf.Net.SASL.User = *userName
	// conf.Net.SASL.Password = *passwd
	// conf.Net.SASL.Handshake = true

	cGroup, err := sarama.NewConsumerGroup(brokers, groupId, conf)
	if err != nil {
		panic(err)
	}
	return cGroup
}
