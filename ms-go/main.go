package main

import (
	KafkaConsumer "ms-go/app/consumers"
	_ "ms-go/db"
	"ms-go/router"
)

func main() {

	consumer := KafkaConsumer.NewConsumer([]string{"kafka:29092"}, "iCasei-ms-go", "rails-to-go")
	defer consumer.Close()

	consumer.ConsumeMessages()
	router.Run()
}
