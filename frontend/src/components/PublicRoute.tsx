import { Navigate, useLocation } from 'react-router-dom'
import { tokenKey } from '../constants/constants'

export const PublicRoute = ({ children }: { children: JSX.Element }) => {
  const location = useLocation()
  const token = localStorage.getItem(tokenKey)
  if (token) {
    return <Navigate to="/dashboard" replace={true} state={{ from: location }} />
  }
  return children
}
