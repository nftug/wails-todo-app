package valueobject

import (
	"strings"
	"unicode/utf8"

	"github.com/nftug/wails-todo-app/library/nullable"
	"github.com/nftug/wails-todo-app/shared/customerr"
	"github.com/samber/lo"
)

type Description interface {
	Value() *string
	String() string
	Equals(other Description) bool
}

type descriptionImpl struct{ nullable.Nullable[string] }

func ReconstructDescription(value *string) Description {
	return descriptionImpl{nullable.New(value)}
}

func NewDescription(value *string) (Description, error) {
	const MaxLength = 200

	trimmed := strings.TrimSpace(lo.FromPtr(value))
	if utf8.RuneCountInString(trimmed) > MaxLength {
		return nil, customerr.NewValidationError("description", "%d文字以内で入力してください", MaxLength)
	}

	return descriptionImpl{nullable.NewByVal(trimmed)}, nil
}

func (t descriptionImpl) Value() *string {
	return t.ToCopiedPtr()
}

func (t descriptionImpl) String() string {
	return t.Nullable.Value
}

func (t descriptionImpl) Equals(other Description) bool {
	return t.Nullable.Value == other.(descriptionImpl).Nullable.Value
}
