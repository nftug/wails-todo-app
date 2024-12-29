import { handleApiError } from '@/lib/api/errors'
import { ITodoApi } from '@/lib/api/todo-api'
import {
  CreateTodo,
  DeleteTodo,
  GetTodoDetails,
  GetTodoList,
  UpdateTodo,
  UpdateTodoStatus
} from '@/types/wailsjs/go/app/TodoApp'
import { dtos, todo } from '@/types/wailsjs/go/models'
import { injectable } from 'inversify'

@injectable()
export default class WailsTodoApiService implements ITodoApi {
  async create(command: todo.CreateCommand): Promise<dtos.CreatedResponse> {
    return await handleApiError(async () => await CreateTodo(command))
  }
  async update(id: number, command: todo.UpdateCommand): Promise<void> {
    await handleApiError(async () => await UpdateTodo(id, command))
  }
  async delete(id: number): Promise<void> {
    await handleApiError(async () => await DeleteTodo(id))
  }
  async updateStatus(id: number, command: todo.UpdateStatusCommand): Promise<void> {
    await handleApiError(async () => await UpdateTodoStatus(id, command))
  }
  async getDetails(id: number): Promise<todo.DetailsResponse> {
    return await handleApiError(async () => await GetTodoDetails(id))
  }
  async search(query: todo.Query): Promise<todo.ItemResponse[]> {
    return await handleApiError(async () => await GetTodoList(query))
  }
}
