import { AppBar, Box, Divider, Stack, Toolbar, Typography } from '@mui/material'
import { Provider } from 'inversify-react'
import { ConfirmProvider } from 'material-ui-confirm'
import { TodoEditModalProvider } from './components/todo/TodoEditModalContext'
import TodoForm from './components/todo/TodoForm'
import TodoList from './components/todo/TodoList'
import { container } from './inversify.config'

const App: React.FC = () => {
  return (
    <Provider container={container}>
      <ConfirmProvider defaultOptions={{ confirmationText: 'OK', cancellationText: 'キャンセル' }}>
        <TodoEditModalProvider>
          <AppBar>
            <Toolbar>
              <Typography variant="h5">Todo App</Typography>
            </Toolbar>
          </AppBar>

          <Box component="main">
            <Toolbar />
            <Stack sx={{ height: 'calc(100vh - 64px - 16px)' }}>
              <Box sx={{ mx: 2 }}>
                <TodoForm />
                <Divider sx={{ py: 1 }} />
              </Box>

              <TodoList sx={{ flexGrow: 1, overflow: 'scroll' }} />
            </Stack>
          </Box>
        </TodoEditModalProvider>
      </ConfirmProvider>
    </Provider>
  )
}

export default App
