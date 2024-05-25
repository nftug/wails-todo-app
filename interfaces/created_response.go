package interfaces

import (
	"github.com/google/uuid"
)

type CreatedResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewCreatedResponse[TEntityPtr Entity[TEntityPtr]](e TEntityPtr) *CreatedResponse {
	return &CreatedResponse{e.ID()}
}
