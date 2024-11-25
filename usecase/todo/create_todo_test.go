package todo_test

import (
	"context"
	"testing"
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	todoInfra "github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/nftug/wails-todo-app/interfaces/enums"
	"github.com/nftug/wails-todo-app/library/testutil"
	todoUseCase "github.com/nftug/wails-todo-app/usecase/todo"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// 正常系
func TestCreateTodo(t *testing.T) {
	// Arrange
	injector := testutil.GetInjector(t)
	uc := do.MustInvoke[todoUseCase.CreateTodoUseCase](injector)

	var cmd = todo.CreateCommand{
		Title:         "Test!",
		Description:   lo.ToPtr("Description"),
		InitialStatus: lo.ToPtr(enums.StatusValue("Todo")),
		DueDate:       lo.ToPtr(time.Now().In(Tz)),
	}

	// Act
	resp, err := uc.Execute(context.Background(), cmd)
	assert.NoError(t, err)

	// Assert
	var actual todoInfra.TodoDBSchema
	db := do.MustInvoke[*gorm.DB](injector)
	err = db.Where("id = ?", resp.ID).Take(&actual).Error
	assert.NoError(t, err)

	dueDateWant := lo.TernaryF(cmd.DueDate == nil,
		func() *time.Time { return nil }, func() *time.Time { return lo.ToPtr(cmd.DueDate.UTC()) })

	assert.Equal(t, 1, actual.PK)
	assert.Equal(t, resp.ID, actual.ID.String())
	assert.Equal(t, cmd.Title, actual.Title)
	assert.Equal(t, cmd.Description, actual.Description)
	assert.Equal(t, lo.FromPtrOr(cmd.InitialStatus, "Todo"), actual.Status)
	assert.Equal(t, Now.UTC(), actual.StatusUpdatedAt)
	assert.Equal(t, dueDateWant, actual.DueDate)
	assert.Equal(t, Now.UTC(), actual.CreatedAt)
	assert.Equal(t, lo.Empty[*time.Time](), actual.UpdatedAt)
}
