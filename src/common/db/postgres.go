package db

import (
	"context"
	"database/sql"
	"github.com/SemyonTolkachyov/message-board/src/common/schema"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	if err := r.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (r *PostgresRepository) InsertMessage(ctx context.Context, message schema.Message) error {
	_, err := r.db.Exec("INSERT INTO messages(id, body, created_at) VALUES($1, $2, $3)", message.Id, message.Body, message.CreatedAt)
	return err
}

func (r *PostgresRepository) ListMessages(ctx context.Context, skip uint64, take uint64) ([]schema.Message, error) {
	rows, err := r.db.Query("SELECT id, body, created_at FROM messages ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Parse all rows into an array of Messages
	var messages []schema.Message
	for rows.Next() {
		message := schema.Message{}
		if err = rows.Scan(&message.Id, &message.Body, &message.CreatedAt); err == nil {
			messages = append(messages, message)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
