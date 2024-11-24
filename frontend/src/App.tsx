import { Divider, Stack, Typography } from '@mui/material'
import { Provider } from 'inversify-react'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'
import { container } from './inversify.config'

const App: React.FC = () => {
  return (
    <Provider container={container}>
      <Stack>
        <Typography variant="h3">Todo App</Typography>
        <TodoForm />
        <Divider />
        <TodoList />
      </Stack>
    </Provider>
  )
}

export default App
