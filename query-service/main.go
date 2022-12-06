package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daluisgarcia/golang-events-cqrs/database"
	"github.com/daluisgarcia/golang-events-cqrs/events"
	"github.com/daluisgarcia/golang-events-cqrs/repository"
	"github.com/daluisgarcia/golang-events-cqrs/search"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresDB            string `envconfig:"POSTGRES_DB"` // Allows to be read by .env
	PostgresUser          string `envconfig:"POSTGRES_USER"`
	PostgresPassword      string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress           string `envconfig:"NATS_ADDRESS"`
	ElasticesearchAddress string `envconfig:"ELASTICSEARCH_ADDRESS"`
}

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/feeds", listFeedsHandler).Methods("GET")
	router.HandleFunc("/search", searchHandler).Methods("GET")
	return
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatalf("%v", err)
	}

	// Connect to Postgres
	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	repo, err := database.NewPostgreRepository(addr)

	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)

	// Connect to ElasticSearch
	esAddr := fmt.Sprintf("http://%s", cfg.ElasticesearchAddress)

	esRepo, err := search.NewElasticSearchRepository(esAddr)

	if err != nil {
		log.Fatal(err)
	}

	search.SetSearchRepository(esRepo)
	defer search.Close()

	// Connect to NATS
	natsAddr := fmt.Sprintf("nats://%s", cfg.NatsAddress)

	n, err := events.NewNatsEventStore(natsAddr)
	if err != nil {
		log.Fatal(err)
	}

	err = n.OnCreateFeed(onCreatedFeed)

	if err != nil {
		log.Fatal(err)
	}

	events.SetEventStore(n)
	defer events.Close()

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
