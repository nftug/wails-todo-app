import {
  Create,
  Delete,
  GetDetails,
  Search,
  Update,
  UpdateStatus
} from '../types/wailsjs/go/app/TodoApp'
import { handleApiError } from './errors'
import { ITodoApi } from './todo-api'

const wailsTodoApi: ITodoApi = {
  create: (command) => handleApiError(async () => await Create(command)),
  delete: (id) => handleApiError(async () => await Delete(id)),
  getDetails: (id) => handleApiError(async () => await GetDetails(id)),
  search: (query) => handleApiError(async () => await Search(query)),
  update: (id, command) => handleApiError(async () => await Update(id, command)),
  updateStatus: (id, command) => handleApiError(async () => await UpdateStatus(id, command))
}

export default wailsTodoApi
