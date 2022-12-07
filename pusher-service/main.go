package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daluisgarcia/golang-events-cqrs/events"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatalf("%v", err)
	}

	hub := NewHub()

	natsAddr := fmt.Sprintf("nats://%s", cfg.NatsAddress)

	n, err := events.NewNatsEventStore(natsAddr)

	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to the events
	err = n.OnCreateFeed(func(msg events.CreatedFeedMessage) {
		hub.Broadcast(
			newCreatedFeedMessage(
				msg.ID,
				msg.Title,
				msg.Description,
				msg.CreatedAt,
			),
			nil,
		)
	})

	if err != nil {
		log.Fatal(err)
	}

	events.SetEventStore(n)
	defer events.Close()

	go hub.Run()

	http.HandleFunc("/ws", hub.HandleWebSocket)
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
