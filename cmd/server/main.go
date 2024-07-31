package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	url := "amqp://guest:guest@localhost:5672/"
	source, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("amqp connection failed: %s", err)
	}
	defer source.Close()
	fmt.Println("amqp connection successful")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
