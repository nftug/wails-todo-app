package interfaces

import (
	"github.com/google/uuid"
)

type CreatedResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewCreatedResponse[TEntity Entity[TEntity]](e TEntity) *CreatedResponse {
	return &CreatedResponse{e.ID()}
}
