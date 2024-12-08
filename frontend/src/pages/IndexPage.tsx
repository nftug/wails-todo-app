import TodoAddFab from '@/components/todo/TodoAddFab'
import TodoList from '@/components/todo/TodoList'
import { TodoEditModalProvider } from '@/context/todo/TodoEditModalContext'
import { fullViewHeightStyle } from '@/lib/layout/styles'

const IndexPage: React.FC = () => {
  return (
    <TodoEditModalProvider>
      <TodoList sx={fullViewHeightStyle} />
      <TodoAddFab />
    </TodoEditModalProvider>
  )
}

export default IndexPage
