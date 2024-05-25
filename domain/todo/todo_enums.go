package todo

import (
	"github.com/nftug/wails-todo-app/domain/todo/internal/enum"
	"github.com/samber/lo"
)

// ステータスのEnum (公開用)
type StatusValue enum.StatusValue

// ステータスのシーケンス (Wailsへの公開用)
var StatusSeq = lo.Map(enum.StatusSeq,
	func(x enum.StatusValue, _ int) StatusValue { return StatusValue(x) })

func (s StatusValue) TSName() string { return string(s) }

// 内部用のEnumへ変換
func (s StatusValue) toInternal() enum.StatusValue {
	return enum.StatusValue(s)
}

// 内部用のEnumのポインタへ変換
func (s *StatusValue) toInternalPtr() *enum.StatusValue {
	val := enum.StatusValue(lo.FromPtr(s))
	return &val
}
