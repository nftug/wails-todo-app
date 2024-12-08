import '@abraham/reflection'

import { ITodoApi, TodoApiType } from '@/lib/api/todo-api'
import WailsTodoApiService from '@/lib/api/wails-todo-api'
import { Container } from 'inversify'

const container = new Container()
container.bind<ITodoApi>(TodoApiType).to(WailsTodoApiService)

export { container }
