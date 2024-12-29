package interfaces

import (
	"time"
)

type Equaler[TSelf any] interface {
	Equals(TSelf) bool
}

type Entity[TSelfPtr any] interface {
	Equaler[TSelfPtr]
	ID() int
	CreatedAt() time.Time
	UpdatedAt() *time.Time
	SetID(id int)
}
