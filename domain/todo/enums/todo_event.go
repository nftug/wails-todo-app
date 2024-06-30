package enums

type TodoEvent string

const (
	NotifyTodo = TodoEvent("NotifyTodo")
)

var TodoEvents = []TodoEvent{NotifyTodo}

func (e TodoEvent) TSName() string { return string(e) }
