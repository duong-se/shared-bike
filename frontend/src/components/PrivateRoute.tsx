import { Navigate, useLocation } from 'react-router-dom'
import { tokenKey } from '../constants/constants'
import { Dropdown } from './Dropdown'

export const PrivateRoute = ({ children }: { children: JSX.Element }) => {
  const token = localStorage.getItem(tokenKey)
  const location = useLocation()
  if (!token) {
    return <Navigate to="/" replace={true} state={{ from: location }} />
  }
  return <div>
    <Dropdown />
    {children}
  </div>
}
