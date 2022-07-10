import { createContext, useCallback, useContext, useEffect, useState } from "react"
import jwtDecode from "jwt-decode";

type User = {
  id: number
  name: string
  username: string
}

type AuthContextType = {
  login: (username: string, password: string, callback: VoidFunction) => void;
  logout: (callback: VoidFunction) => void;
  error?: string
  isLoading: boolean
  user?: User
}

const AuthContext = createContext<AuthContextType>(null!);

export const useAuth = () => {
  return useContext(AuthContext);
}

export const tokenKey = "accessToken";

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [error, setError] = useState<string| undefined>(undefined)
  const [isLoading, setIsLoading] = useState(false)
  const [user, setUser] = useState<User|undefined>()
  const token = localStorage.getItem(tokenKey);
  const decodeToken = useCallback((value: string | null) => {
    if (value) {
      const decoded = jwtDecode<User>(value);
      setUser(decoded);
    }
  }, [])
  useEffect(() => {
    decodeToken(token)
  }, [decodeToken, token])
  const login = useCallback(async (username: string, password: string, callback: VoidFunction) => {
    setIsLoading(true)
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
    }).then(async (resp) => {
      const data = await resp.json()
      if(resp.ok) {
        return data
      }
      throw new Error(data)
    }).then(async (result) => {
      const { accessToken } = result
      localStorage.setItem(tokenKey, accessToken)
      decodeToken(accessToken)
      callback()
    }).catch((error) => {
      setError(error.message)
    }).finally(() => setIsLoading(false))
  }, [decodeToken])

  const logout = useCallback(() => {
    localStorage.removeItem(tokenKey)
  }, [])


  return <AuthContext.Provider value={{ login, logout, error, isLoading, user }}>{children}</AuthContext.Provider>;
}
