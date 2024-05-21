package todo

import (
	"reflect"
	"time"

	"github.com/nftug/wails-todo-app/domain/myerror"
	"github.com/samber/lo"
)

type Status struct {
	value     StatusItem
	updatedAt time.Time
}

func ReconstructStatus(value StatusItem, updatedAt time.Time) Status {
	return Status{value, updatedAt}
}

func NewStatus(value StatusItem) (*Status, error) {
	if !lo.Contains(StatusSeq, value) {
		return nil, myerror.NewInvalidArgError("status", "不正なステータスです")
	}
	return &Status{value, time.Now()}, nil
}

func (s Status) String() string { return string(s.value) }

func (s *Status) Value() StatusItem { return s.value }

func (s *Status) UpdatedAt() time.Time { return s.updatedAt }

func (s *Status) Equals(other *Status) bool {
	return reflect.DeepEqual(s.value, other.value)
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
