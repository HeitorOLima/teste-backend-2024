package KafkaProducer

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	PRODUCER_TOPIC    = "go-to-rails"
	DEFAULT_PARTITION = 0
	BROKER_ADDRESS    = "localhost:9092"
)

type Producer struct {
	Conn *kafka.Conn
}

func NewProducer() *Producer {
	conn, err := kafka.DialLeader(context.Background(), "tcp", BROKER_ADDRESS, PRODUCER_TOPIC, DEFAULT_PARTITION)
	if err != nil {
		log.Fatal("Falha ao se conectar com o Leader:", err)
	}

	return &Producer{
		Conn: conn,
	}
}

func (c *Producer) Close() {
	if err := c.Conn.Close(); err != nil {
		log.Fatal("Falha ao fechar a conexão:", err)
	}
}

func (c *Producer) WriteMessage(message []byte) (*Producer, error) {
	err := c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return c, err
	}

	_, err = c.Conn.WriteMessages(
		kafka.Message{
			Topic: PRODUCER_TOPIC,
			Value: message,
		},
	)

	log.Println("Mensagem postada!")
	return c, err
}
