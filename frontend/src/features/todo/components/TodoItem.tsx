import { useDeleteTodo } from '@/features/todo/hooks/useDeleteTodo'
import { overflowEllipsisStyle } from '@/lib/layout/styles'
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
import { todo } from '@wailsjs/go/models'
import { useConfirm } from 'material-ui-confirm'

interface TodoItemProps {
  item: todo.ItemResponse
  onClickEdit: (itemId: number) => void
}

const TodoItem: React.FC<TodoItemProps> = ({ item, onClickEdit }) => {
  const deleteTodo = useDeleteTodo()
  const confirm = useConfirm()

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

    deleteTodo.mutate(item.id)
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
        <Button
          size="small"
          color="primary"
          onClick={() => onClickEdit(item.id)}
          startIcon={<EditIcon />}
        >
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
