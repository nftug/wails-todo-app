package todo

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*Todo, error)
	Save(ctx context.Context, entity *Todo) error
	Delete(ctx context.Context, entity *Todo) error
	FindAllForNotification(ctx context.Context, dueDate time.Time) ([]*Todo, error)
}

type TodoQueryService interface {
	Find(ctx context.Context, id uuid.UUID) (*DetailsResponse, error)
	FindAll(ctx context.Context, q Query) ([]*ItemResponse, error)
}

type TodoNotificationSender interface {
	Send(ctx context.Context, item *Todo) error
}
