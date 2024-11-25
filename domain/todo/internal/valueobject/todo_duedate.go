package valueobject

import (
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/library/nullable"
	"github.com/nftug/wails-todo-app/shared/customerr"
	"github.com/samber/lo"
)

type DueDate interface {
	Value() *time.Time
	String() string
	Equals(other DueDate) bool
}

type dueDateImpl struct{ nullable.Nullable[time.Time] }

func ReconstructDueDate(value *time.Time) DueDate {
	return dueDateImpl{nullable.NewByPtr(value)}
}

func NewDueDate(value *time.Time) (DueDate, error) {
	if value == nil {
		return dueDateImpl{nullable.NewEmpty[time.Time]()}, nil
	}

	v := nullable.NewByVal(lo.FromPtr(value).UTC())
	if v.Value().Unix() < flextime.Now().UTC().Unix() {
		return nil, customerr.NewValidationError("dueDate", "過去の日付は指定できません。")
	}
	return dueDateImpl{v}, nil
}

func (d dueDateImpl) String() string {
	return d.RealValue().String()
}

func (d dueDateImpl) Equals(other DueDate) bool {
	return d.Nullable.Equals(other.(dueDateImpl).Nullable)
}
