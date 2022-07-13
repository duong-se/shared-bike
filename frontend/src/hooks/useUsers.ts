import { useMutation, UseMutationOptions } from 'react-query'
import { login, register } from '../apis/users'
import { LoginResponse, LoginVariables } from '../typings/types'
import { RegisterResponse, RegisterVariables } from '../typings/types'

export const useLogin = (options?: Omit<UseMutationOptions<LoginResponse, unknown, LoginVariables, unknown>, 'mutationFn'>) => {
  return useMutation(login, options)
}

export const useRegister = (options?: Omit<UseMutationOptions<RegisterResponse, unknown, RegisterVariables, unknown>, 'mutationFn'>) => {
  return useMutation(register, options)
}
