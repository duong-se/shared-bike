import { lazy, Suspense } from 'react'
import {
  Routes,
  Route,
} from 'react-router-dom'
import { PrivateRoute } from './components/PrivateRoute'
import { PublicRoute } from './components/PublicRoute'
import { AuthProvider } from './hooks/AuthProvider'
import {
  QueryClient,
  QueryClientProvider,
} from 'react-query'
import { Spinner } from './components/Spinner'

const LoginPage = lazy(() => import('./pages/LoginPage'))
const BikeMapPage = lazy(() => import('./pages/BikeMapPage'))

const App = () => {
  const queryClient = new QueryClient()
  return (
    <Suspense fallback={<Spinner />}>
      <QueryClientProvider client={queryClient} >
        <AuthProvider>
          <Routes>
            <Route path="/" element={<PublicRoute><LoginPage /></PublicRoute>} />
            <Route path="/dashboard" element={<PrivateRoute><BikeMapPage /></PrivateRoute>} />
          </Routes>
        </AuthProvider>
      </QueryClientProvider>
    </Suspense>
  )
}

export default App
