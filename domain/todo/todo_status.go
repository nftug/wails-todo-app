package todo

import (
	"reflect"
	"time"

	"github.com/nftug/wails-todo-app/interfaces"
	"github.com/samber/lo"
)

type Status interface {
	Value() StatusItem
	String() string
	UpdatedAt() time.Time
	Equals(other Status) bool
}

type statusImpl struct {
	value     StatusItem
	updatedAt time.Time
}

func ReconstructStatus(value StatusItem, updatedAt time.Time) Status {
	return &statusImpl{value, updatedAt}
}

func NewInitialStatus(value *StatusItem) (Status, error) {
	return NewStatus(lo.FromPtrOr(value, StatusTodo))
}

func NewStatus(value StatusItem) (Status, error) {
	if !lo.Contains(StatusSeq, value) {
		return nil, interfaces.NewInvalidArgError("status", "不正なステータスです")
	}
	return &statusImpl{value, time.Now()}, nil
}

func (s statusImpl) String() string { return string(s.value) }

func (s statusImpl) Value() StatusItem { return s.value }

func (s statusImpl) UpdatedAt() time.Time { return s.updatedAt }

func (s statusImpl) Equals(other Status) bool {
	return reflect.DeepEqual(s.Value(), other.Value())
}

// ステータスのEnum
type StatusItem string

const (
	StatusBacklog = StatusItem("Backlog")
	StatusTodo    = StatusItem("Todo")
	StatusDoing   = StatusItem("Doing")
	StatusDone    = StatusItem("Done")
)

var StatusSeq = []StatusItem{StatusBacklog, StatusTodo, StatusDoing, StatusDone}

func (s StatusItem) TSName() string { return string(s) }
