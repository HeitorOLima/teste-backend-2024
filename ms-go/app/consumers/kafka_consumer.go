package KafkaConsumer

import (
	"context"
	"encoding/json"
	"log"
	"ms-go/app/models"
	"ms-go/app/services/products"

	"github.com/segmentio/kafka-go"
)

const (
	CONSUMER_TOPIC    = "rails-to-go"
	DEFAULT_PARTITION = 0
	BROKER_ADDRESS    = "localhost:9092"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *KafkaConsumer {
	config := kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MaxBytes: 10e6,
	}
	reader := kafka.NewReader(config)
	return &KafkaConsumer{
		Reader: reader,
	}
}

func (kc *KafkaConsumer) ReadMessages() {
	for {
		log.Printf("Buscando mensagens")
		m, err := kc.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			continue
		}

		var product models.Product
		err = json.Unmarshal(m.Value, &product)
		if err != nil {
			log.Printf("Erro ao decodificar mensagem JSON: %v", err)
			continue
		}

		existingProduct, _ := products.Details(product)
		if existingProduct != nil {
			_, err := products.Update(product, false)
			if err != nil {
				log.Printf("Erro ao atualizar o produto: %v", err)
				continue
			}
		} else {
			_, err := products.Create(product, false)
			if err != nil {
				log.Printf("Erro ao criar o produto: %v", err)
				continue
			}
		}
	}
}

func (kc *KafkaConsumer) Close() {
	kc.Reader.Close()
}
