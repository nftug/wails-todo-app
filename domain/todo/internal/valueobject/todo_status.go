package valueobject

import (
	"reflect"
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/nftug/wails-todo-app/interfaces/enums"
	"github.com/samber/lo"
)

type Status interface {
	Value() enums.StatusValue
	String() string
	UpdatedAt() time.Time
	Equals(other Status) bool
	ChangeStatus(value enums.StatusValue) (Status, error)
}

type statusImpl struct {
	value     enums.StatusValue
	updatedAt time.Time
}

func ReconstructStatus(value enums.StatusValue, updatedAt time.Time) Status {
	return &statusImpl{value, updatedAt}
}

func NewStatus(value *enums.StatusValue) (Status, error) {
	return newStatus(lo.FromPtrOr(value, enums.StatusTodo))
}

func (s statusImpl) ChangeStatus(value enums.StatusValue) (Status, error) {
	if s.value == value {
		return nil, interfaces.NewInvalidArgError("status", "現在と異なるステータスを設定してください")
	}
	return newStatus(value)
}

func newStatus(value enums.StatusValue) (Status, error) {
	if !value.Validate() {
		return nil, interfaces.NewInvalidArgError("status", "不正なステータスです")
	}
	return &statusImpl{value, flextime.Now().UTC()}, nil
}

func (s statusImpl) String() string { return string(s.value) }

func (s statusImpl) Value() enums.StatusValue { return s.value }

func (s statusImpl) UpdatedAt() time.Time { return s.updatedAt }

func (s statusImpl) Equals(other Status) bool { return reflect.DeepEqual(s.Value(), other.Value()) }
