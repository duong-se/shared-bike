import { render, screen, waitFor } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { initialize, mockInstances } from '@googlemaps/jest-mocks'
import * as reactQuery from 'react-query'
import * as useAuth from './hooks/AuthProvider'
import * as useMap from './hooks/useMap'
import * as useBikes from './hooks/useBikes'

import App from './App'
import { tokenKey } from './constants/constants'
import { Bike } from './typings/types'

describe('App', () => {
  beforeEach(() => {
    initialize()
  })
  it('renders login page', async () => {
    render(<App />, { wrapper: BrowserRouter})
    await waitFor(() => {
      const linkElement = screen.getByText(/Shared bike platform for everyone/i)
      expect(linkElement).toBeInTheDocument()
    })
  })

  it('renders dashboard page', async () => {
    jest.spyOn(reactQuery, 'useQueryClient').mockReturnValue({
      setQueryData: jest.fn(),
    } as unknown as reactQuery.QueryClient)
    jest.spyOn(useAuth, 'useAuth').mockReturnValue({
      user: {
        id: 1,
        name: 'mockName',
        username: 'mockUsername',
      },
      setUser: jest.fn()
    })
    jest.spyOn(useBikes, 'useBikes').mockReturnValue({
      isLoading: false,
      isSuccess: true,
      isError: false,
      error: '',
      data: null
    } as unknown as reactQuery.UseQueryResult<Bike[], unknown>)
    jest.spyOn(useMap, 'useMap').mockReturnValue({
      map: mockInstances.get(Map)[0] as unknown as google.maps.Map,
      mapRef: jest.fn()
    })
    localStorage.setItem(tokenKey, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsIm5hbWUiOiJUZXN0IFVzZXIiLCJwZXJtaXNzaW9ucyI6bnVsbCwiZXhwIjoxNjU3MzUyODM0fQ.jU5yp2y-3H-dxXP1hdDW-FYEYv5F9GhAVDCbafphUzI')
    render(<App />, { wrapper: BrowserRouter})
    await waitFor(() => {
      const logoutElement = screen.getByText(/Logout/i)
      expect(logoutElement).toBeInTheDocument()
    })
  })
})
