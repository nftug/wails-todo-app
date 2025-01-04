import { handleApiError } from '@/lib/api/errors'
import { useQuery } from '@tanstack/react-query'
import { GetTodoList } from '@wailsjs/go/app/TodoApp'
import { todo } from '@wailsjs/go/models'
import { useConfirm } from 'material-ui-confirm'
import { useEffect } from 'react'

export const useTodoQuery = (query: todo.Query) => {
  const { data, error, isPending } = useQuery({
    queryKey: ['todo', 'list', query],
    queryFn: async () => {
      return await handleApiError(async () => await GetTodoList(query))
    }
  })
  const confirm = useConfirm()

  useEffect(() => {
    if (!error) return
    confirm({ title: 'エラー', description: error.message, hideCancelButton: true })
  }, [error])

  return { data, error, isPending }
}
