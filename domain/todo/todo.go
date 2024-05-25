package todo

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/nftug/wails-todo-app/domain/todo/internal"
	"github.com/nftug/wails-todo-app/library/nullable"
)

type Todo struct {
	pk          int
	id          uuid.UUID
	title       internal.Title
	description internal.Description
	status      internal.Status
	dueDate     internal.DueDate
	createdAt   time.Time
	updatedAt   nullable.Nullable[time.Time]
}

// Getter

func (t Todo) PK() int                    { return t.pk }
func (t Todo) ID() uuid.UUID              { return t.id }
func (t Todo) Title() string              { return t.title.Value() }
func (t Todo) Description() *string       { return t.description.Value() }
func (t Todo) Status() StatusValue        { return StatusValue(t.status.Value()) }
func (t Todo) StatusUpdatedAt() time.Time { return t.status.UpdatedAt() }
func (t Todo) DueDate() *time.Time        { return t.dueDate.Value() }
func (t Todo) CreatedAt() time.Time       { return t.createdAt }
func (t Todo) UpdatedAt() *time.Time      { return t.updatedAt.Value() }

func Reconstruct(
	pk int,
	id uuid.UUID,
	title string,
	description *string,
	status StatusValue,
	statusUpdatedAt time.Time,
	dueDate *time.Time,
	createdAt time.Time,
	updatedAt *time.Time) *Todo {
	return &Todo{
		pk:          pk,
		id:          id,
		title:       internal.ReconstructTitle(title),
		description: internal.ReconstructDescription(description),
		status:      internal.ReconstructStatus(status.toInternal(), statusUpdatedAt),
		dueDate:     internal.ReconstructDueDate(dueDate),
		createdAt:   createdAt,
		updatedAt:   nullable.NewByPtr(updatedAt),
	}
}

func NewTodo(command CreateCommand) (*Todo, error) {
	title, err := internal.NewTitle(command.Title)
	if err != nil {
		return nil, err
	}
	desc, err := internal.NewDescription(command.Description)
	if err != nil {
		return nil, err
	}
	status, err := internal.NewInitialStatus(command.InitialStatus.toInternalPtr())
	if err != nil {
		return nil, err
	}
	dueDate, err := internal.NewDueDate(command.DueDate)
	if err != nil {
		return nil, err
	}

	return &Todo{
		id:          uuid.New(),
		title:       title,
		description: desc,
		status:      status,
		dueDate:     dueDate,
		createdAt:   time.Now(),
	}, nil
}

func (t *Todo) Update(command UpdateCommand) error {
	title, err := internal.NewTitle(command.Title)
	if err != nil {
		return err
	}
	desc, err := internal.NewDescription(command.Description)
	if err != nil {
		return err
	}
	dueDate, err := internal.NewDueDate(command.DueDate)
	if err != nil {
		return err
	}

	t.title = title
	t.description = desc
	t.dueDate = dueDate
	t.updatedAt = nullable.NewByVal(time.Now())
	return nil
}

func (t *Todo) UpdateStatus(command UpdateStatusCommand) error {
	status, err := internal.NewStatus(command.Status.toInternal())
	if err != nil {
		return err
	}

	t.status = status
	return nil
}

func (t *Todo) SetPK(pk int) { t.pk = pk }

func (t Todo) Equals(other *Todo) bool {
	return reflect.DeepEqual(t.id, other.id)
}
