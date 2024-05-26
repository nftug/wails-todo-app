package todo

import (
	"context"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*Todo, error)
	Save(ctx context.Context, entity *Todo) error
	Delete(ctx context.Context, entity *Todo) error
}

type TodoQueryService interface {
	Find(ctx context.Context, id uuid.UUID) (*DetailResponse, error)
	FindAll(ctx context.Context, q Query) ([]*ItemResponse, error)
}
