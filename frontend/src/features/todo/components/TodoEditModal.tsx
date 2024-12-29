import { useTodoEditForm } from '@/features/todo/hooks/useTodoEditForm'
import DateTimePickerField from '@/lib/form/components/DateTimePickerField'
import FormTextField from '@/lib/form/components/FormTextField'
import { Close, Replay, Save } from '@mui/icons-material'
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
import { FormProvider } from 'react-hook-form'

interface Props {
  open: boolean
  itemId: number | null
  onClose: () => void
}

const TodoEditModal: React.FC<Props> = ({ open, itemId, onClose }) => {
  const { form, mutation } = useTodoEditForm({
    itemId,
    onSuccess: onClose,
    dialogOpened: open
  })
  const confirm = useConfirm()

  const handleClose = async (_: {}, reason: 'backdropClick' | 'escapeKeyDown') => {
    if (reason === 'backdropClick') return

    if (form.formState.isDirty) {
      try {
        await confirm({ title: '確認', description: '変更を保存しないで閉じますか？' })
      } catch {
        return
      }
    }

    onClose()
  }

  const closeDialog = () => handleClose({}, 'escapeKeyDown')

  const onReset = (e: React.FormEvent) => {
    e.preventDefault()
    form.reset()
  }

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      PaperProps={{
        component: 'form',
        onSubmit: form.handleSubmit((data) => mutation.mutate(data)),
        onReset
      }}
    >
      <DialogTitle sx={{ m: 0, p: 2 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
          {itemId ? 'Todoの編集' : 'Todoの新規作成'}
          <IconButton
            aria-label="close"
            onClick={closeDialog}
            sx={(theme) => ({ color: theme.palette.grey[500] })}
            children={<Close />}
            title="閉じる"
          />
        </Box>
      </DialogTitle>

      <DialogContent dividers>
        <FormProvider {...form}>
          <FormTextField name="title" label="タイトル" fullWidth margin="normal" />
          <FormTextField
            name="description"
            label="説明"
            fullWidth
            margin="normal"
            multiline
            rows={4}
          />
          <DateTimePickerField
            name="dueDate"
            label="期限"
            views={['year', 'day', 'hours', 'minutes']}
            textFieldProps={{ fullWidth: true, margin: 'normal' }}
            pickerProps={{ format: 'YYYY/MM/DD HH:mm', ampm: false, disablePast: true }}
          />
        </FormProvider>
      </DialogContent>

      <DialogActions>
        <Button type="reset" startIcon={<Replay />} disabled={!form.formState.isDirty}>
          リセット
        </Button>
        <Button
          type="submit"
          variant="contained"
          color="primary"
          startIcon={<Save />}
          disabled={Object.keys(form.formState.errors).length > 0 || mutation.isPending}
        >
          保存
        </Button>
      </DialogActions>
    </Dialog>
  )
}

export default TodoEditModal
