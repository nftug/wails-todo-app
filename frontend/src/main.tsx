import App from '@/App'
import { container } from '@/inversify.config'
import { Provider } from 'inversify-react'
import React from 'react'
import { createRoot } from 'react-dom/client'

const root = createRoot(document.getElementById('root')!)

root.render(
  <React.StrictMode>
    <Provider container={container}>
      <App />
    </Provider>
  </React.StrictMode>
)
