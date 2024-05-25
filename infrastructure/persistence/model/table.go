package model

import "github.com/nftug/wails-todo-app/interfaces"

type EntityTable[TEntity interfaces.Entity[TEntity]] interface {
	ToEntity() TEntity
	Transfer(e TEntity)
	GetPK() int
}

type ResponseTable[TDetail any, TItem any] interface {
	ToDetailResponse() TDetail
	ToItemResponse() TItem
}
