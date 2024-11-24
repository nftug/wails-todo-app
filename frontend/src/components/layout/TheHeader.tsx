import MenuIcon from '@mui/icons-material/Menu'
import { AppBar, IconButton, Toolbar, Typography } from '@mui/material'
import { useAtom } from 'jotai'
import { drawerOpenAtom } from '../../atoms/layout-atoms'

const TheHeader: React.FC = () => {
  const [, setDrawerOpen] = useAtom(drawerOpenAtom)

  const toggleDrawer = () => {
    setDrawerOpen((x) => !x)
  }

  return (
    <AppBar>
      <Toolbar>
        <IconButton
          size="large"
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{ mr: 2 }}
          onClick={toggleDrawer}
        >
          <MenuIcon />
        </IconButton>
        <Typography variant="h5">Todo App</Typography>
      </Toolbar>
    </AppBar>
  )
}

export default TheHeader
