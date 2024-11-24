import TodoAddFab from '../components/todo/TodoAddFab'
import { TodoEditModalProvider } from '../components/todo/TodoEditModalContext'
import TodoList from '../components/todo/TodoList'

const IndexPage: React.FC = () => {
  return (
    <TodoEditModalProvider>
      <TodoList />
      <TodoAddFab />
    </TodoEditModalProvider>
  )
}

export default IndexPage
