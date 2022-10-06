package consumer

import (
	"SICKHackathon/responseService/src/internal/kafka/internalConfig"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumerHandler struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumerHandler() *KafkaConsumerHandler {
	kcfg := internalConfig.GetKafkaConfig()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kcfg.BootstrapServers,
		"security.protocol": kcfg.SecurityProtocol,
		"sasl.mechanisms":   kcfg.SaslMech,
		"sasl.username":     kcfg.SaslUsername,
		"sasl.password":     kcfg.SaslPassword,
		"group.id":          kcfg.GroupID,
		"auto.offset.reset": kcfg.AutoOffsetReset,
	})

	if err != nil {
		panic(err)
	}
	return &KafkaConsumerHandler{consumer: c}
}

func (kc *KafkaConsumerHandler) ConsumeKafkaTopic() error {

	kc.consumer.SubscribeTopics([]string{"topic_0", "^aRegex.*[Tt]opic"}, nil)
	for {
		msg, err := kc.consumer.ReadMessage(-1)
		if err != nil {
			return err
		}
		fmt.Println(msg)
		return nil
	}
	kc.consumer.Close()

	return nil
}
