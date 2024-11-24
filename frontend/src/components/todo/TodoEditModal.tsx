import { Save } from '@mui/icons-material'
import CloseIcon from '@mui/icons-material/Close'
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton
} from '@mui/material'
import { useConfirm } from 'material-ui-confirm'
import { useEffect, useState } from 'react'
import { useTodoApi } from '../../api/todo-api'
import { todo } from '../../types/wailsjs/go/models'
import TodoFormContent from './TodoFormContent'
import TodoFormProvider from './TodoFormProvider'

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

  const handleClose = async (_: {}, reason: 'backdropClick' | 'escapeKeyDown') => {
    if (reason === 'backdropClick') return

    if (isDirty) {
      try {
        await confirm({ title: '確認', description: '変更を保存しないで閉じますか？' })
      } catch {
        return
      }
    }

    onClose()
  }

  const closeDialog = () => handleClose({}, 'escapeKeyDown')

  return (
    <Dialog open={open} onClose={handleClose}>
      <TodoFormProvider originData={originData} onSubmitFinished={onClose} onSetDirty={setIsDirty}>
        {(context) => (
          <>
            <DialogTitle sx={{ m: 0, p: 2 }}>
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                {itemId ? 'Todoの編集' : 'Todoの新規作成'}
                <IconButton
                  aria-label="close"
                  onClick={closeDialog}
                  sx={(theme) => ({ color: theme.palette.grey[500] })}
                  children={<CloseIcon />}
                  title="閉じる"
                />
              </Box>
            </DialogTitle>

            <DialogContent dividers>
              <TodoFormContent context={context} />
            </DialogContent>

            <DialogActions>
              <Button onClick={closeDialog}>キャンセル</Button>
              <Button type="submit" variant="contained" color="primary" startIcon={<Save />}>
                保存
              </Button>
            </DialogActions>
          </>
        )}
      </TodoFormProvider>
    </Dialog>
  )
}

export default TodoEditModal
