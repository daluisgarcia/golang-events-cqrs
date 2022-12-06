package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/daluisgarcia/golang-events-cqrs/events"
	"github.com/daluisgarcia/golang-events-cqrs/models"
	"github.com/daluisgarcia/golang-events-cqrs/repository"
	"github.com/daluisgarcia/golang-events-cqrs/search"
)

func onCreatedFeed(msg events.CreatedFeedMessage) {
	feed := models.Feed{
		ID:          msg.ID,
		Title:       msg.Title,
		Description: msg.Description,
		CreatedAt:   msg.CreatedAt,
	}

	if err := search.IndexFeed(context.Background(), feed); err != nil {
		log.Printf("Error indexing feed: %v", err)
	}
}

func listFeedsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	feeds, err := repository.ListFeeds(ctx)
	if err != nil {
		http.Error(w, "Error listing feeds", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Getting query param "q"
	query := r.URL.Query().Get("q")

	if len(query) == 0 {
		http.Error(w, "Missing query param", http.StatusBadRequest)
		return
	}

	feeds, err := search.SearchFeed(ctx, query)
	if err != nil {
		http.Error(w, "Error searching feeds", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(feeds)
}
