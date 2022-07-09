import { Navigate, useLocation } from 'react-router-dom'
import { useCookies } from 'react-cookie'

export const PrivateRoute = ({ children }: { children: JSX.Element }) => {
  const [cookies] = useCookies<"session">();
  console.log({ cookies })
  let location = useLocation();
  if (!cookies.session) {
    return <Navigate to="/" state={{ from: location }} replace />;
  }
  return children;
}
