package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
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

	ch, err := source.Channel()
	if err != nil {
		log.Fatalf("creating channel failed: %s", err)
	}
	defer ch.Close()
	fmt.Println("channel created successful")

	gamelogic.PrintServerHelp()

	for {
		words := gamelogic.GetInput()
		if len(words) == 0 {
			continue
		}

		switch words[0] {
		case "pause":
			fmt.Println("Sending pause message...")
			err = pubsub.PublishJSON(
				ch,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{
					IsPaused: true,
				},
			)
			if err != nil {
				log.Fatalf("could not publish pause: %v", err)
			}
			fmt.Println("Pause message sent!")
		case "resume":
			fmt.Println("Sending resume message...")
			err = pubsub.PublishJSON(
				ch,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{
					IsPaused: false,
				},
			)
			if err != nil {
				log.Fatalf("could not publish resume: %v", err)
			}
			fmt.Println("Resume message sent!")
		case "help":
			gamelogic.PrintServerHelp()
		case "quit":
			fmt.Println("Exiting peril server...")
			return
		default:
			fmt.Printf("Unrecognized command: %v\n", words[0])
		}
	}
}
