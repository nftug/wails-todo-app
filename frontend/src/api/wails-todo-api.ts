import {
  Create,
  Delete,
  GetDetails,
  Search,
  Update,
  UpdateStatus
} from '../types/wailsjs/go/app/TodoApp'
import { todo } from '../types/wailsjs/go/models'
import { handleApiError } from './errors'
import { ITodoApi } from './todo-api'

const useWailsTodoApi: ITodoApi = {
  create: (command) =>
    handleApiError(async () => await Create(todo.CreateCommand.createFrom(command))),
  delete: (id) => handleApiError(async () => await Delete(id)),
  getDetails: (id) => handleApiError(async () => await GetDetails(id)),
  search: (query) => handleApiError(async () => await Search(todo.Query.createFrom(query))),
  update: (id, command) =>
    handleApiError(async () => await Update(id, todo.UpdateCommand.createFrom(command))),
  updateStatus: (id, command) =>
    handleApiError(async () => await UpdateStatus(id, todo.UpdateStatusCommand.createFrom(command)))
}

export default useWailsTodoApi
