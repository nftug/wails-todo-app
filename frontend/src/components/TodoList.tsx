import { List, ListItem, SxProps, Theme } from '@mui/material'
import { useEffect } from 'react'
import useTodoAtoms from '../atoms/todo-atoms'
import TodoItem from './TodoItem'

interface Props {
  sx?: SxProps<Theme>
}

const TodoList: React.FC<Props> = ({ sx }) => {
  const { todoList, updateList } = useTodoAtoms()

  useEffect(() => {
    updateList()
  }, [])

  return (
    <List sx={sx}>
      {todoList.map((item) => (
        <ListItem key={item.id}>
          <TodoItem item={item} />
        </ListItem>
      ))}
    </List>
  )
}

export default TodoList
