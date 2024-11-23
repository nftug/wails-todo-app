import { atom, useAtomValue } from 'jotai'
import type {
  CreatedResponse,
  TodoCreateCommand,
  TodoDetails,
  TodoItem,
  TodoQuery,
  TodoUpdateCommand,
  TodoUpdateStatusCommand
} from '../types/todo-dto'

export interface ITodoApi {
  create(command: TodoCreateCommand): Promise<CreatedResponse>
  delete(id: string): Promise<void>
  getDetails(id: string): Promise<TodoDetails>
  search(query: TodoQuery): Promise<TodoItem[]>
  update(id: string, command: TodoUpdateCommand): Promise<void>
  updateStatus(id: string, command: TodoUpdateStatusCommand): Promise<void>
}

export const todoApiAtom = atom<ITodoApi>({
  create: () => {
    throw new Error('Service is not set.')
  },
  delete: () => {
    throw new Error('Service is not set.')
  },
  getDetails: () => {
    throw new Error('Service is not set.')
  },
  search: () => {
    throw new Error('Service is not set.')
  },
  update: () => {
    throw new Error('Service is not set.')
  },
  updateStatus: () => {
    throw new Error('Service is not set.')
  }
})

export const useTodoApi = () => useAtomValue(todoApiAtom)
