import TodoEditModal from '@/components/todo/TodoEditModal'
import { createContext, useState } from 'react'
import sleep from 'sleep-promise'

interface TodoEditModalContextType {
  openModal: (itemId?: string) => void
}

export const TodoEditModalContext = createContext<TodoEditModalContextType>(undefined!)

export const TodoEditModalProvider = ({ children }: { children?: React.ReactNode }) => {
  const [open, setOpen] = useState(false)
  const [itemId, setItemId] = useState<string | null>(null)

  const openModal = (id?: string) => {
    setItemId(id ?? null) // IDを更新して、データ読み込み開始
    setOpen(true)
  }

  const closeModal = async () => {
    setOpen(false)
    await sleep(100) // ダイアログが完全に閉じるまで待機
    setItemId(null)
  }

  return (
    <TodoEditModalContext.Provider value={{ openModal }}>
      {children}
      <TodoEditModal open={open} itemId={itemId} onClose={closeModal} />
    </TodoEditModalContext.Provider>
  )
}
