package todo

import (
	"context"

	"github.com/gen2brain/beeep"
	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/domain/todo/enums"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type todoNotificationSender struct {
	notifiedItems []*todo.ItemResponse
}

func NewTodoNotificationSender(i *do.Injector) (todo.TodoNotificationSender, error) {
	return &todoNotificationSender{[]*todo.ItemResponse{}}, nil
}

func (t *todoNotificationSender) Send(ctx context.Context, item *todo.ItemResponse) error {
	if lo.ContainsBy(t.notifiedItems,
		func(x *todo.ItemResponse) bool { return x.ID == item.ID }) {
		return nil
	}

	beeep.Alert(item.Title, lo.FromPtrOr(item.Description, "No description"), "")

	runtime.Show(ctx)
	runtime.EventsEmit(ctx, string(enums.NotifyTodo), item)

	return nil
}
