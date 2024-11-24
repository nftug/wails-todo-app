import { Box, Button, TextField } from '@mui/material'
import { DateTimePicker } from '@mui/x-date-pickers'
import dayjs from 'dayjs'
import { useMemo } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { ApiError } from '../api/errors'
import useTodoAtoms from '../atoms/todo-atoms'
import { todo } from '../types/wailsjs/go/models'

type TodoFormValue = todo.CreateCommand | todo.UpdateCommand

interface Props {
  originData?: todo.DetailsResponse
}

const TodoForm: React.FC<Props> = ({ originData }) => {
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
    <Box sx={{ my: 2 }}>
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
        <Controller
          name="dueDate"
          control={control}
          render={({ field: { onChange, value, ref } }) => (
            <DateTimePicker
              label="Due Date"
              onChange={onChange}
              value={value ? dayjs(value) : null}
              inputRef={ref}
              views={['year', 'day', 'hours', 'minutes']}
              slotProps={{
                textField: {
                  fullWidth: true,
                  margin: 'normal',
                  error: !!errors.dueDate,
                  helperText: errors.dueDate?.message as string
                }
              }}
            />
          )}
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
