package main

import (
	KafkaConsumer "ms-go/app/consumers"
	KafkaProducer "ms-go/app/producers"
	_ "ms-go/db"
	"ms-go/router"
	"time"
)

func main() {

	time.Sleep(time.Second * 10)

	producer := KafkaProducer.NewProducer()
	producer.Close()

	consumer := KafkaConsumer.NewConsumer([]string{KafkaConsumer.BROKER_ADDRESS}, KafkaConsumer.CONSUMER_TOPIC)
	defer consumer.Close()

	go consumer.ReadMessages()

	router.Run()
}
