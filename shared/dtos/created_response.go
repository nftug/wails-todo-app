package dtos

import "github.com/nftug/wails-todo-app/shared/interfaces"

type CreatedResponse struct {
	ID int `json:"id"`
}

func NewCreatedResponse[TEntityPtr interfaces.Entity[TEntityPtr]](e TEntityPtr) *CreatedResponse {
	return &CreatedResponse{e.ID()}
}
