import TodoItem from '@/features/todo/components/TodoItem'
import { useTodoQuery } from '@/features/todo/hooks/useTodoQuery'
import { Box, List, ListItem, SxProps, Theme, Typography } from '@mui/material'
import { todo } from '@wailsjs/go/models'

interface Props {
  query?: todo.Query
  onClickEdit: (itemId: number) => void
  sx?: SxProps<Theme>
}

const TodoList: React.FC<Props> = ({ query, onClickEdit, sx }) => {
  const { data } = useTodoQuery(query ?? {})

  return (
    <>
      <List sx={sx}>
        {!data?.length ? (
          <Box display="flex" justifyContent="center" alignItems="center" height={1}>
            <Typography variant="h6" color="textSecondary">
              Todoがありません。
            </Typography>
          </Box>
        ) : (
          data.map((item) => (
            <ListItem key={item.id}>
              <TodoItem item={item} onClickEdit={onClickEdit} />
            </ListItem>
          ))
        )}
      </List>
    </>
  )
}

export default TodoList
