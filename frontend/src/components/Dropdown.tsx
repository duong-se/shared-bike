import { useCallback } from 'react'
import { useNavigate } from 'react-router-dom'
import { tokenKey } from '../constants/constants'
import { useAuth } from '../hooks/AuthProvider'

export const Dropdown: React.FC = () => {
  const { user } = useAuth()
  const navigation = useNavigate()
  const handleLogout = useCallback(() => {
    localStorage.removeItem(tokenKey)
    navigation('/', { replace: true })
  }, [navigation])

  if (!user) {
    return null
  }
  return (
    <div className="absolute z-10 right-1">
      <div className="dropdown dropdown-end">
        <label data-testid="avatar" tabIndex={0} className="btn btn-ghost btn-circle avatar placeholder">
          <div className="bg-neutral-focus text-neutral-content w-10 rounded-full text-xs">
            <span className="text-xs">{user.name[0]}</span>
          </div>
        </label>
        <ul tabIndex={0} className="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52">
          <li><button onClick={handleLogout}>Logout</button></li>
        </ul>
      </div>
    </div>
  )
}
