import { Box, createTheme, CssBaseline, ThemeProvider, Toolbar } from '@mui/material'
import { ConfirmProvider } from 'material-ui-confirm'
import { Route, Routes } from 'react-router-dom'
import TheDrawer from './components/layout/TheDrawer'
import TheHeader from './components/layout/TheHeader'
import AboutPage from './pages/AboutPage'
import IndexPage from './pages/IndexPage'
import SettingsPage from './pages/SettingsPage'

const App: React.FC = () => {
  const theme = createTheme({
    colorSchemes: { dark: true }
  })
  const confirmOptions = { confirmationText: 'OK', cancellationText: 'キャンセル' } as const

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />

      <ConfirmProvider defaultOptions={confirmOptions}>
        <TheHeader />
        <TheDrawer />

        <Box component="main">
          <Toolbar />
          <Routes>
            <Route index element={<IndexPage />} />
            <Route path="/about" element={<AboutPage />} />
            <Route path="/settings" element={<SettingsPage />} />
          </Routes>
        </Box>
      </ConfirmProvider>
    </ThemeProvider>
  )
}

export default App
