import { atom, useAtom } from 'jotai'
import { useTodoApi } from '../api/todo-api'
import type {
  TodoCreateCommand,
  TodoDetails,
  TodoItem,
  TodoQuery,
  TodoUpdateCommand
} from '../types/todo-dto'
import { enums } from '../types/wailsjs/go/models'

const todoListAtom = atom<TodoItem[]>([])
const queryAtom = atom<TodoQuery>({})
const selectedTodoAtom = atom<TodoDetails | null>(null)

const useTodoAtoms = () => {
  const [query, setQuery] = useAtom(queryAtom)
  const [todoList, setTodoList] = useAtom(todoListAtom)
  const [selectedTodo, setSelectedTodo] = useAtom(selectedTodoAtom)
  const api = useTodoApi()

  const updateList = async () => {
    setTodoList(await api.search(query))
  }

  const selectTodo = async (id: string | null) => {
    if (id) {
      const todo = await api.getDetails(id)
      setSelectedTodo(todo)
    } else {
      setSelectedTodo(null)
    }
  }

  const createTodo = async (command: TodoCreateCommand) => {
    const ret = await api.create(command)
    await updateList()
    return ret
  }

  const updateTodo = async (id: string, command: TodoUpdateCommand) => {
    await api.update(id, command)
    await updateList()
  }

  const updateStatus = async (id: string, status: enums.StatusValue) => {
    await api.updateStatus(id, { status })
    await updateList()
  }

  return {
    updateList,
    selectTodo,
    createTodo,
    updateTodo,
    updateStatus,
    todoList,
    selectedTodo,
    query,
    setQuery
  }
}

export default useTodoAtoms
