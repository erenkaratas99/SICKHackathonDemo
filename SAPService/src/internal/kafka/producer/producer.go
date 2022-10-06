package producer

import (
	"SICKHackathon/SAPService/src/internal/kafka/internalConfig"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
	"github.com/labstack/gommon/log"
	"net/http"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() *KafkaProducer {
	kcfg := internalConfig.GetKafkaConfig()
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kcfg.BootstrapServers,
		"security.protocol": kcfg.SecurityProtocol,
		"sasl.mechanisms":   kcfg.SaslMech,
		"sasl.username":     kcfg.SaslUsername,
		"sasl.password":     kcfg.SaslPassword,
	})
	if err != nil {
		panic(err)
	}
	return &KafkaProducer{
		producer: p,
	}
}

func (kp *KafkaProducer) ProduceMsg(topic string, msg []byte) error {
	go func() {
		for e := range kp.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					errMsg := fmt.Sprintf("Delivery failed: %v\n", ev.TopicPartition)
					customErrors.NewHTTPError(http.StatusInternalServerError, "KafkaProducerErr", errMsg)
				} else {
					log.Info("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)

	// Wait for message deliveries before shutting down
	kp.producer.Flush(15 * 1000)
	kp.producer.Close()
	return nil
}
