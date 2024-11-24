import { createContext, useContext, useState } from 'react'
import TodoEditModal from './TodoEditModal'

interface TodoEditModalContextType {
  openModal: (itemId?: string) => void
}

const TodoEditModalContext = createContext<TodoEditModalContextType>(undefined!)

export const useTodoEditModal = () => useContext(TodoEditModalContext)

export const TodoEditModalProvider = ({ children }: { children?: React.ReactNode }) => {
  const [open, setOpen] = useState(false)
  const [itemId, setItemId] = useState<string | null>(null)

  const openModal = (id?: string) => {
    setItemId(id ?? null)
    setOpen(true)
  }

  const closeModal = () => {
    setOpen(false)
    // setItemId(null)
  }

  return (
    <TodoEditModalContext.Provider value={{ openModal }}>
      {children}
      <TodoEditModal open={open} itemId={itemId} onClose={closeModal} />
    </TodoEditModalContext.Provider>
  )
}
