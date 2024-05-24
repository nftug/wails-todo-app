package todo

import (
	"reflect"
	"time"

	"github.com/nftug/wails-todo-app/domain/shared/myerror"
	"github.com/samber/lo"
)

type DueDate interface {
	Value() *time.Time
	String() string
	Equals(other DueDate) bool
}

type dueDateImpl struct {
	value time.Time
}

func ReconstructDueDate(value *time.Time) DueDate {
	return &dueDateImpl{lo.FromPtr(value)}
}

func NewDueDate(value *time.Time) (DueDate, error) {
	if value == nil {
		return &dueDateImpl{}, nil
	}

	v := lo.FromPtr(value)
	if v.Unix() < time.Now().Unix() {
		return nil, myerror.NewInvalidArgError("dueDate", "過去の日付は指定できません。")
	}

	return &dueDateImpl{v}, nil
}

func (d *dueDateImpl) Value() *time.Time {
	return lo.Ternary(lo.IsEmpty(d.value), nil, lo.ToPtr(d.value))
}

func (d *dueDateImpl) String() string {
	return d.value.String()
}

func (d *dueDateImpl) Equals(other DueDate) bool {
	return reflect.DeepEqual(d.Value(), other.Value())
}
