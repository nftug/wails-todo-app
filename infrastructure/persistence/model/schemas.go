package model

import "github.com/nftug/wails-todo-app/interfaces"

type RepositorySchema[TEntityPtr interfaces.Entity[TEntityPtr], TSelf any] interface {
	ToEntity() TEntityPtr
	Transfer(e TEntityPtr) TSelf
	GetPK() int
}

type QueryServiceSchema[TDetail any, TItem any] interface {
	ToDetailResponse() TDetail
	ToItemResponse() TItem
}
