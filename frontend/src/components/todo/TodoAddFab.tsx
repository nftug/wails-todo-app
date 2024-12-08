import { TodoEditModalContext } from '@/context/todo/TodoEditModalContext'
import { Add } from '@mui/icons-material'
import { Fab } from '@mui/material'
import { useContext } from 'react'

const TodoAddFab: React.FC = () => {
  const { openModal } = useContext(TodoEditModalContext)

  return (
    <Fab
      sx={{ position: 'fixed', bottom: 16, right: 16 }}
      title="Todoを追加"
      color="primary"
      onClick={() => openModal()}
      children={<Add />}
    />
  )
}

export default TodoAddFab
