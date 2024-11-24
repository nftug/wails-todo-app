import { atom, useAtom } from 'jotai'
import { useTodoApi } from '../api/todo-api'
import { enums, todo } from '../types/wailsjs/go/models'

const todoListAtom = atom<todo.ItemResponse[]>([])
const queryAtom = atom<todo.Query>({})
const selectedTodoAtom = atom<todo.DetailsResponse | null>(null)

const useTodoAtoms = () => {
  const [query, setQuery] = useAtom(queryAtom)
  const [todoList, setTodoList] = useAtom(todoListAtom)
  const api = useTodoApi()

  const updateList = async () => {
    setTodoList(await api.search(query))
  }

  const createTodo = async (command: todo.CreateCommand) => {
    const ret = await api.create(command)
    await updateList()
    return ret
  }

  const updateTodo = async (id: string, command: todo.UpdateCommand) => {
    await api.update(id, command)
    await updateList()
  }

  const deleteTodo = async (id: string) => {
    await api.delete(id)
    await updateList()
  }

  const updateStatus = async (id: string, status: enums.StatusValue) => {
    await api.updateStatus(id, { status })
    await updateList()
  }

  return {
    updateList,
    createTodo,
    updateTodo,
    deleteTodo,
    updateStatus,
    todoList,
    query,
    setQuery
  }
}

export default useTodoAtoms
