import { Divider, Stack, Typography } from '@mui/material'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'
import ServiceProvider from './contexts/ServiceProvider'

const App: React.FC = () => {
  return (
    <ServiceProvider>
      <Stack>
        <Typography variant="h3">Todo App</Typography>
        <TodoForm />
        <Divider />
        <TodoList />
      </Stack>
    </ServiceProvider>
  )
}

export default App
