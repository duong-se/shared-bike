import { Navigate, useLocation } from 'react-router-dom'
import { tokenKey } from '../hooks/AuthProvider';

export const PublicRoute = ({ children }: { children: JSX.Element }) => {
  let location = useLocation();
  const token = localStorage.getItem(tokenKey)
  if (token) {
    return <Navigate to="/dashboard" state={{ from: location }} replace />;
  }
  return children;
}
