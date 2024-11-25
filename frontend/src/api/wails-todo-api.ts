import { injectable } from 'inversify'
import {
  Create,
  Delete,
  GetDetails,
  Search,
  Update,
  UpdateStatus
} from '../types/wailsjs/go/app/TodoApp'
import { dtos, todo } from '../types/wailsjs/go/models'
import { handleApiError } from './errors'
import { ITodoApi } from './todo-api'

@injectable()
export default class WailsTodoApiService implements ITodoApi {
  async create(command: todo.CreateCommand): Promise<dtos.CreatedResponse> {
    return await handleApiError(async () => await Create(command))
  }
  async update(id: string, command: todo.UpdateCommand): Promise<void> {
    await handleApiError(async () => await Update(id, command))
  }
  async delete(id: string): Promise<void> {
    await handleApiError(async () => await Delete(id))
  }
  async updateStatus(id: string, command: todo.UpdateStatusCommand): Promise<void> {
    await handleApiError(async () => await UpdateStatus(id, command))
  }
  async getDetails(id: string): Promise<todo.DetailsResponse> {
    return await handleApiError(async () => await GetDetails(id))
  }
  async search(query: todo.Query): Promise<todo.ItemResponse[]> {
    return await handleApiError(async () => await Search(query))
  }
}
