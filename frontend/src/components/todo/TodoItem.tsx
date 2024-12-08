import { useTodoCommandAtoms } from '@/atoms/todo-atoms'
import { TodoEditModalContext } from '@/context/todo/TodoEditModalContext'
import { overflowEllipsisStyle } from '@/lib/layout/styles'
import { todo } from '@/types/wailsjs/go/models'
import DeleteIcon from '@mui/icons-material/Delete'
import EditIcon from '@mui/icons-material/Edit'
import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Chip,
  DialogContentText,
  Typography
} from '@mui/material'
import { useConfirm } from 'material-ui-confirm'
import { useContext } from 'react'

interface TodoItemProps {
  item: todo.ItemResponse
}

const TodoItem: React.FC<TodoItemProps> = ({ item }) => {
  const { deleteTodo } = useTodoCommandAtoms()
  const confirm = useConfirm()
  const { openModal } = useContext(TodoEditModalContext)

  const handleEditItem = () => {
    openModal(item.id)
  }

  const handleDeleteItem = async () => {
    try {
      await confirm({
        title: 'Todoの削除',
        content: (
          <DialogContentText sx={overflowEllipsisStyle}>
            次のTodoを削除しますか？
            <br />
            {item.title}
          </DialogContentText>
        )
      })
    } catch {
      return
    }
    await deleteTodo(item.id)
  }

  return (
    <Card sx={{ marginBottom: '10px', width: 1 }}>
      <CardContent>
        <Box sx={{ width: 1 }}>
          <Box sx={{ display: 'flex', alignItems: 'center', mt: 1, mb: 3 }}>
            <Chip label={item.status} sx={{ mr: 1 }} />
            <Typography variant="h5" sx={overflowEllipsisStyle}>
              {item.title}
            </Typography>
          </Box>

          <Box sx={{ my: 1, mx: 1 }}>
            <Typography variant="body2" color="textSecondary" sx={overflowEllipsisStyle}>
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
