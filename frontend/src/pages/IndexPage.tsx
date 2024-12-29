import TodoAddFab from '@/features/todo/components/TodoAddFab'
import TodoEditModal from '@/features/todo/components/TodoEditModal'
import TodoList from '@/features/todo/components/TodoList'
import { fullViewHeightStyle } from '@/lib/layout/styles'
import { useState } from 'react'

const IndexPage: React.FC = () => {
  const [openEdit, setOpenEdit] = useState(false)
  const [selectedItemId, setSelectedItemId] = useState<number | null>(null)

  const onClickEdit = (itemId: number | null) => {
    setSelectedItemId(itemId)
    setOpenEdit(true)
  }

  return (
    <>
      <TodoList sx={fullViewHeightStyle} onClickEdit={onClickEdit} />
      <TodoAddFab onClickAdd={() => onClickEdit(null)} />

      <TodoEditModal open={openEdit} itemId={selectedItemId} onClose={() => setOpenEdit(false)} />
    </>
  )
}

export default IndexPage
