package todo

import (
	"context"
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/domain/todo/enums"
	"github.com/samber/do"
	"github.com/samber/lo"
)

type NotifyTodoUseCase interface {
	Execute(ctx context.Context)
}

type notifyTodoUseCase struct {
	query        todo.TodoQueryService
	notification todo.TodoNotificationSender
}

func NewNotifyTodoUseCase(i *do.Injector) (NotifyTodoUseCase, error) {
	return &notifyTodoUseCase{
		do.MustInvoke[todo.TodoQueryService](i),
		do.MustInvoke[todo.TodoNotificationSender](i),
	}, nil
}

func (u *notifyTodoUseCase) Execute(ctx context.Context) {
	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()

		for range tick.C {
			query := todo.Query{
				Status:     lo.ToPtr(enums.StatusTodo),
				After:      lo.ToPtr(time.Now()),
				IsNotified: lo.ToPtr(false),
			}
			items, err := u.query.FindAll(ctx, query)
			if err != nil {
				panic(err)
			}

			for _, item := range items {
				if err := u.notification.Send(ctx, item); err != nil {
					panic(err)
				}
			}
		}
	}()
}
