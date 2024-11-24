import { Box, Toolbar } from '@mui/material'
import { Provider } from 'inversify-react'
import { ConfirmProvider } from 'material-ui-confirm'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import TheDrawer from './components/layout/TheDrawer'
import TheHeader from './components/layout/TheHeader'
import { container } from './inversify.config'
import AboutPage from './pages/AboutPage'
import IndexPage from './pages/IndexPage'

const App: React.FC = () => {
  return (
    <Provider container={container}>
      <ConfirmProvider defaultOptions={{ confirmationText: 'OK', cancellationText: 'キャンセル' }}>
        <BrowserRouter>
          <TheHeader />
          <TheDrawer />

          <Box component="main">
            <Toolbar />

            <Routes>
              <Route index element={<IndexPage />} />
              <Route path="/about" element={<AboutPage />} />
            </Routes>
          </Box>
        </BrowserRouter>
      </ConfirmProvider>
    </Provider>
  )
}

export default App
