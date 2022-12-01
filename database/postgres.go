package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/daluisgarcia/golang-events-cqrs/models"
)

type PostgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(urlDB string) (*PostgreRepository, error) {
	db, err := sql.Open("postgres", urlDB)

	if err != nil {
		return nil, err
	}

	return &PostgreRepository{db: db}, nil
}

func (r *PostgreRepository) Close() {
	r.db.Close()
}

func (r *PostgreRepository) InsertFeed(ctx context.Context, feed *models.Feed) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO feeds (id, title, description) VALUES ($1, $2, $3)", feed.ID, feed.Title, feed.Description)
	return err
}

func (r *PostgreRepository) ListFeeds(ctx context.Context) ([]*models.Feed, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, created_at FROM feeds")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feeds []*models.Feed

	for rows.Next() {
		feed := new(models.Feed)
		if err := rows.Scan(&feed.ID, &feed.Title, &feed.Description, &feed.CreatedAt); err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
