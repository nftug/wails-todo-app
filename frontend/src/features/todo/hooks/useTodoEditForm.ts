import { todoFieldSchema, TodoFieldSchemaType } from '@/features/todo/todoFieldSchema'
import { ApiError, handleApiError } from '@/lib/api/errors'
import { CreateTodo, GetTodoDetails, UpdateTodo } from '@/types/wailsjs/go/app/TodoApp'
import { todo } from '@/types/wailsjs/go/models'
import { yupResolver } from '@hookform/resolvers/yup'
import { useMutation, useQuery, useQueryClient, UseQueryResult } from '@tanstack/react-query'
import { useConfirm } from 'material-ui-confirm'
import { useEffect } from 'react'
import { useForm } from 'react-hook-form'

type UseTodoEditFormOptions = {
  itemId: number | null
  onSuccess: () => void
  dialogOpened: boolean
}

export const useTodoEditForm = ({ itemId, onSuccess, dialogOpened }: UseTodoEditFormOptions) => {
  const queryClient = useQueryClient()
  const confirm = useConfirm()

  // Query
  const query = useQuery({
    queryKey: ['todo', 'details', itemId],
    queryFn: async () => {
      if (!itemId) return
      return await handleApiError(async () => await GetTodoDetails(itemId))
    },
    enabled: !!itemId
  }) as UseQueryResult<todo.DetailsResponse, ApiError>

  useEffect(() => {
    if (!query.error) return
    confirm({ title: 'エラー', description: query.error.message, hideCancelButton: true })
  }, [query.error])

  // Form
  const form = useForm({
    resolver: yupResolver(todoFieldSchema),
    mode: 'onChange'
  })

  useEffect(() => {
    const getFields = (data: todo.DetailsResponse): TodoFieldSchemaType => ({
      title: data.title,
      description: data.description ?? '',
      dueDate: data.dueDate
    })
    form.reset(query.data ? getFields(query.data) : todoFieldSchema.getDefault())
  }, [query.data, dialogOpened])

  // Mutation
  const mutation = useMutation({
    mutationFn: async (fields: TodoFieldSchemaType) => {
      return await handleApiError(async () =>
        itemId ? await UpdateTodo(itemId, fields) : await CreateTodo(fields)
      )
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['todo'],
        predicate: ({ queryKey }) => queryKey[1] === 'list' || queryKey.at(-1) === itemId
      })
      onSuccess()
      setTimeout(() => mutation.reset(), 500)
    },
    onError: (error: ApiError) => {
      form.setError(error.data?.field ?? 'root', { message: error.data?.message })
    }
  })

  return { query, form, mutation }
}
