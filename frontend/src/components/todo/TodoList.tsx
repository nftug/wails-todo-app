import { Box, List, ListItem, SxProps, Theme, Typography } from '@mui/material'
import { useEffect } from 'react'
import useTodoAtoms from '../../atoms/todo-atoms'
import TodoItem from './TodoItem'

interface Props {
  sx?: SxProps<Theme>
}

const TodoList: React.FC<Props> = ({ sx }) => {
  const { todoList, updateList, query } = useTodoAtoms()

  useEffect(() => {
    updateList()
  }, [query])

  return (
    <List sx={sx}>
      {todoList.length === 0 ? (
        <Box display="flex" justifyContent="center" alignItems="center" height={1}>
          <Typography variant="h6" color="textSecondary">
            Todoがありません。
          </Typography>
        </Box>
      ) : (
        todoList.map((item) => (
          <ListItem key={item.id}>
            <TodoItem item={item} />
          </ListItem>
        ))
      )}
    </List>
  )
}

export default TodoList
