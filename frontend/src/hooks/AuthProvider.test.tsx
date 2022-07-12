import { render, screen, waitFor } from '@testing-library/react'
import { tokenKey } from '../constants/constants'
import { AuthProvider, useAuth } from './AuthProvider'

describe('AuthProvider', () => {
  it('should run correctly and show user', async () => {
    const MockComponent = () => {
      const { user } = useAuth()
      return <div>{user?.name}</div>
    }
    localStorage.setItem(tokenKey, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsIm5hbWUiOiJUZXN0IFVzZXIiLCJwZXJtaXNzaW9ucyI6bnVsbCwiZXhwIjoxNjU3MzUyODM0fQ.jU5yp2y-3H-dxXP1hdDW-FYEYv5F9GhAVDCbafphUzI')
    render(<AuthProvider><MockComponent/></AuthProvider>)
    const user = await screen.findByText('Test User')
    await waitFor(() => {
      expect(user).toBeInTheDocument()
    })
  })

  it('should not show user', async () => {
    const MockComponent = () => {
      const { user } = useAuth()
      return <div>{user?.name}</div>
    }
    localStorage.clear()
    const { container } = render(<AuthProvider><MockComponent/></AuthProvider>)
    expect(container).toMatchSnapshot()
  })
})
