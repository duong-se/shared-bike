import { fireEvent, render, screen, waitFor } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { tokenKey } from '../constants/constants'
import { AuthProvider } from '../hooks/AuthProvider'
import { Dropdown } from './Dropdown'

describe('Dropdown', () => {
  it('should render correctly and clear token when logout', () => {
    localStorage.setItem(tokenKey, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsIm5hbWUiOiJUZXN0IFVzZXIiLCJwZXJtaXNzaW9ucyI6bnVsbCwiZXhwIjoxNjU3MzUyODM0fQ.jU5yp2y-3H-dxXP1hdDW-FYEYv5F9GhAVDCbafphUzI')
    const token = localStorage.getItem(tokenKey)
    expect(token).not.toBeNull()
    render(<AuthProvider><Dropdown /></AuthProvider>, { wrapper: BrowserRouter })
    const avatarButton = screen.getByTestId('avatar')
    expect(avatarButton).toBeInTheDocument()
    fireEvent.click(avatarButton)
    const dropdown = screen.getByRole('button')
    fireEvent.click(dropdown)
    waitFor(() => {
      const clearedToken = localStorage.getItem(tokenKey)
      expect(clearedToken).toBeNull()
    })
  })

  it('should render correctly null', () => {
    localStorage.clear()
    const token = localStorage.getItem(tokenKey)
    expect(token).toBeNull()
    const { container } = render(<AuthProvider><Dropdown /></AuthProvider>, { wrapper: BrowserRouter })
    expect(container).toMatchSnapshot()
  })
})
