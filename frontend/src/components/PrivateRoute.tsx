import { Navigate, useLocation } from 'react-router-dom'
import { tokenKey } from '../constants/constants';

export const PrivateRoute = ({ children }: { children: JSX.Element }) => {
  const token = localStorage.getItem(tokenKey)
  let location = useLocation()
  if (!token) {
    return <Navigate to="/" replace={true} state={{ from: location }} />
  }
  return children;
}
