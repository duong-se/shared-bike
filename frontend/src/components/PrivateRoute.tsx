import { Navigate, useLocation } from 'react-router-dom'
import { tokenKey } from '../hooks/AuthProvider'

export const PrivateRoute = ({ children }: { children: JSX.Element }) => {
  const token = localStorage.getItem(tokenKey)
  let location = useLocation();
  if (!token) {
    return <Navigate to="/" state={{ from: location }} replace />;
  }
  return children;
}
