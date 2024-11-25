package nullable

import (
	"reflect"

	"github.com/samber/lo"
)

// nilになりうる値を保持する値オブジェクト
type Nullable[T comparable] struct {
	value T
}

// nilになりうる値を保持する値オブジェクトを生成する。
func NewByPtr[T comparable](val *T) Nullable[T] {
	return Nullable[T]{lo.FromPtr(val)}
}

// nilになりうる値を保持する値オブジェクトを生成する。(値渡し)
func NewByVal[T comparable](val T) Nullable[T] {
	return Nullable[T]{val}
}

// 空の値が入った値オブジェクトを生成する。
func NewEmpty[T comparable]() Nullable[T] {
	return Nullable[T]{}
}

// 値がゼロ値の場合nilを返す。それ以外は値のコピーのポインタを返す。
func (nv Nullable[T]) Value() *T {
	if lo.IsEmpty(nv.value) {
		return nil
	} else {
		return lo.ToPtr(nv.value)
	}
}

// 値の実体を返す。ポインタがnilの場合はデフォルト値を返す。
func (nv Nullable[T]) RealValue() T { return nv.value }

// 値が等しい場合はtrueを返す。
func (nv Nullable[T]) Equals(other Nullable[T]) bool {
	return reflect.DeepEqual(nv.value, other.value)
}

// 値がデフォルト値かどうかを判定する。
func (nv Nullable[T]) IsEmpty() bool {
	return lo.IsEmpty(nv.value)
}
