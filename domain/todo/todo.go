package todo

import (
	"reflect"
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/domain/todo/internal/valueobject"
	"github.com/nftug/wails-todo-app/library/nullable"
	"github.com/nftug/wails-todo-app/shared/enums"
)

type Todo struct {
	id          int
	title       valueobject.Title
	description valueobject.Description
	status      valueobject.Status
	dueDate     valueobject.DueDate
	notifiedAt  nullable.Nullable[time.Time]
	createdAt   time.Time
	updatedAt   nullable.Nullable[time.Time]
}

// Getter

func (t *Todo) ID() int                    { return t.id }
func (t *Todo) Title() string              { return t.title.Value() }
func (t *Todo) Description() *string       { return t.description.Value() }
func (t *Todo) Status() enums.StatusValue  { return t.status.Value() }
func (t *Todo) StatusUpdatedAt() time.Time { return t.status.UpdatedAt() }
func (t *Todo) DueDate() *time.Time        { return t.dueDate.Value() }
func (t *Todo) NotifiedAt() *time.Time     { return t.notifiedAt.Value() }
func (t *Todo) CreatedAt() time.Time       { return t.createdAt }
func (t *Todo) UpdatedAt() *time.Time      { return t.updatedAt.Value() }

func Reconstruct(
	id int,
	title string,
	description *string,
	status enums.StatusValue,
	statusUpdatedAt time.Time,
	dueDate *time.Time,
	notifiedAt *time.Time,
	createdAt time.Time,
	updatedAt *time.Time) *Todo {
	return &Todo{
		id:          id,
		title:       valueobject.ReconstructTitle(title),
		description: valueobject.ReconstructDescription(description),
		status:      valueobject.ReconstructStatus(status, statusUpdatedAt),
		dueDate:     valueobject.ReconstructDueDate(dueDate),
		notifiedAt:  nullable.NewByPtr(notifiedAt),
		createdAt:   createdAt,
		updatedAt:   nullable.NewByPtr(updatedAt),
	}
}

func NewTodo(command CreateCommand) (*Todo, error) {
	title, err := valueobject.NewTitle(command.Title)
	if err != nil {
		return nil, err
	}
	desc, err := valueobject.NewDescription(command.Description)
	if err != nil {
		return nil, err
	}
	status, err := valueobject.NewStatus(command.InitialStatus)
	if err != nil {
		return nil, err
	}
	dueDate, err := valueobject.NewDueDate(command.DueDate)
	if err != nil {
		return nil, err
	}

	return &Todo{
		title:       title,
		description: desc,
		status:      status,
		dueDate:     dueDate,
		notifiedAt:  nullable.NewEmpty[time.Time](),
		createdAt:   flextime.Now().UTC(),
		updatedAt:   nullable.NewEmpty[time.Time](),
	}, nil
}

func (t *Todo) Update(command UpdateCommand) error {
	title, err := valueobject.NewTitle(command.Title)
	if err != nil {
		return err
	}
	desc, err := valueobject.NewDescription(command.Description)
	if err != nil {
		return err
	}
	dueDate, err := valueobject.NewDueDate(command.DueDate)
	if err != nil {
		return err
	}

	// Due dateが以前と異なる場合は、未通知状態に戻す
	if !dueDate.Equals(t.dueDate) {
		t.notifiedAt = nullable.NewEmpty[time.Time]()
	}

	t.dueDate = dueDate
	t.title = title
	t.description = desc
	t.updatedAt = nullable.NewByVal(flextime.Now().UTC())
	return nil
}

func (t *Todo) UpdateStatus(command UpdateStatusCommand) error {
	status, err := t.status.ChangeStatus(command.Status)
	if err != nil {
		return err
	}

	t.status = status
	return nil
}

func (t *Todo) SetNotifiedAt(notifiedAt time.Time) {
	t.notifiedAt = nullable.NewByVal(notifiedAt.UTC())
}

func (t *Todo) ClearNotifiedAt() {
	t.notifiedAt = nullable.NewEmpty[time.Time]()
}

func (t *Todo) SetID(id int) { t.id = id }

func (t *Todo) Equals(other *Todo) bool {
	return reflect.DeepEqual(t.id, other.id)
}
