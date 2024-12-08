import { DrawerDispatchContext } from '@/context/layout/HeaderContext'
import MenuIcon from '@mui/icons-material/Menu'
import { AppBar, IconButton, Toolbar, Typography } from '@mui/material'
import { useContext } from 'react'

const TheHeader: React.FC = () => {
  const setDrawerOpened = useContext(DrawerDispatchContext)

  const toggleDrawer = () => {
    setDrawerOpened((x) => !x)
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
