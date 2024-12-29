package nullable

import "github.com/samber/lo"

type Nullable[T any] struct {
	Value T
}

func New[T any](x *T) Nullable[T] {
	return Nullable[T]{lo.FromPtr(x)}
}

func NewByVal[T any](x T) Nullable[T] {
	return Nullable[T]{x}
}

func (n Nullable[T]) ToCopiedPtr() *T {
	p := lo.EmptyableToPtr(n.Value)
	if p == nil {
		return nil
	}
	return lo.ToPtr(*p)
}
