package valueobject

import (
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/library/nullable"
	"github.com/nftug/wails-todo-app/shared/customerr"
)

type DueDate interface {
	Value() *time.Time
	String() string
	Equals(other DueDate) bool
}

type dueDateImpl struct{ nullable.Nullable[time.Time] }

func ReconstructDueDate(value *time.Time) DueDate {
	return dueDateImpl{nullable.New(value)}
}

func NewDueDate(value *time.Time) (DueDate, error) {
	if value == nil {
		return dueDateImpl{nullable.Nullable[time.Time]{}}, nil
	}

	v := value.UTC()
	if v.Before(flextime.Now().UTC()) {
		return nil, customerr.NewValidationError("dueDate", "過去の日付は指定できません。")
	}
	return dueDateImpl{nullable.NewByVal(v)}, nil
}

func (d dueDateImpl) Value() *time.Time {
	return d.ToCopiedPtr()
}

func (d dueDateImpl) String() string {
	return d.Nullable.Value.String()
}

func (d dueDateImpl) Equals(other DueDate) bool {
	return d.Nullable.Value == other.(dueDateImpl).Nullable.Value
}
