import axios, { AxiosError, AxiosResponse } from 'axios'
import { tokenKey } from '../constants/constants'
import { handleInterceptConfig, handleInterceptRequestError, handleInterceptResponse, handleInterceptResponseError, HTTP_STATUS } from './axiosInstance'

describe('handleInterceptRequestError', () => {
  it('should return correct error', () => {
    let err
    jest.spyOn(axios, 'create')
    try {
      handleInterceptRequestError(new Error('mockError'))
    } catch (error) {
      err = error
    }
    expect(err).toBeInstanceOf(Error)
    expect((err as unknown as Error).message).toEqual('mockError')
  })
})

describe('handleInterceptConfig', () => {
  it('should return config', () => {
  jest.spyOn(axios, 'create')
  const result = handleInterceptConfig({})
    expect(result).toEqual({})
  })

  it('should return config with authentication', () => {
  jest.spyOn(axios, 'create')
  localStorage.setItem(tokenKey, 'mockToken')
  const result = handleInterceptConfig({})
    expect(result).toEqual({
      headers: {
        Authorization: 'Bearer mockToken'
      }
    })
  })
})

describe('handleInterceptResponse', () => {
  it('should run and return response', () => {
    const result = handleInterceptResponse({ data: 'mockData' } as AxiosResponse<any, any>)
    expect(result).toEqual({ data: 'mockData' })
  })
})

describe('handleInterceptResponseError', () => {
  it('should run and return normal error', () => {
    const mockError = new Error('mockError')
    let err
    try {
      handleInterceptResponseError(mockError)
    } catch (error) {
      err = error
    }
    expect((err as Error).message).toEqual('mockError')
  })

  it('should run and return response', () => {
    const mockError = new AxiosError(
      'mockError',
      HTTP_STATUS.UNAUTHORIZED as unknown as string,
      {},
      {},
      { status: HTTP_STATUS.UNAUTHORIZED as unknown as number, statusText: 'Unauthorized', data: 'mockError' } as AxiosResponse)
    let err
    localStorage.setItem(tokenKey, 'mockToken')
    try {
      handleInterceptResponseError(mockError)
    } catch (error) {
      err = error
    }
    expect((err as Error).message).toEqual('mockError')
    expect(localStorage.getItem(tokenKey)).toBeNull()
  })

  it('should run and not clear localStorage', () => {
    const mockError = new AxiosError(
      'mockError',
      HTTP_STATUS.OK as unknown as string,
      {},
      {},
      { status: HTTP_STATUS.OK as unknown as number, statusText: 'Unauthorized', data: 'mockError' } as AxiosResponse)
    let err
    localStorage.setItem(tokenKey, 'mockToken')
    try {
      handleInterceptResponseError(mockError)
    } catch (error) {
      err = error
    }
    expect((err as Error).message).toEqual('mockError')
    expect(localStorage.getItem(tokenKey)).toEqual('mockToken')
  })
})
