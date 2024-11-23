import { atom, useAtomValue } from 'jotai'
import { interfaces, todo } from '../types/wailsjs/go/models'

export interface ITodoApi {
  create(command: todo.CreateCommand): Promise<interfaces.CreatedResponse>
  delete(id: string): Promise<void>
  getDetails(id: string): Promise<todo.DetailsResponse>
  search(query: todo.Query): Promise<todo.ItemResponse[]>
  update(id: string, command: todo.UpdateCommand): Promise<void>
  updateStatus(id: string, command: todo.UpdateStatusCommand): Promise<void>
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
