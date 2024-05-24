package dialog

import (
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DialogType string

const (
	InfoDialog     DialogType = DialogType(runtime.InfoDialog)
	WarningDialog  DialogType = DialogType(runtime.WarningDialog)
	ErrorDialog    DialogType = DialogType(runtime.ErrorDialog)
	QuestionDialog DialogType = DialogType(runtime.QuestionDialog)
)

var AllDialogTypes = []DialogType{InfoDialog, WarningDialog, ErrorDialog, QuestionDialog}

func (t DialogType) TSName() string {
	typeStr := string(t)
	return strings.ToUpper(typeStr[:1]) + typeStr[1:]
}

func (t *DialogType) ToRuntimeType() runtime.DialogType {
	if t == nil {
		return runtime.InfoDialog
	} else {
		return runtime.DialogType(*t)
	}
}
