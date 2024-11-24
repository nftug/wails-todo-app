import { Home, Info } from '@mui/icons-material'
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText
} from '@mui/material'
import { useAtom } from 'jotai'
import { Link } from 'react-router-dom'
import { drawerOpenAtom } from '../../atoms/layout-atoms'

type DrawerItem = {
  name: string
  href: string
  icon?: React.JSX.Element
}

const TheDrawer: React.FC = () => {
  const [drawerOpen, setDrawerOpen] = useAtom(drawerOpenAtom)

  const menuItems: DrawerItem[] = [
    {
      name: 'ホーム',
      href: '/',
      icon: <Home />
    },
    {
      name: 'このアプリについて',
      href: '/about',
      icon: <Info />
    }
  ]

  return (
    <Drawer open={drawerOpen} onClose={() => setDrawerOpen(false)}>
      <Box sx={{ width: 250 }} role="presentation" onClick={() => setDrawerOpen(false)}>
        <List>
          {menuItems.map((item) => (
            <ListItem key={item.name} disablePadding>
              <ListItemButton component={Link} to={item.href}>
                <ListItemIcon children={item.icon} />
                <ListItemText primary={item.name} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Box>
    </Drawer>
  )
}

export default TheDrawer
