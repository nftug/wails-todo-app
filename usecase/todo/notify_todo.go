package todo

import (
	"context"
	"time"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/samber/do"
)

type NotifyTodoUseCase interface {
	Execute(ctx context.Context)
}

type notifyTodoUseCase struct {
	repo         todo.TodoRepository
	notification todo.TodoNotificationSender
}

func NewNotifyTodoUseCase(i *do.Injector) (NotifyTodoUseCase, error) {
	return &notifyTodoUseCase{
		do.MustInvoke[todo.TodoRepository](i),
		do.MustInvoke[todo.TodoNotificationSender](i),
	}, nil
}

func (u *notifyTodoUseCase) Execute(ctx context.Context) {
	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()

		for range tick.C {
			items, err := u.repo.FindAllForNotification(ctx, time.Now())
			if err != nil {
				panic(err)
			}

			for _, item := range items {
				if err := u.notification.Send(ctx, item); err != nil {
					panic(err)
				}

				item.SetNotifiedAt(time.Now())
				if err := u.repo.Save(ctx, item); err != nil {
					panic(err)
				}
			}
		}
	}()
}
