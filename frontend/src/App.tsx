import { lazy } from "react"
import {
  Routes,
  Route,
} from "react-router-dom"
import { CookiesProvider } from "react-cookie"
import { PrivateRoute } from "./components/PrivateRoute"
import { PublicRoute } from "./components/PublicRoute"
import { AuthProvider } from "./hooks/AuthProvider"

const LoginPage = lazy(() => import("./pages/LoginPage"))
const BikeMapPage = lazy(() => import("./pages/BikeMapPage"))

const App = () => {
  return (
    <CookiesProvider >
      <AuthProvider>
        <Routes>
          <Route path="/" element={<PublicRoute><LoginPage /></PublicRoute>} />
          <Route path="/dashboard" element={<PrivateRoute><BikeMapPage /></PrivateRoute>} />
        </Routes>
      </AuthProvider>
    </CookiesProvider>
  )
}

export default App
