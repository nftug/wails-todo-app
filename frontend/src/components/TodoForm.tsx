import { Button, TextField } from '@mui/material'
import { useMemo } from 'react'
import { useForm } from 'react-hook-form'
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
    reset
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

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
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
        error={!!errors.description}
        helperText={errors.description?.message}
      />
      {/*
      <TextField
        {...register('dueDate')}
        label="Due Date"
        type="date"
        InputLabelProps={{ shrink: true }}
        fullWidth
        margin="normal"
        error={!!errors.dueDate}
        helperText={errors.dueDate?.message as string}
      />
      */}
      <Button type="submit" variant="contained" color="primary">
        Save
      </Button>
    </form>
  )
}

export default TodoForm
