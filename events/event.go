package events

import (
	"context"

	"github.com/daluisgarcia/golang-events-cqrs/models"
)

type EventStore interface {
	Close()
	PublishCreatedFeed(ctx context.Context, feed *models.Feed) error
	SuscribeCreatedFeed(ctx context.Context) (<-chan *models.Feed, error)
	OnCreateFeed(f func(CreatedFeedMessage)) error // Callback that reacts when a new feed has been created
}

var eventStore EventStore

func SetEventStore(eStore EventStore) {
	eventStore = eStore
}

func Close() {
	eventStore.Close()
}

func PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishCreatedFeed(ctx, feed)
}

func SuscribeCreatedFeed(ctx context.Context) (<-chan *models.Feed, error) {
	return eventStore.SuscribeCreatedFeed(ctx)
}

func OnCreateFeed(f func(CreatedFeedMessage)) error {
	return eventStore.OnCreateFeed(f)
}
