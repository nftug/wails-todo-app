import { Card, CardContent, Chip, Typography } from '@mui/material'
import { useEffect, useMemo, useRef } from 'react'
import useTodoAtoms from '../atoms/todo-atoms'
import { todo } from '../types/wailsjs/go/models'
import TodoForm from './TodoForm'

interface TodoItemProps {
  item: todo.ItemResponse
}

const TodoItem: React.FC<TodoItemProps> = ({ item }) => {
  const { selectTodo, selectedTodo } = useTodoAtoms()
  const isSelected = useMemo(() => selectedTodo?.id === item.id, [selectedTodo, item])
  const cardRef = useRef<HTMLDivElement>(null)

  // カードの外側クリックを検出する
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (cardRef.current && !cardRef.current.contains(event.target as Node)) {
        selectTodo(null)
      }
    }
    document.addEventListener('click', handleClickOutside)
    return () => {
      document.removeEventListener('click', handleClickOutside)
    }
  }, [])

  const onClickItem = async () => {
    if (!isSelected) await selectTodo(item.id)
  }

  return (
    <Card
      ref={cardRef}
      onClick={onClickItem}
      sx={{ marginBottom: '10px', cursor: isSelected ? undefined : 'pointer', width: 1 }}
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
