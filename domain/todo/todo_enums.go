package todo

import (
	"github.com/nftug/wails-todo-app/domain/todo/internal"
	"github.com/samber/lo"
)

// ステータスのEnum (公開用)
type StatusValue internal.StatusValue

// ステータスのシーケンス (Wailsへの公開用)
var StatusSeq = lo.Map(internal.StatusSeq,
	func(x internal.StatusValue, _ int) StatusValue { return StatusValue(x) })

func (s StatusValue) TSName() string { return string(s) }

// 内部用のEnumへ変換
func (s StatusValue) toInternal() internal.StatusValue {
	return internal.StatusValue(s)
}

// 内部用のEnumのポインタへ変換
func (s *StatusValue) toInternalPtr() *internal.StatusValue {
	val := internal.StatusValue(lo.FromPtr(s))
	return &val
}
