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

func TestCreateTodoValidCase(t *testing.T) {
	var cmdBase = todo.CreateCommand{
		Title:         "Test!",
		Description:   lo.ToPtr("Description"),
		InitialStatus: lo.ToPtr(todo.StatusValue("Todo")),
		DueDate:       lo.ToPtr(time.Now().In(Tz)),
	}

	tests := []struct {
		name  string
		value todo.CreateCommand
	}{
		{"空欄なし", cmdBase},
		{"説明欄が空欄", arrange(cmdBase, func(c *todo.CreateCommand) { c.Description = nil })},
		{"ステータスが空欄", arrange(cmdBase, func(c *todo.CreateCommand) { c.InitialStatus = nil })},
		{"期限が空欄", arrange(cmdBase, func(c *todo.CreateCommand) { c.DueDate = nil })},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			a := presentation.CreateAppRootMock()
			cmd := test.value

			// Act
			resp, err := a.Todo.Create(cmd)
			if err != nil {
				t.Errorf("error: %+v", err)
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
		})

	}
}
