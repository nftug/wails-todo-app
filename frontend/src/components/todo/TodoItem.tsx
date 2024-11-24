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
import useTodoAtoms from '../../atoms/todo-atoms'
import { todo } from '../../types/wailsjs/go/models'
import { overflowEllipsisStyle } from '../layout/styles'

interface TodoItemProps {
  item: todo.ItemResponse
  onClickEdit: (itemId: string) => void
}

const TodoItem: React.FC<TodoItemProps> = ({ item, onClickEdit }) => {
  const { deleteTodo } = useTodoAtoms()
  const confirm = useConfirm()

  const handleEditItem = () => {
    onClickEdit(item.id)
  }

  const handleDeleteItem = async () => {
    try {
      await confirm({
        title: '確認',
        content: (
          <DialogContentText sx={overflowEllipsisStyle}>
            次のアイテムを削除しますか？
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
