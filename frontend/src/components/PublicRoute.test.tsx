import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { tokenKey } from '../constants/constants'
import { PublicRoute } from './PublicRoute'

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  Navigate: ({ children }: React.PropsWithChildren) => <div>{children}</div>
}));

describe('PublicRoute', () => {
  it('should run and return children', async () => {
    localStorage.setItem(tokenKey, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsIm5hbWUiOiJUZXN0IFVzZXIiLCJwZXJtaXNzaW9ucyI6bnVsbCwiZXhwIjoxNjU3MzUyODM0fQ.jU5yp2y-3H-dxXP1hdDW-FYEYv5F9GhAVDCbafphUzI')
    const { container } = render(<PublicRoute><div data-testid="mockChildren">Test</div></PublicRoute>, { wrapper: BrowserRouter })
    expect(container).toMatchSnapshot()
  })

  it('should run and navigate', async () => {
    localStorage.clear()
    render(<PublicRoute><div data-testid="mockChildren">Test</div></PublicRoute>, { wrapper: BrowserRouter })
    const child = screen.getByTestId('mockChildren')
    expect(child).toBeInTheDocument()
  })
})
