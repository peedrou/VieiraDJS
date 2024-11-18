package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %v", err)
	}

	return &KafkaProducer{
		producer: producer,
	}, nil
}

func (kp *KafkaProducer) Close() error {
	return kp.producer.Close()
}

func (kp *KafkaProducer) SendMessage(topic string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := kp.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	return nil
}
