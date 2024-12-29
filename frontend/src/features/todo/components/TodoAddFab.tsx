import { Add } from '@mui/icons-material'
import { Fab } from '@mui/material'

interface Props {
  onClickAdd: () => void
}

const TodoAddFab: React.FC<Props> = ({ onClickAdd }) => {
  return (
    <Fab
      sx={{ position: 'fixed', bottom: 16, right: 16 }}
      title="Todoを追加"
      color="primary"
      onClick={onClickAdd}
      children={<Add />}
    />
  )
}

export default TodoAddFab
