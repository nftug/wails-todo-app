import { Provider } from 'inversify-react'
import React from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import App from './App'
import { container } from './inversify.config'

const root = createRoot(document.getElementById('root')!)

root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Provider container={container}>
        <App />
      </Provider>
    </BrowserRouter>
  </React.StrictMode>
)
