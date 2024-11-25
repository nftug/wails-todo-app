package valueobject

import (
	"strings"
	"unicode/utf8"

	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/nftug/wails-todo-app/library/nullable"
	"github.com/samber/lo"
)

type Description interface {
	Value() *string
	String() string
	Equals(other Description) bool
}

type descriptionImpl struct{ nullable.Nullable[string] }

func ReconstructDescription(value *string) Description {
	return descriptionImpl{nullable.NewByPtr(value)}
}

func NewDescription(value *string) (Description, error) {
	const MaxLength = 200

	trimmed := strings.TrimSpace(lo.FromPtr(value))
	if utf8.RuneCountInString(trimmed) > MaxLength {
		return nil, interfaces.NewInvalidArgError("description", "%d文字以内で入力してください", MaxLength)
	}

	return descriptionImpl{nullable.NewByVal(trimmed)}, nil
}

func (t descriptionImpl) String() string {
	return t.RealValue()
}

func (t descriptionImpl) Equals(other Description) bool {
	return t.Nullable.Equals(other.(descriptionImpl).Nullable)
}
