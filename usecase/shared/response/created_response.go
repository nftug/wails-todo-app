package response

import (
	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/shared/entity"
)

type CreatedResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewCreatedResponse[TEntity entity.Entity[TEntity]](e TEntity) *CreatedResponse {
	return &CreatedResponse{e.ID()}
}
