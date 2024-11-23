package todo

import (
	"context"

	"github.com/gen2brain/beeep"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces/enums"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type todoNotificationSender struct{}

func NewTodoNotificationSender(i *do.Injector) (todo.TodoNotificationSender, error) {
	return &todoNotificationSender{}, nil
}

func (t *todoNotificationSender) Send(ctx context.Context, item *todo.Todo) error {
	dto := todo.ItemResponse{
		ID:          item.ID().String(),
		Title:       item.Title(),
		Description: item.Description(),
		Status:      item.Status(),
		NotifiedAt:  item.NotifiedAt(),
		DueDate:     item.DueDate(),
	}

	beeep.Alert(dto.Title, lo.FromPtrOr(dto.Description, "No description"), "")

	runtime.Show(ctx)
	runtime.EventsEmit(ctx, string(enums.NotifyTodo), dto)

	return nil
}
