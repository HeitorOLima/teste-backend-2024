package KafkaProducer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	producerInstance *KafkaProducer
	once             sync.Once
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		MaxAttempts:            10,
		WriteTimeout:           10 * time.Second,
		ReadTimeout:            10 * time.Second,
		RequiredAcks:           kafka.RequireAll,
	}
	return &KafkaProducer{writer: writer}
}

func GetProducer() *KafkaProducer {
	once.Do(func() {
		producerInstance = NewProducer([]string{"localhost:9092"}, "go-to-rails")
	})
	return producerInstance
}

func (producer *KafkaProducer) Close() {
	producer.writer.Close()
}

func (producer *KafkaProducer) ProduceMessage(key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := producer.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("Erro ao produzir mensagem: %v", err)
	} else {
		log.Printf("Mensagem produzida: %s", value)
	}

	return err
}
