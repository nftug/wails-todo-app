import { TextField } from '@mui/material'
import { UseFormReturn } from 'react-hook-form'
import { todo } from '../../types/wailsjs/go/models'
import DateTimePickerField from '../common/DateTimePickerField'

type TodoFormValue = todo.CreateCommand | todo.UpdateCommand

interface Props {
  context: UseFormReturn<TodoFormValue>
}

const TodoFormContent: React.FC<Props> = ({ context }) => {
  const {
    register,
    formState: { errors },
    control,
    watch
  } = context

  return (
    <>
      <TextField
        {...register('title')}
        label="タイトル"
        fullWidth
        margin="normal"
        error={!!errors.title}
        helperText={errors.title?.message}
        slotProps={{ inputLabel: { shrink: !!watch('title') } }}
      />
      <TextField
        {...register('description')}
        label="説明"
        fullWidth
        margin="normal"
        multiline
        rows={4}
        error={!!errors.description}
        helperText={errors.description?.message}
        slotProps={{ inputLabel: { shrink: !!watch('description') } }}
      />
      <DateTimePickerField
        name="dueDate"
        control={control}
        label="期限"
        views={['year', 'day', 'hours', 'minutes']}
        fullWidth
        margin="normal"
        error={!!errors.dueDate}
        helperText={errors.dueDate?.message as string}
      />
    </>
  )
}

export default TodoFormContent
