package todo

import (
	"context"
	"time"
)

type TodoRepository interface {
	Find(ctx context.Context, id int) (*Todo, error)
	Save(ctx context.Context, entity *Todo) error
	Delete(ctx context.Context, id int) error
	FindAllForNotification(ctx context.Context, dueDate time.Time) ([]*Todo, error)
}

type TodoQueryService interface {
	Find(ctx context.Context, id int) (*DetailsResponse, error)
	FindAll(ctx context.Context, q Query) ([]*ItemResponse, error)
}

type TodoNotificationSender interface {
	Send(ctx context.Context, item *Todo) error
}
