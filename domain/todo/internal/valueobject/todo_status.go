package valueobject

import (
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/shared/customerr"
	"github.com/nftug/wails-todo-app/shared/enums"
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
	return statusImpl{value, updatedAt}
}

func NewStatus(value *enums.StatusValue) (Status, error) {
	return newStatus(lo.FromPtrOr(value, enums.StatusTodo))
}

func (s statusImpl) ChangeStatus(value enums.StatusValue) (Status, error) {
	if s.value == value {
		return nil, customerr.NewValidationError("status", "現在と異なるステータスを設定してください")
	}
	return newStatus(value)
}

func newStatus(value enums.StatusValue) (Status, error) {
	if !value.Validate() {
		return nil, customerr.NewValidationError("status", "不正なステータスです")
	}
	return statusImpl{value, flextime.Now().UTC()}, nil
}

func (s statusImpl) String() string { return string(s.value) }

func (s statusImpl) Value() enums.StatusValue { return s.value }

func (s statusImpl) UpdatedAt() time.Time { return s.updatedAt }

func (s statusImpl) Equals(other Status) bool {
	return s.Value() == other.Value()
}
