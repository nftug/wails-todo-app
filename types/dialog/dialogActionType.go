package dialog

import (
	"runtime"

	"github.com/samber/lo"
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

		return lo.
			Map(buttons, func(b DialogButton, _ int) string {
				return lo.Ternary(runtime.GOOS == "darwin", MacButtonNames[b], string(b))
			})
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
		return lo.Ternary(runtime.GOOS == "darwin", MacButtonNames[action], string(action))
	}
}
