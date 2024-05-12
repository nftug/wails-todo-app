package dialog

import (
	"runtime"

	"github.com/ahmetb/go-linq/v3"
)

type DialogActionType string

const (
	OkAction       DialogActionType = "Ok"
	OkCancelAction DialogActionType = "OkCancel"
	YesNoAction    DialogActionType = "YesNo"
)

var AllDialogActionTypes = []DialogActionType{OkAction, OkCancelAction, YesNoAction}

func (t DialogActionType) TSName() string {
	return string(t)
}

func (t *DialogActionType) GetButtons() []string {
	if t == nil {
		return []string{}
	} else {
		var buttons []DialogButton
		switch *t {
		case OkCancelAction:
			buttons = []DialogButton{Cancel, Ok}
		case YesNoAction:
			buttons = []DialogButton{No, Yes}
		default:
			buttons = []DialogButton{}
		}

		var results []string
		if runtime.GOOS == "darwin" {
			linq.From(buttons).
				SelectT(func(b DialogButton) string { return MacButtonNames[b] }).
				ToSlice(&results)
		} else {
			linq.From(buttons).
				SelectT(func(b DialogButton) string { return string(b) }).
				ToSlice(&results)
		}

		return results
	}
}

func (t *DialogActionType) GetDefaultButton() string {
	if t == nil {
		return *new(string)
	} else {
		var action DialogButton
		switch *t {
		case OkCancelAction:
			action = Ok
		case YesNoAction:
			action = Yes
		default:
			action = DialogButton(*new(string))
		}

		var button string
		if runtime.GOOS == "darwin" {
			button = MacButtonNames[action]
		} else {
			button = string(action)
		}

		return button
	}
}
