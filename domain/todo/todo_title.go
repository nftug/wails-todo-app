package todo

import (
	"reflect"
	"strings"

	"github.com/nftug/wails-todo-app/domain/myerror"
)

type Title struct {
	value string
}

func ReconstructTitle(value string) Title {
	return Title{value}
}

func NewTitle(value string) (*Title, error) {
	const MaxLength = 150

	value = strings.TrimSpace(value)
	if value == "" {
		return nil, myerror.NewInvalidArgError("title", "タイトルを設定してください")
	}
	if len(value) > MaxLength {
		return nil, myerror.NewInvalidArgError("title", "%d文字以内で入力してください", MaxLength)
	}

	return &Title{value}, nil
}

func (t *Title) Value() string { return t.value }

func (t *Title) String() string { return t.value }

func (t *Title) Equals(other *Title) bool {
	return reflect.DeepEqual(t.value, other.value)
}
