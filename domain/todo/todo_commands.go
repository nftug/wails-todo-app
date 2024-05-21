package todo

type CreateCommand struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      StatusItem `json:"status"`
}

type EditCommand struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

type UpdateStatusCommand struct {
	Status StatusItem `json:"status"`
}
