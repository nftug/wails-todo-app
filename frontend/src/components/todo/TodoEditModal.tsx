import CloseIcon from '@mui/icons-material/Close'
import { Dialog, DialogContent, DialogTitle, IconButton } from '@mui/material'
import { useConfirm } from 'material-ui-confirm'
import { useEffect, useState } from 'react'
import { useTodoApi } from '../../api/todo-api'
import { todo } from '../../types/wailsjs/go/models'
import TodoForm from './TodoForm'

interface Props {
  open: boolean
  itemId: string | null
  onClose: () => void
}

const TodoEditModal: React.FC<Props> = ({ open, itemId, onClose }) => {
  const api = useTodoApi()
  const confirm = useConfirm()
  const [originData, setOriginData] = useState<todo.DetailsResponse | null>(null)
  const [isDirty, setIsDirty] = useState(false)

  const loadData = async () => {
    if (itemId) {
      setOriginData(await api.getDetails(itemId))
    } else {
      setOriginData(null)
    }
  }

  useEffect(() => {
    loadData()
  }, [itemId])

  const handleClose = async ({ reason }: { reason: 'backdropClick' | 'escapeKeyDown' }) => {
    if (reason === 'backdropClick') return

    if (isDirty) {
      try {
        await confirm({ title: '確認', description: '変更を保存しないで閉じますか？' })
        onClose()
      } catch {
        return
      }
    } else {
      onClose()
    }
  }

  return (
    <Dialog open={open} onClose={handleClose}>
      <DialogTitle sx={{ m: 0, p: 2 }} id="customized-dialog-title">
        {itemId ? 'Todoの編集' : 'Todoの新規作成'}
      </DialogTitle>
      <IconButton
        aria-label="close"
        onClick={() => handleClose({ reason: 'escapeKeyDown' })}
        sx={(theme) => ({
          position: 'absolute',
          right: 8,
          top: 8,
          color: theme.palette.grey[500]
        })}
      >
        <CloseIcon />
      </IconButton>

      <DialogContent>
        <TodoForm
          originData={originData}
          onSetDirty={(v) => setIsDirty(v)}
          onSubmitFinished={onClose}
        />
      </DialogContent>
    </Dialog>
  )
}

export default TodoEditModal
