package search

import (
	"context"

	"github.com/daluisgarcia/golang-events-cqrs/models"
)

type SearchRepository interface {
	Close()
	IndexFeed(ctx context.Context, feed *models.Feed) error
	SeatchFeed(ctx context.Context, query string) ([]*models.Feed, error)
}

var repo SearchRepository

func SetSearcRepository(r SearchRepository) {
	repo = r
}

func Close() {
	repo.Close()
}

func IndexFeed(ctx context.Context, feed *models.Feed) error {
	return repo.IndexFeed(ctx, feed)
}

func SearchFeed(ctx context.Context, query string) ([]*models.Feed, error) {
	return repo.SeatchFeed(ctx, query)
}
