import { useMutation, useQueryClient } from '@tanstack/react-query'
import { DeleteTodo } from '@wailsjs/go/app/TodoApp'
import { useRef } from 'react'
import { handleApiError } from '../../../lib/api/errors'

export const useDeleteTodo = () => {
  const queryClient = useQueryClient()
  const idRef = useRef<number | null>(null)

  const mutation = useMutation({
    mutationFn: async (id: number) => {
      idRef.current = id
      return await handleApiError(async () => await DeleteTodo(id))
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ['todo'],
        predicate: ({ queryKey }) => queryKey[1] === 'list' || queryKey.at(-1) !== idRef.current
      })
    }
  })

  return mutation
}
