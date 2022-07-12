import React, { useEffect } from "react"
import jwtDecode from 'jwt-decode'
import { tokenKey } from "../constants/constants";
import { User } from "../typings/types";


type AuthContextType = {
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>
  user?: User
}

const AuthContext = React.createContext<AuthContextType>(null!);

export const useAuth = () => {
  return React.useContext(AuthContext);
}

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = React.useState<User|undefined>()
  useEffect(() => {
    const token = localStorage.getItem(tokenKey)
    if (token) {
      const decoded = jwtDecode<User>(token);
      setUser(decoded);
    }
  }, [])
  return <AuthContext.Provider value={{ user, setUser }}>{children}</AuthContext.Provider>;
}
