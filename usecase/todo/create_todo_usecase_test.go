package todo_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	infra "github.com/nftug/wails-todo-app/infrastructure/todo"
	"github.com/nftug/wails-todo-app/presentation"
	"github.com/samber/lo"
)

func TestCreateTodo(t *testing.T) {
	t.Parallel()
	a := presentation.CreateAppRootMock()

	cmdBase := todo.CreateCommand{
		Title:         "Test!",
		Description:   lo.ToPtr("Description"),
		InitialStatus: lo.ToPtr(todo.StatusValue("Todo")),
		DueDate:       lo.ToPtr(time.Now().In(Tz)),
	}
	cmds := lo.RepeatBy(4, func(_ int) todo.CreateCommand { return cmdBase })
	cmds[1].Description = nil
	cmds[2].InitialStatus = nil
	cmds[3].DueDate = nil

	// Act
	for _, cmd := range cmds {
		resp, err := a.Todo.Create(cmd)
		if err != nil {
			t.Errorf("cannot create the item: %+v", err)
		}

		// Assert
		var actual infra.TodoTable
		if err := a.DB.Where("id = ?", resp.ID).Take(&actual).Error; err != nil {
			t.Errorf("db error while asserting: %+v", err)
		}

		dueDateWant := lo.TernaryF(cmd.DueDate == nil,
			func() *time.Time { return nil }, func() *time.Time { return lo.ToPtr(cmd.DueDate.UTC()) })
		want := infra.TodoTable{
			PK:              actual.PK,
			ID:              resp.ID,
			Title:           cmd.Title,
			Description:     cmd.Description,
			Status:          lo.FromPtrOr(cmd.InitialStatus, "Todo"),
			StatusUpdatedAt: Now.UTC(),
			DueDate:         dueDateWant,
			CreatedAt:       Now.UTC(),
			UpdatedAt:       nil,
		}
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("in and want don't match.\nin: %+v\nwant: %+v", actual, want)
		}
	}
}
