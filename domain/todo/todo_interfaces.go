package todo

import (
	"context"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Find(id uuid.UUID, ctx context.Context) (*Todo, error)
	Save(entity *Todo, ctx context.Context) error
	Delete(entity *Todo, ctx context.Context) error
}

type TodoQueryService interface {
	Find(id uuid.UUID, ctx context.Context) (*DetailResponse, error)
	FindAll(q Query, ctx context.Context) ([]*ItemResponse, error)
}
