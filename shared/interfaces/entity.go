package interfaces

import (
	"time"

	"github.com/google/uuid"
)

type Equaler[TSelf any] interface {
	Equals(TSelf) bool
}

type Entity[TSelfPtr any] interface {
	Equaler[TSelfPtr]
	PK() int
	ID() uuid.UUID
	CreatedAt() time.Time
	UpdatedAt() *time.Time
	SetPK(pk int)
}
