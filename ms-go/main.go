package main

import (
	KafkaConsumer "ms-go/app/consumers"
	KafkaProducer "ms-go/app/producers"
	_ "ms-go/db"
	"ms-go/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	consumer := KafkaConsumer.NewConsumer([]string{"kafka:29092"}, "iCasei-ms-go", "rails-to-go")
	defer consumer.Close()

	go func() {
		consumer.ConsumeMessages()
	}()

	producer := KafkaProducer.GetProducer()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		router.Run()
	}()

	<-sigs

	producer.Close()
}
