package interfaces

type CreatedResponse struct {
	ID string `json:"id"`
}

func NewCreatedResponse[TEntityPtr Entity[TEntityPtr]](e TEntityPtr) *CreatedResponse {
	return &CreatedResponse{e.ID().String()}
}
