import { Card, CardContent, Chip, Typography } from '@mui/material'
import { useMemo } from 'react'
import useTodoAtoms from '../atoms/todo-atoms'
import type { TodoItem } from '../types/todo-dto'
import TodoForm from './TodoForm'

interface TodoItemProps {
  item: TodoItem
}

const TodoItem: React.FC<TodoItemProps> = ({ item }) => {
  const { selectTodo, selectedTodo } = useTodoAtoms()
  const isSelected = useMemo(() => selectedTodo?.id === item.id, [selectedTodo, item])

  const onClickItem = useMemo(() => {
    if (isSelected) return undefined
    return async () => {
      await selectTodo(item.id)
    }
  }, [isSelected, item])

  return (
    <Card
      onClick={onClickItem}
      style={{ marginBottom: '10px', cursor: isSelected ? undefined : 'pointer' }}
    >
      <CardContent>
        {isSelected && selectedTodo ? (
          <TodoForm originData={selectedTodo} />
        ) : (
          <>
            <Typography variant="h5">{item.title}</Typography>
            <Typography variant="body2" color="textSecondary">
              {item.description}
            </Typography>
            <Chip label={item.status} style={{ marginTop: '5px' }} />
          </>
        )}
      </CardContent>
    </Card>
  )
}

export default TodoItem
