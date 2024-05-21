package todo

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	id          uuid.UUID
	title       Title
	description Description
	status      Status
	createdAt   time.Time
	updatedAt   *time.Time
}

// Getter

func (t *Todo) ID() uuid.UUID              { return t.id }
func (t *Todo) Title() string              { return t.title.Value() }
func (t *Todo) Description() *string       { return t.description.Value() }
func (t *Todo) Status() StatusItem         { return t.status.Value() }
func (t *Todo) StatusUpdatedAt() time.Time { return t.status.UpdatedAt() }
func (t *Todo) CreatedAt() time.Time       { return t.createdAt }
func (t *Todo) UpdatedAt() *time.Time      { return t.updatedAt }

func Reconstruct(
	id uuid.UUID,
	title string,
	description *string,
	status StatusItem,
	statusUpdatedAt time.Time,
	createdAt time.Time,
	updatedAt *time.Time,
) *Todo {
	return &Todo{
		id:          id,
		title:       ReconstructTitle(title),
		description: ReconstructDescription(description),
		status:      ReconstructStatus(status, statusUpdatedAt),
		createdAt:   createdAt,
		updatedAt:   updatedAt,
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

	return &Todo{
		id:          uuid.New(),
		title:       *title,
		description: *desc,
		status:      *status,
		createdAt:   time.Now(),
	}, nil
}

func (t *Todo) Edit(command EditCommand) error {
	title, err := NewTitle(command.Title)
	if err != nil {
		return err
	}
	desc, err := NewDescription(command.Description)
	if err != nil {
		return err
	}
	updatedAt := time.Now()

	t.title = *title
	t.description = *desc
	t.updatedAt = &updatedAt
	return nil
}

func (t *Todo) UpdateStatus(command UpdateStatusCommand) error {
	status, err := NewStatus(command.Status)
	if err != nil {
		return err
	}

	t.status = *status
	return nil
}

func (t *Todo) Equals(other *Todo) bool {
	return reflect.DeepEqual(t.id, other.id)
}
