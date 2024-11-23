import { List, ListItem } from '@mui/material'
import { useEffect } from 'react'
import useTodoAtoms from '../atoms/todo-atoms'
import TodoItem from './TodoItem'

const TodoList: React.FC = () => {
  const { todoList, updateList } = useTodoAtoms()

  useEffect(() => {
    updateList()
  }, [])

  return (
    <List>
      {todoList.map((item) => (
        <ListItem key={item.id}>
          <TodoItem item={item} />
        </ListItem>
      ))}
    </List>
  )
}

export default TodoList
