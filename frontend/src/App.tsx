import { Box, createTheme, CssBaseline, ThemeProvider, Toolbar } from '@mui/material'
import { Provider } from 'inversify-react'
import { ConfirmProvider } from 'material-ui-confirm'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import TheDrawer from './components/layout/TheDrawer'
import TheHeader from './components/layout/TheHeader'
import { container } from './inversify.config'
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

      <Provider container={container}>
        <ConfirmProvider defaultOptions={confirmOptions}>
          <BrowserRouter>
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
          </BrowserRouter>
        </ConfirmProvider>
      </Provider>
    </ThemeProvider>
  )
}

export default App
