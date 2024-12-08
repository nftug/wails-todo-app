import TheDrawer from '@/components/layout/TheDrawer'
import TheHeader from '@/components/layout/TheHeader'
import { HeaderProvider } from '@/context/layout/HeaderContext'
import AboutPage from '@/pages/AboutPage'
import IndexPage from '@/pages/IndexPage'
import SettingsPage from '@/pages/SettingsPage'
import { Box, createTheme, CssBaseline, ThemeProvider, Toolbar } from '@mui/material'
import { ConfirmProvider } from 'material-ui-confirm'
import { BrowserRouter, Route, Routes } from 'react-router-dom'

const App: React.FC = () => {
  const theme = createTheme({ colorSchemes: { dark: true } })
  const confirmOptions = { confirmationText: 'OK', cancellationText: 'キャンセル' } as const

  return (
    <BrowserRouter>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <ConfirmProvider defaultOptions={confirmOptions}>
          <HeaderProvider>
            <TheHeader />
            <TheDrawer />
          </HeaderProvider>

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
    </BrowserRouter>
  )
}

export default App
