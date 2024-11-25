package dtos

import "github.com/nftug/wails-todo-app/shared/interfaces"

type CreatedResponse struct {
	ID string `json:"id"`
}

func NewCreatedResponse[TEntityPtr interfaces.Entity[TEntityPtr]](e TEntityPtr) *CreatedResponse {
	return &CreatedResponse{e.ID().String()}
}
