import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { tokenKey } from '../constants/constants'
import { AuthProvider } from '../hooks/AuthProvider'
import { PrivateRoute } from './PrivateRoute'

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  Navigate: ({ children }: React.PropsWithChildren) => <div>{children}</div>
}))

describe('PrivateRoute', () => {
  it('should run and return children', () => {
    localStorage.setItem(tokenKey, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsIm5hbWUiOiJUZXN0IFVzZXIiLCJwZXJtaXNzaW9ucyI6bnVsbCwiZXhwIjoxNjU3MzUyODM0fQ.jU5yp2y-3H-dxXP1hdDW-FYEYv5F9GhAVDCbafphUzI')
    render(<AuthProvider><PrivateRoute><div data-testid="mockChildren">Test</div></PrivateRoute></AuthProvider>, { wrapper: BrowserRouter })
    const child = screen.getByTestId('mockChildren')
    expect(child).toBeInTheDocument()
  })

  it('should run and navigate', async () => {
    localStorage.clear()
    const { container } = render(<AuthProvider><PrivateRoute><div data-testid="mockChildren">Test</div></PrivateRoute></AuthProvider>, { wrapper: BrowserRouter })
    expect(container).toMatchSnapshot()
  })
})
