import { Box, Button, TextField } from '@mui/material'
import { useEffect, useMemo } from 'react'
import { useForm } from 'react-hook-form'
import { ApiError } from '../../api/errors'
import useTodoAtoms from '../../atoms/todo-atoms'
import { todo } from '../../types/wailsjs/go/models'
import DateTimePickerField from '../common/DateTimePickerField'

type TodoFormValue = todo.CreateCommand | todo.UpdateCommand

interface Props {
  originData?: todo.DetailsResponse | null
  onSetDirty?: (value: boolean) => void
  onSubmitFinished?: () => void
}

const TodoForm: React.FC<Props> = ({ originData, onSetDirty, onSubmitFinished }) => {
  const { createTodo, updateTodo } = useTodoAtoms()

  const defaultValues = useMemo<TodoFormValue>(
    () => ({
      title: originData?.title ?? '',
      description: originData?.description,
      dueDate: originData?.dueDate
    }),
    [originData]
  )

  const {
    register,
    handleSubmit,
    formState: { errors, isDirty },
    setError,
    reset,
    control,
    watch
  } = useForm({ defaultValues })

  // defaultValueが変更されるたびにデフォルト値を更新
  useEffect(() => reset(defaultValues), [defaultValues])

  useEffect(() => onSetDirty && onSetDirty(isDirty), [isDirty])

  const onSubmit = async (form: TodoFormValue) => {
    try {
      if (originData) {
        await updateTodo(originData.id, form)
      } else {
        await createTodo(form)
      }
      onSubmitFinished && onSubmitFinished()
    } catch (error) {
      if (error instanceof ApiError) {
        const field = error.data?.field as keyof TodoFormValue
        setError(field ?? 'root', { message: error.data?.message })
        return
      }
      throw error
    }
  }

  const onReset = (e: React.FormEvent) => {
    e.preventDefault()
    reset()
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)} onReset={onReset}>
      <TextField
        {...register('title')}
        label="Title"
        fullWidth
        margin="normal"
        error={!!errors.title}
        helperText={errors.title?.message}
        slotProps={{ inputLabel: { shrink: !!watch('title') } }}
      />
      <TextField
        {...register('description')}
        label="Description"
        fullWidth
        multiline
        margin="normal"
        rows={3}
        error={!!errors.description}
        helperText={errors.description?.message}
        slotProps={{ inputLabel: { shrink: !!watch('description') } }}
      />
      <DateTimePickerField
        name="dueDate"
        control={control}
        label="Due Date"
        views={['year', 'day', 'hours', 'minutes']}
        fullWidth
        margin="normal"
        error={!!errors.dueDate}
        helperText={errors.dueDate?.message as string}
      />

      <Box sx={{ mt: 1 }}>
        <Button type="submit" variant="contained" color="primary" sx={{ mr: 1 }}>
          Save
        </Button>
        <Button type="reset" variant="contained" color="secondary">
          Reset
        </Button>
      </Box>
    </form>
  )
}

export default TodoForm
