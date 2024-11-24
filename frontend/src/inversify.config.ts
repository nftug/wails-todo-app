import '@abraham/reflection'
import { Container } from 'inversify'
import { ITodoApi, TodoApiType } from './api/todo-api'
import WailsTodoApiService from './api/wails-todo-api'

const container = new Container()
container.bind<ITodoApi>(TodoApiType).to(WailsTodoApiService)

export { container }
