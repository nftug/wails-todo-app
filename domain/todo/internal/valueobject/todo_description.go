package valueobject

import (
	"strings"

	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/nftug/wails-todo-app/library/nullable"
)

type Description interface {
	Value() *string
	String() string
	Equals(other Description) bool
}

type descriptionImpl struct {
	nullable.Nullable[string]
}

func ReconstructDescription(value *string) Description {
	return &descriptionImpl{nullable.NewByPtr(value)}
}

func NewDescription(value *string) (Description, error) {
	const MaxLength = 1000
	v := nullable.NewByPtr(value)

	if v.IsEmpty() {
		return &descriptionImpl{v}, nil
	}

	trimmed := strings.TrimSpace(v.RawValue())
	if len(trimmed) > MaxLength {
		return nil, interfaces.NewInvalidArgError("description", "%d文字以内で入力してください", MaxLength)
	}

	return &descriptionImpl{nullable.NewByVal(trimmed)}, nil
}

func (t descriptionImpl) String() string { return t.RawValue() }

func (t descriptionImpl) Equals(other Description) bool {
	return t.EqualsByVal(*other.Value())
}
