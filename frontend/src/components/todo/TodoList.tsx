import { Box, List, ListItem, SxProps, Theme, Typography } from '@mui/material'
import { useEffect } from 'react'
import useTodoAtoms from '../../atoms/todo-atoms'
import { TodoEditModalProvider, useTodoEditModal } from './TodoEditModalContext'
import TodoItem from './TodoItem'

interface Props {
  sx?: SxProps<Theme>
}

const TodoList: React.FC<Props> = ({ sx }) => {
  const { todoList, updateList } = useTodoAtoms()
  const { openModal } = useTodoEditModal()

  useEffect(() => {
    updateList()
  }, [])

  return (
    <TodoEditModalProvider>
      <List sx={sx}>
        {todoList.length === 0 ? (
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
              height: '100%'
            }}
          >
            <Typography variant="h6" color="textSecondary">
              TODOがありません。
            </Typography>
          </Box>
        ) : (
          todoList.map((item) => (
            <ListItem key={item.id}>
              <TodoItem item={item} onClickEdit={(id) => openModal(id)} />
            </ListItem>
          ))
        )}
      </List>
    </TodoEditModalProvider>
  )
}

export default TodoList
