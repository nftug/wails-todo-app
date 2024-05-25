package nullable

import (
	"reflect"

	"github.com/samber/lo"
)

// nilになりうる値を保持する値オブジェクト
type Nullable[T comparable] interface {
	Value() *T
	RawValue() T
	Equals(other Nullable[T]) bool
	EqualsByVal(other T) bool
	IsEmpty() bool
}

type nullableImpl[T comparable] struct {
	value T
}

// nilになりうる値を保持する値オブジェクトを生成する。
func NewByPtr[T comparable](val *T) Nullable[T] {
	return &nullableImpl[T]{lo.FromPtr(val)}
}

// nilになりうる値を保持する値オブジェクトを生成する。(デフォルト値付き)
func NewByPtrOr[T comparable](val *T, fallback T) Nullable[T] {
	return &nullableImpl[T]{lo.FromPtrOr(val, fallback)}
}

// nilになりうる値を保持する値オブジェクトを生成する。(値渡し)
func NewByVal[T comparable](val T) Nullable[T] {
	return &nullableImpl[T]{val}
}

// 空の値が入った値オブジェクトを生成する。
func NewEmpty[T comparable]() Nullable[T] { return &nullableImpl[T]{} }

// 値がゼロ値の場合nilを返す。それ以外は値のコピーのポインタを返す。
func (nv nullableImpl[T]) Value() *T {
	// 生のポインタを渡すと不変性が崩れるため、lo.ToPtr()で値のコピーのポインタを渡す。
	return lo.Ternary(lo.IsEmpty(nv.value), nil, lo.ToPtr(nv.value))
}

// 生の値を返す。オリジナルがnilの場合はゼロ値を返す。
func (nv nullableImpl[T]) RawValue() T { return nv.value }

// 値の内容を等価で判定する。
func (nv nullableImpl[T]) Equals(other Nullable[T]) bool {
	return reflect.DeepEqual(nv.Value(), other.Value())
}

// 値の内容を等価で判定する。
func (nv nullableImpl[T]) EqualsByVal(other T) bool {
	return nv.value == other
}

// 値がデフォルト値かどうかを判定する。
func (nv nullableImpl[T]) IsEmpty() bool {
	return lo.IsEmpty(nv)
}
