package todo

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Todo struct {
	pk          int
	id          uuid.UUID
	title       Title
	description Description
	status      Status
	dueDate     DueDate
	createdAt   time.Time
	updatedAt   time.Time
}

// Getter

func (t *Todo) PK() int                    { return t.pk }
func (t *Todo) ID() uuid.UUID              { return t.id }
func (t *Todo) Title() string              { return t.title.Value() }
func (t *Todo) Description() *string       { return t.description.Value() }
func (t *Todo) Status() StatusItem         { return t.status.Value() }
func (t *Todo) StatusUpdatedAt() time.Time { return t.status.UpdatedAt() }
func (t *Todo) DueDate() *time.Time        { return t.dueDate.Value() }
func (t *Todo) CreatedAt() time.Time       { return t.createdAt }
func (t *Todo) UpdatedAt() *time.Time {
	return lo.Ternary(lo.IsEmpty(t.updatedAt), nil, lo.ToPtr(t.updatedAt))
}

func Reconstruct(
	pk int,
	id uuid.UUID,
	title string,
	description *string,
	status StatusItem,
	statusUpdatedAt time.Time,
	dueDate *time.Time,
	createdAt time.Time,
	updatedAt *time.Time) *Todo {
	return &Todo{
		pk:          pk,
		id:          id,
		title:       ReconstructTitle(title),
		description: ReconstructDescription(description),
		status:      ReconstructStatus(status, statusUpdatedAt),
		dueDate:     ReconstructDueDate(dueDate),
		createdAt:   createdAt,
		updatedAt:   lo.FromPtr(updatedAt),
	}
}

func NewTodo(command CreateCommand) (*Todo, error) {
	title, err := NewTitle(command.Title)
	if err != nil {
		return nil, err
	}
	desc, err := NewDescription(command.Description)
	if err != nil {
		return nil, err
	}
	status, err := NewStatus(command.Status)
	if err != nil {
		return nil, err
	}
	dueDate, err := NewDueDate(command.DueDate)
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
	title, err := NewTitle(command.Title)
	if err != nil {
		return err
	}
	desc, err := NewDescription(command.Description)
	if err != nil {
		return err
	}
	dueDate, err := NewDueDate(command.DueDate)
	if err != nil {
		return err
	}

	t.title = title
	t.description = desc
	t.dueDate = dueDate
	t.updatedAt = time.Now()
	return nil
}

func (t *Todo) UpdateStatus(command UpdateStatusCommand) error {
	status, err := NewStatus(command.Status)
	if err != nil {
		return err
	}

	t.status = status
	return nil
}

func (t *Todo) SetPK(pk int) { t.pk = pk }

func (t *Todo) Equals(other *Todo) bool {
	return reflect.DeepEqual(t.id, other.id)
}
