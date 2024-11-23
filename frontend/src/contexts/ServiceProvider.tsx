import { createStore, Provider } from 'jotai'
import { todoApiAtom } from '../api/todo-api'
import useWailsTodoApi from '../api/wails-todo-api'

interface ServiceProviderProps {
  children?: React.ReactNode
}

const ServiceProvider: React.FC<ServiceProviderProps> = ({ children }) => {
  // Inject
  const store = createStore()
  store.set(todoApiAtom, useWailsTodoApi)

  return <Provider store={store} children={children} />
}

export default ServiceProvider
