package internal

import (
	"reflect"
	"time"

	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
)

type Status interface {
	Value() StatusValue
	String() string
	UpdatedAt() time.Time
	Equals(other Status) bool
}

type statusImpl struct {
	value     StatusValue
	updatedAt time.Time
}

func ReconstructStatus(value StatusValue, updatedAt time.Time) Status {
	return &statusImpl{StatusValue(value), updatedAt}
}

func NewInitialStatus(value *StatusValue) (Status, error) {
	return NewStatus(lo.FromPtrOr(value, StatusTodo))
}

func NewStatus(value StatusValue) (Status, error) {
	v := StatusValue(value)
	if v.Validate() {
		return nil, interfaces.NewInvalidArgError("status", "不正なステータスです")
	}
	return &statusImpl{v, time.Now()}, nil
}

func (s statusImpl) String() string { return string(s.value) }

func (s statusImpl) Value() StatusValue { return s.value }

func (s statusImpl) UpdatedAt() time.Time { return s.updatedAt }

func (s statusImpl) Equals(other Status) bool {
	return reflect.DeepEqual(s.Value(), other.Value())
}
