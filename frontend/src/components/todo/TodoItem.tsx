import DeleteIcon from '@mui/icons-material/Delete'
import EditIcon from '@mui/icons-material/Edit'
import { Box, Button, Card, CardActions, CardContent, Chip, Typography } from '@mui/material'
import { useRef } from 'react'
import useTodoAtoms from '../../atoms/todo-atoms'
import { todo } from '../../types/wailsjs/go/models'

interface TodoItemProps {
  item: todo.ItemResponse
  onClickEdit: (itemId: string) => void
}

const TodoItem: React.FC<TodoItemProps> = ({ item, onClickEdit }) => {
  const { deleteTodo } = useTodoAtoms()
  const cardRef = useRef<HTMLDivElement>(null)

  const handleEditItem = () => {
    onClickEdit(item.id)
  }

  const handleDeleteItem = async () => {
    await deleteTodo(item.id)
  }

  return (
    <Card ref={cardRef} sx={{ marginBottom: '10px', width: 1 }}>
      <CardContent>
        <Box sx={{ width: 1 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', mt: 1, mb: 3 }}>
            <Chip label={item.status} sx={{ mr: 1 }} />
            <Typography
              variant="h5"
              overflow="hidden"
              textOverflow="ellipsis"
              whiteSpace="nowrap"
              maxWidth={1}
            >
              {item.title}
            </Typography>
          </Box>

          <Box sx={{ my: 1, mx: 1 }}>
            <Typography
              variant="body2"
              color="textSecondary"
              overflow="hidden"
              textOverflow="ellipsis"
              whiteSpace="nowrap"
              maxWidth={1}
            >
              {item.description ?? 'No description'}
            </Typography>
          </Box>
        </Box>
      </CardContent>

      <CardActions>
        <Button size="small" color="primary" onClick={handleEditItem} startIcon={<EditIcon />}>
          Edit
        </Button>
        <Button size="small" color="error" onClick={handleDeleteItem} startIcon={<DeleteIcon />}>
          Delete
        </Button>
      </CardActions>
    </Card>
  )
}

export default TodoItem
