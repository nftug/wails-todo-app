package dialog

type DialogButton string

const (
	Ok     DialogButton = "Ok"
	Cancel DialogButton = "Cancel"
	Yes    DialogButton = "Yes"
	No     DialogButton = "No"
)

var AllDialogButtons = []DialogButton{Ok, Cancel, Yes, No}

func (t DialogButton) TSName() string {
	return string(t)
}

var MacButtonNames = map[DialogButton]string{
	Ok:     "OK",
	Cancel: "Cancel",
	Yes:    "Yes",
	No:     "No",
}

func GetDialogButtonResult(button string) DialogButton {
	switch button {
	case string(Ok), string(Cancel), string(Yes), string(No):
		return DialogButton(button)
	default:
		keys := keysByValue(MacButtonNames, button)
		if len(keys) > 0 {
			return keys[0]
		} else {
			return DialogButton(button)
		}
	}
}

func keysByValue[T comparable, V comparable](m map[T]V, value V) []T {
	var keys []T
	for k, v := range m {
		if value == v {
			keys = append(keys, k)
		}
	}
	return keys
}
