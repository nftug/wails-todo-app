package dialog

import (
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DialogOptions struct {
	Message    string            `json:"message"`
	Title      *string           `json:"title"`
	Type       *DialogType       `json:"type"`
	ActionType *DialogActionType `json:"actionType"`
}

func (o *DialogOptions) ToRuntimeOptions() runtime.MessageDialogOptions {
	message := strings.ToUpper(o.Message[:1]) + o.Message[1:]

	var title string
	if o.Title != nil {
		title = *o.Title
	} else {
		title = "Message"
	}

	return runtime.MessageDialogOptions{
		Type:          o.Type.ToRuntimeType(),
		Title:         title,
		Message:       message,
		Buttons:       o.ActionType.GetButtons(),
		DefaultButton: o.ActionType.GetDefaultButton(),
	}
}
