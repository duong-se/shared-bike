import { createContext, useCallback, useContext, useState } from "react"
import { useCookies } from 'react-cookie'

type AuthContextType = {
  login: (username: string, password: string, callback: VoidFunction) => void;
  logout: (callback: VoidFunction) => void;
  error?: string
}

const AuthContext = createContext<AuthContextType>(null!);

export const useAuth = () => {
  return useContext(AuthContext);
}

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [, setCookie] = useCookies()
  const [error, setError] = useState<string| undefined>(undefined)
  const login = useCallback(async (username: string, password: string, callback: VoidFunction) => {
    const loginUrl = `${window.sharedBike.config.baseUrl}/users/login`
    fetch(loginUrl, {
      method: "POST",
      credentials: "include",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        username,
        password,
      })
    }).then((resp) => {
      resp.status === 204 ? callback() : setError("Login failed")
    }).catch((error) => setError(error))
  }, [])

  const logout = useCallback(() => {
   setCookie('session', null)
  }, [setCookie])


  return <AuthContext.Provider value={{ login, logout, error }}>{children}</AuthContext.Provider>;
}
