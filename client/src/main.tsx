import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import App from './App.tsx'
import { ThemeProvider } from './hooks/ThemeProvier.tsx'

const queryClient = new QueryClient()

createRoot(document.getElementById('root')!).render(

   <ThemeProvider defaultTheme='dark' storageKey='vite-ui-theme'>
      <StrictMode>
         <QueryClientProvider client={queryClient}>
            <App />
         </QueryClientProvider>
      </StrictMode>,
   </ThemeProvider>
)
