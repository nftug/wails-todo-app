package todo

import "time"

type Query struct {
	Search      *string     `json:"search"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Status      *StatusItem `json:"status"`
	Until       *time.Time  `json:"until"`
}
