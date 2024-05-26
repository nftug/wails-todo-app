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
	ChangeStatus(value enum.StatusValue) (Status, error)
}

type statusImpl struct {
	value     enum.StatusValue
	updatedAt time.Time
}

func ReconstructStatus(value enum.StatusValue, updatedAt time.Time) Status {
	return &statusImpl{enum.StatusValue(value), updatedAt}
}

func NewStatus(value *enum.StatusValue) (Status, error) {
	return newStatus(lo.FromPtrOr(value, enum.StatusTodo))
}

func (s statusImpl) ChangeStatus(value enum.StatusValue) (Status, error) {
	if s.value == value {
		return nil, interfaces.NewInvalidArgError("status", "現在と異なるステータスを設定してください")
	}
	return newStatus(value)
}

func newStatus(value enum.StatusValue) (Status, error) {
	if !value.Validate() {
		return nil, interfaces.NewInvalidArgError("status", "不正なステータスです")
	}
	return &statusImpl{value, flextime.Now().UTC()}, nil
}

func (s statusImpl) String() string { return string(s.value) }

func (s statusImpl) Value() enum.StatusValue { return s.value }

func (s statusImpl) UpdatedAt() time.Time { return s.updatedAt }

func (s statusImpl) Equals(other Status) bool { return reflect.DeepEqual(s.Value(), other.Value()) }
