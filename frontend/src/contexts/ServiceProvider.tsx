import { createStore, Provider } from 'jotai'
import { todoApiAtom } from '../api/todo-api'
import wailsTodoApi from '../api/wails-todo-api'

interface ServiceProviderProps {
  children?: React.ReactNode
}

const ServiceProvider: React.FC<ServiceProviderProps> = ({ children }) => {
  // Inject
  const store = createStore()
  store.set(todoApiAtom, wailsTodoApi)

  return <Provider store={store} children={children} />
}

export default ServiceProvider
