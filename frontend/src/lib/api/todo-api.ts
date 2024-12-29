import { dtos, todo } from '@/types/wailsjs/go/models'
import { useInjection } from 'inversify-react'

export interface ITodoApi {
  create(command: todo.CreateCommand): Promise<dtos.CreatedResponse>
  delete(id: number): Promise<void>
  getDetails(id: number): Promise<todo.DetailsResponse>
  search(query: todo.Query): Promise<todo.ItemResponse[]>
  update(id: number, command: todo.UpdateCommand): Promise<void>
  updateStatus(id: number, command: todo.UpdateStatusCommand): Promise<void>
}

export const TodoApiType = Symbol.for('TodoApi')

export const useTodoApi = () => useInjection<ITodoApi>(TodoApiType)
