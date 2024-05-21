package todo

import (
	"reflect"
	"strings"

	"github.com/nftug/wails-todo-app/domain/myerror"
)

type Description struct {
	value *string
}

func ReconstructDescription(value *string) Description {
	return Description{value}
}

func NewDescription(value *string) (*Description, error) {
	const MaxLength = 1000

	if value == nil {
		return &Description{}, nil
	}
	trimmed := strings.TrimSpace(*value)
	if len(trimmed) > MaxLength {
		return nil, myerror.NewInvalidArgError("description", "%d文字以内で入力してください", MaxLength)
	}

	return &Description{value}, nil
}

func (t *Description) Value() *string { return t.value }

func (t *Description) String() string {
	if t.value == nil {
		return ""
	} else {
		return (*t.value)
	}
}

func (t *Description) Equals(other *Description) bool {
	return reflect.DeepEqual(t.value, other.value)
}
