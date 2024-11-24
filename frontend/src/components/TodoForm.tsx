import { Box, Button, SxProps, TextField, Theme } from '@mui/material'
import { useMemo } from 'react'
import { useForm } from 'react-hook-form'
import { ApiError } from '../api/errors'
import useTodoAtoms from '../atoms/todo-atoms'
import { todo } from '../types/wailsjs/go/models'
import DateTimePickerField from './common/DateTimePickerField'

type TodoFormValue = todo.CreateCommand | todo.UpdateCommand

interface Props {
  originData?: todo.DetailsResponse
  sx?: SxProps<Theme>
}

const TodoForm: React.FC<Props> = ({ originData, sx }) => {
  const { createTodo, updateTodo, selectTodo } = useTodoAtoms()

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
    formState: { errors },
    setError,
    reset,
    control
  } = useForm({ defaultValues })

  const onSubmit = async (form: TodoFormValue) => {
    try {
      if (originData) {
        await updateTodo(originData.id, form)
      } else {
        await createTodo(form)
      }

      // フォーム内容とTodoの選択をリセット
      reset()
      selectTodo(null)
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
    <Box sx={sx}>
      <form onSubmit={handleSubmit(onSubmit)} onReset={onReset}>
        <TextField
          {...register('title')}
          label="Title"
          fullWidth
          margin="normal"
          error={!!errors.title}
          helperText={errors.title?.message}
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
    </Box>
  )
}

export default TodoForm
