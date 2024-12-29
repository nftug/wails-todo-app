package util

import (
	"github.com/samber/lo"
)

func EmptyableToCopiedPtr[T any](x T) *T {
	p := lo.EmptyableToPtr(x)
	if p == nil {
		return nil
	}
	return lo.ToPtr(*p)
}
