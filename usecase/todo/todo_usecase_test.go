package todo_test

import (
	"context"
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/nftug/wails-todo-app/domain/todo"
	infra "github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/nftug/wails-todo-app/usecase"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var Now time.Time
var Tz *time.Location = time.Local

func TestMain(t *testing.M) {
	// Setup
	Now = time.Now().In(Tz)
	flextime.Fix(Now)

	t.Run()

	// Teardown
}

func Arrange[T any](base T, f func(c *T)) T {
	c := base
	f(&c)
	return c
}

// 正常系
func TestCreateTodo(t *testing.T) {
	var cmd = todo.CreateCommand{
		Title:         "Test!",
		Description:   lo.ToPtr("Description"),
		InitialStatus: lo.ToPtr(todo.StatusValue("Todo")),
		DueDate:       lo.ToPtr(time.Now().In(Tz)),
	}

	// Arrange
	a := usecase.InitUseCaseAdapterMock()

	// Act
	resp, err := a.CreateTodo.Execute(cmd, context.Background())
	assert.NoError(t, err)

	// Assert
	var actual infra.TodoTable
	err = a.DB.Where("id = ?", resp.ID).Take(&actual).Error
	assert.NoError(t, err)

	dueDateWant := lo.TernaryF(cmd.DueDate == nil,
		func() *time.Time { return nil }, func() *time.Time { return lo.ToPtr(cmd.DueDate.UTC()) })

	assert.Equal(t, 1, actual.PK)
	assert.Equal(t, resp.ID, actual.ID)
	assert.Equal(t, cmd.Title, actual.Title)
	assert.Equal(t, cmd.Description, actual.Description)
	assert.Equal(t, lo.FromPtrOr(cmd.InitialStatus, "Todo"), actual.Status)
	assert.Equal(t, Now.UTC(), actual.StatusUpdatedAt)
	assert.Equal(t, dueDateWant, actual.DueDate)
	assert.Equal(t, Now.UTC(), actual.CreatedAt)
	assert.Equal(t, lo.Empty[*time.Time](), actual.UpdatedAt)
}
