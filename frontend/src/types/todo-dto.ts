import type { interfaces, todo } from './wailsjs/go/models'

export type TodoItem = ClassFields<todo.ItemResponse>

export type TodoDetails = ClassFields<todo.DetailResponse>

export type TodoQuery = ClassFields<todo.Query>

export type TodoCreateCommand = ClassFields<todo.CreateCommand>

export type TodoUpdateCommand = ClassFields<todo.UpdateCommand>

export type TodoUpdateStatusCommand = ClassFields<todo.UpdateStatusCommand>

export type CreatedResponse = ClassFields<interfaces.CreatedResponse>
