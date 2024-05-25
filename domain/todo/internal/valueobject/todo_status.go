package valueobject

import (
	"reflect"
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/domain/todo/internal/enum"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
)

type Status interface {
	Value() enum.StatusValue
	String() string
	UpdatedAt() time.Time
	Equals(other Status) bool
}

type statusImpl struct {
	value     enum.StatusValue
	updatedAt time.Time
}

func ReconstructStatus(value enum.StatusValue, updatedAt time.Time) Status {
	return &statusImpl{enum.StatusValue(value), updatedAt}
}

func NewInitialStatus(value *enum.StatusValue) (Status, error) {
	return NewStatus(lo.FromPtrOr(value, enum.StatusTodo))
}

func NewStatus(value enum.StatusValue) (Status, error) {
	v := enum.StatusValue(value)
	if !v.Validate() {
		return nil, interfaces.NewInvalidArgError("status", "不正なステータスです")
	}
	return &statusImpl{v, flextime.Now().UTC()}, nil
}

func (s statusImpl) String() string { return string(s.value) }

func (s statusImpl) Value() enum.StatusValue { return s.value }

func (s statusImpl) UpdatedAt() time.Time { return s.updatedAt }

func (s statusImpl) Equals(other Status) bool {
	return reflect.DeepEqual(s.Value(), other.Value())
}
