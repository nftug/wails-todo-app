import { useTodoApi } from '@/lib/api/todo-api'
import { enums, todo } from '@/types/wailsjs/go/models'
import { atom, useAtom, useAtomValue, useSetAtom } from 'jotai'

const todoListAtom = atom<todo.ItemResponse[]>([])
const queryAtom = atom<todo.Query>({})

export const useTodoQueryAtoms = () => {
  const api = useTodoApi()

  const [query, setQuery] = useAtom(queryAtom)
  const [todoList, setTodoList] = useAtom(todoListAtom)

  const updateList = async () => {
    setTodoList(await api.search(query))
  }

  return { query, setQuery, todoList, updateList }
}

export const useTodoCommandAtoms = () => {
  const api = useTodoApi()

  const query = useAtomValue(queryAtom)
  const setTodoList = useSetAtom(todoListAtom)

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

  return { createTodo, updateTodo, deleteTodo, updateStatus }
}
