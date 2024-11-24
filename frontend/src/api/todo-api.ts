import { useInjection } from 'inversify-react'
import { interfaces, todo } from '../types/wailsjs/go/models'

export interface ITodoApi {
  create(command: todo.CreateCommand): Promise<interfaces.CreatedResponse>
  delete(id: string): Promise<void>
  getDetails(id: string): Promise<todo.DetailsResponse>
  search(query: todo.Query): Promise<todo.ItemResponse[]>
  update(id: string, command: todo.UpdateCommand): Promise<void>
  updateStatus(id: string, command: todo.UpdateStatusCommand): Promise<void>
}

export const TodoApiType = Symbol.for('TodoApi')

export const useTodoApi = () => useInjection<ITodoApi>(TodoApiType)