import { useTodoCommandAtoms } from '@/atoms/todo-atoms'
import { ApiError } from '@/lib/api/errors'
import { todo } from '@/types/wailsjs/go/models'
import { useEffect } from 'react'
import { UseFormReturn, useForm } from 'react-hook-form'

type TodoFormValue = todo.CreateCommand | todo.UpdateCommand

interface Props {
  originData?: todo.DetailsResponse | null
  onSubmitFinished?: () => void
  onSetDirty?: (value: boolean) => void
  children: (form: UseFormReturn<TodoFormValue>) => React.ReactNode
}

const TodoFormProvider: React.FC<Props> = ({
  originData,
  onSubmitFinished,
  onSetDirty,
  children
}) => {
  const { createTodo, updateTodo } = useTodoCommandAtoms()
  const form = useForm<TodoFormValue>()
  const {
    handleSubmit,
    setError,
    reset,
    formState: { isDirty }
  } = form

  // defaultValuesが変更されたときのリセット処理 (デフォルト値の更新)
  useEffect(() => {
    reset({
      title: originData?.title ?? '',
      description: originData?.description ?? '',
      dueDate: originData?.dueDate
    })
  }, [originData])

  console.log('rerendered')

  // isDirtyが更新されたときにPropsに通知
  useEffect(() => onSetDirty && onSetDirty(isDirty), [isDirty])

  // フォームの送信処理
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
      {children(form)}
    </form>
  )
}

export default TodoFormProvider
