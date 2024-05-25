package todo

import (
	"reflect"
	"strings"

	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
)

type Description interface {
	Value() *string
	String() string
	Equals(other Description) bool
}

type descriptionImpl struct {
	value string
}

func ReconstructDescription(value *string) Description {
	return &descriptionImpl{lo.FromPtr(value)}
}

func NewDescription(value *string) (Description, error) {
	const MaxLength = 1000

	if value == nil {
		return &descriptionImpl{}, nil
	}

	trimmed := strings.TrimSpace(*value)
	if len(trimmed) > MaxLength {
		return nil, interfaces.NewInvalidArgError("description", "%d文字以内で入力してください", MaxLength)
	}

	return &descriptionImpl{trimmed}, nil
}

func (t descriptionImpl) Value() *string {
	return lo.Ternary(lo.IsEmpty(t.value), nil, lo.ToPtr(t.value))
}

func (t descriptionImpl) String() string { return t.value }

func (t descriptionImpl) Equals(other Description) bool {
	return reflect.DeepEqual(t.Value(), other.Value())
}
