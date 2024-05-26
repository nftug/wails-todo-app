package valueobject

import (
	"reflect"
	"strings"

	"github.com/nftug/wails-todo-app/interfaces"
)

type Title interface {
	Value() string
	String() string
	Equals(other Title) bool
}

type titleImpl struct{ value string }

func ReconstructTitle(value string) Title { return &titleImpl{value} }

func NewTitle(value string) (Title, error) {
	const MaxLength = 150

	value = strings.TrimSpace(value)
	if value == "" {
		return nil, interfaces.NewInvalidArgError("title", "タイトルを設定してください")
	}
	if len(value) > MaxLength {
		return nil, interfaces.NewInvalidArgError("title", "%d文字以内で入力してください", MaxLength)
	}

	return &titleImpl{value}, nil
}

func (t titleImpl) Value() string { return t.value }

func (t titleImpl) String() string { return t.value }

func (t titleImpl) Equals(other Title) bool { return reflect.DeepEqual(t.Value(), other.Value()) }
