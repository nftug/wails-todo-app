package model

import "github.com/nftug/wails-todo-app/domain/shared/entity"

type EntityTable[TEntity entity.Entity[TEntity]] interface {
	ToEntity() TEntity
	Transfer(e TEntity)
	GetPK() int
}

type ResponseTable[TDetail any, TItem any] interface {
	ToDetailResponse() TDetail
	ToItemResponse() TItem
}
