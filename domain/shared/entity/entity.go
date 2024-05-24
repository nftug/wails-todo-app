package entity

import (
	"time"

	"github.com/google/uuid"
)

type Equaler[T any] interface {
	Equals(T) bool
}

type Entity[T any] interface {
	Equaler[T]
	PK() int
	ID() uuid.UUID
	CreatedAt() time.Time
	UpdatedAt() *time.Time
	SetPK(pk int)
}
