import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import { RegisterForm, onSuccessHandler } from './RegisterForm'
import * as AuthProvider from '../hooks/AuthProvider'
import { UseMutationResult } from 'react-query'
import { RegisterResponse, RegisterVariables } from '../typings/types'
import * as useUsers from '../hooks/useUsers'

describe('RegisterForm', () => {
  it('should run corect', async () => {
    const mockSetUser = jest.fn()
    jest.spyOn(AuthProvider, 'useAuth').mockReturnValue({
      setUser: mockSetUser
    })
    const mockMutate = jest.fn()
    const mockOnSuccess = jest.fn()
    jest.spyOn(useUsers, 'useRegister').mockReturnValue({
      isLoading: true,
      mutate: mockMutate,
      error: '',
      isError: false,
      onSuccess: mockOnSuccess
    } as unknown as UseMutationResult<RegisterResponse, unknown, RegisterVariables, unknown>)
    const mockProps = {
      onClickLogin: jest.fn(),
    }
    render(<RegisterForm {...mockProps} />, { wrapper: BrowserRouter })
    const usernameInput = screen.getByLabelText('Username')
    fireEvent.change(usernameInput, { target: { value: 'mockUsername' } })
    const passwordInput = screen.getByLabelText('Password')
    const confirmPasswordInput = screen.getByLabelText('Confirm Password')
    const nameInput = screen.getByLabelText('Name')
    fireEvent.change(passwordInput, { target: { value: 'mockPassword' } })
    fireEvent.change(confirmPasswordInput, { target: { value: 'mockPassword1' } })
    fireEvent.change(confirmPasswordInput, { target: { value: 'mockPassword' } })
    fireEvent.change(nameInput, { target: { value: 'mockName' } })
    const buttons = screen.getAllByRole('button')
    fireEvent.click(buttons[1])
    await waitFor(() => {
      expect(mockProps.onClickLogin).toBeCalled()
    })
    fireEvent.click(buttons[0])
    await waitFor(() => {
      expect(mockMutate).toBeCalledWith({
        username: 'mockUsername',
        password: 'mockPassword',
        name: 'mockName'
      })
    })
  })

  it('should show error', async () => {
    const mockSetUser = jest.fn()
    jest.spyOn(AuthProvider, 'useAuth').mockReturnValue({
      setUser: mockSetUser
    })
    const mockMutate = jest.fn()
    const mockOnSuccess = jest.fn()
    jest.spyOn(useUsers, 'useRegister').mockReturnValue({
      isLoading: true,
      mutate: mockMutate,
      error: '',
      isError: false,
      onSuccess: mockOnSuccess
    } as unknown as UseMutationResult<RegisterResponse, unknown, RegisterVariables, unknown>)
    const mockProps = {
      onClickLogin: jest.fn(),
    }
    render(<RegisterForm {...mockProps} />, { wrapper: BrowserRouter })
    const buttons = screen.getAllByRole('button')
    fireEvent.click(buttons[0])
    waitFor(() => {
      const usernameError = screen.getByText('Username is required')
      expect(usernameError).toBeInTheDocument()
      const passwordError = screen.getByText('Password is required')
      expect(passwordError).toBeInTheDocument()
      const confirmPasswordError = screen.getByText('Confirm password is required')
      expect(confirmPasswordError).toBeInTheDocument()
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
