import { Navigate, useLocation } from 'react-router-dom'
import { useCookies } from 'react-cookie'

export const PublicRoute = ({ children }: { children: JSX.Element }) => {
  const [cookies] = useCookies<"session">();
  console.log({ cookies })
  let location = useLocation();
  if (cookies.session) {
    return <Navigate to="/dashboard" state={{ from: location }} replace />;
  }
  return children;
}
