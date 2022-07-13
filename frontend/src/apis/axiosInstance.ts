import axios, { AxiosRequestConfig, AxiosResponse } from 'axios'
import { commonHeaders, tokenKey } from '../constants/constants'

export enum HTTP_STATUS {
  UNAUTHORIZED = 401,
  OK = 200,
}

export const axiosApiInstance = axios.create({
  headers: commonHeaders,
})

export const handleInterceptRequestError = (error: unknown) => {
  throw error
}

export const handleInterceptConfig =  (config: AxiosRequestConfig<unknown>) => {
  const token = localStorage.getItem(tokenKey)
  if (!token) {
    return config
  }
  config.headers = {
    ...config.headers,
    Authorization: `Bearer ${token}` as string,
  }
  return config
}

axiosApiInstance.interceptors.request.use(
  handleInterceptConfig,
  handleInterceptRequestError
)

export const handleInterceptResponse = (response: AxiosResponse<unknown, unknown>) => {
  return response
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const handleInterceptResponseError = (error: any) => {
  if (error?.response?.status === HTTP_STATUS.UNAUTHORIZED) {
    localStorage.clear()
    window.location.href = '/'
  }
  const errorMessage = error.response?.data ?? error.message
  throw new Error(errorMessage)
}

axiosApiInstance.interceptors.response.use(
  handleInterceptResponse,
  handleInterceptResponseError,
)
