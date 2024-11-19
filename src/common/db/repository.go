package db

import (
	"context"
	"github.com/SemyonTolkachyov/message-board/src/common/schema"
)

type Repository interface {
	Close()
	InsertMessage(ctx context.Context, message schema.Message) error
	ListMessages(ctx context.Context, skip uint64, take uint64) ([]schema.Message, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertMessage(ctx context.Context, message schema.Message) error {
	return impl.InsertMessage(ctx, message)
}

func ListMessages(ctx context.Context, skip uint64, take uint64) ([]schema.Message, error) {
	return impl.ListMessages(ctx, skip, take)
}
