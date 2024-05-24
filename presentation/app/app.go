package app

import (
	"context"

	"github.com/nftug/wails-todo-app/presentation/types/dialog"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App { return &App{} }

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ShowMessageDialog(opt dialog.DialogOptions) dialog.DialogButton {
	ret, _ := runtime.MessageDialog(a.ctx, opt.ToRuntimeOptions())
	return dialog.GetDialogButtonResult(ret)
}
