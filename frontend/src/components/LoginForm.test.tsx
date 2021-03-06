import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { LoginForm, onSuccessHandler } from './LoginForm'
import * as AuthProvider from '../hooks/AuthProvider'
import { UseMutationResult } from 'react-query'
import { LoginResponse, LoginVariables } from '../typings/types'
import * as useUsers from '../hooks/useUsers'

describe('LoginForm', () => {
  it('should run corect', async () => {
    const mockSetUser = jest.fn()
    jest.spyOn(AuthProvider, 'useAuth').mockReturnValue({
      setUser: mockSetUser
    })
    const mockMutate = jest.fn()
    const mockOnSuccess = jest.fn()
    jest.spyOn(useUsers, 'useLogin').mockReturnValue({
      isLoading: false,
      mutate: mockMutate,
      error: '',
      isError: false,
      onSuccess: mockOnSuccess
    } as unknown as UseMutationResult<LoginResponse, unknown, LoginVariables, unknown>)
    const mockProps = {
      onClickRegister: jest.fn(),
    }
    render(<LoginForm {...mockProps} />, { wrapper: BrowserRouter })
    const usernameInput = screen.getByLabelText('Username')
    fireEvent.change(usernameInput, { target: { value: 'mockUsername' } })
    const passwordInput = screen.getByLabelText('Password')
    fireEvent.change(passwordInput, { target: { value: 'mockPassword' } })
    const buttons = screen.getAllByRole('button')
    fireEvent.click(buttons[1])
    await waitFor(() => {
      expect(mockProps.onClickRegister).toBeCalled()
    })
    fireEvent.click(buttons[0])
    await waitFor(() => {
      expect(mockMutate).toBeCalledWith({
        username: 'mockUsername',
        password: 'mockPassword'
      })
    })
  })

  it('should show form errors', async () => {
    const mockSetUser = jest.fn()
    jest.spyOn(AuthProvider, 'useAuth').mockReturnValue({
      setUser: mockSetUser
    })
    const mockMutate = jest.fn()
    const mockOnSuccess = jest.fn()
    jest.spyOn(useUsers, 'useLogin').mockReturnValue({
      isLoading: false,
      mutate: mockMutate,
      error: 'Network Error',
      isError: true,
      onSuccess: mockOnSuccess
    } as unknown as UseMutationResult<LoginResponse, unknown, LoginVariables, unknown>)
    const mockProps = {
      onClickRegister: jest.fn(),
    }
    render(<LoginForm {...mockProps} />, { wrapper: BrowserRouter })
    const buttons = screen.getAllByRole('button')
    fireEvent.click(buttons[0])
    waitFor(() => {
      const networkError = screen.getByText('Network Error')
      expect(networkError).toBeInTheDocument()
    })
  })

  it('should show error after submit', async () => {
    const mockSetUser = jest.fn()
    jest.spyOn(AuthProvider, 'useAuth').mockReturnValue({
      setUser: mockSetUser
    })
    const mockMutate = jest.fn()
    const mockOnSuccess = jest.fn()
    jest.spyOn(useUsers, 'useLogin').mockReturnValue({
      isLoading: false,
      mutate: mockMutate,
      error: '',
      isError: false,
      onSuccess: mockOnSuccess
    } as unknown as UseMutationResult<LoginResponse, unknown, LoginVariables, unknown>)
    const mockProps = {
      onClickRegister: jest.fn(),
    }
    render(<LoginForm {...mockProps} />, { wrapper: BrowserRouter })
    const buttons = screen.getAllByRole('button')
    fireEvent.click(buttons[0])
    waitFor(() => {
      const usernameError = screen.getByText('Username is required')
      expect(usernameError).toBeInTheDocument()
      const passwordError = screen.getByText('Password is required')
      expect(passwordError).toBeInTheDocument()
      const nameError = screen.getByText('Name is required')
      expect(nameError).toBeInTheDocument()
    })
  })

  it('onSuccessHandler should call navigate and setUser', () => {
    const mockSetUser = jest.fn()
    const mockNavigate = jest.fn()
    const fn = onSuccessHandler(mockSetUser, mockNavigate)
    fn({ accessToken: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsIm5hbWUiOiJUZXN0IFVzZXIiLCJwZXJtaXNzaW9ucyI6bnVsbCwiZXhwIjoxNjU3MzUyODM0fQ.jU5yp2y-3H-dxXP1hdDW-FYEYv5F9GhAVDCbafphUzI' })
    expect(mockSetUser).toBeCalledWith({
      'exp': 1657352834,
      'id': 1,
      'name': 'Test User',
      'permissions': null,
      'username': 'test1',
    })
    expect(mockNavigate).toBeCalledWith('/dashboard', { replace: true })
  })
})
