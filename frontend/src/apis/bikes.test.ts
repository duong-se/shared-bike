import { AxiosError } from 'axios'
import { BikeStatus } from '../typings/types'
import { axiosApiInstance, HTTP_STATUS } from './axiosInstance'
import { fetchBikes, rentBike, returnBike } from './bikes'
describe('fetchBikes', () => {
  it('should return all bikes', async () => {
    const mockBikes = [
      {
        id: 1,
        name: 'mockName',
        lat: '50.123456',
        long: '8.123456',
        status: BikeStatus.RENTED,
        nameOfRenter: 'mockRenter',
        userId: 1,
      },
      {
        id: 2,
        name: 'mockName1',
        lat: '50.123456',
        long: '8.123456',
        status: BikeStatus.RENTED,
        nameOfRenter: 'mockRenter1',
        userId: 1,
      },
    ]
    jest.spyOn(axiosApiInstance, 'get').mockResolvedValue({
      data: mockBikes,
      status: HTTP_STATUS.OK
    })
    const result = await fetchBikes()
    expect(result).toEqual(mockBikes)
  })
  it('should throw error normal', async () => {
    jest.spyOn(axiosApiInstance, 'get').mockRejectedValue(new Error('mockError'))
    let err
    try {
      await fetchBikes()
    } catch (error) {
      err = error
    }
    expect(err).toEqual('mockError')
  })

  it('should throw error axios', async () => {
    jest.spyOn(axiosApiInstance, 'get').mockRejectedValue(new AxiosError('mockError'))
    let err
    try {
      await fetchBikes()
    } catch (error) {
      err = error
    }
    expect(err).toEqual('mockError')
  })
})

describe('rentBike', () => {
  it('should rent a bikes', async () => {
    const mockBike = {
      id: 1,
      name: 'mockName',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    }
    jest.spyOn(axiosApiInstance, 'patch').mockResolvedValue({
      data: mockBike,
      status: HTTP_STATUS.OK
    })
    const result = await rentBike({ bikeId: 1 })
    expect(result).toEqual(mockBike)
  })
  it('should throw error normal', async () => {
    jest.spyOn(axiosApiInstance, 'patch').mockRejectedValue(new Error('mockError'))
    let err
    try {
      await rentBike({ bikeId: 1 })
    } catch (error) {
      err = error
    }
    expect(err).toEqual('mockError')
  })
})

describe('returnBike', () => {
  it('should return a bikes', async () => {
    const mockBike = {
      id: 1,
      name: 'mockName',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.AVAILABLE,
      nameOfRenter: '',
      userId: 0,
    }
    jest.spyOn(axiosApiInstance, 'patch').mockResolvedValue({
      data: mockBike,
      status: HTTP_STATUS.OK
    })
    const result = await returnBike({ bikeId: 1 })
    expect(result).toEqual(mockBike)
  })
  it('should throw error normal', async () => {
    jest.spyOn(axiosApiInstance, 'patch').mockRejectedValue(new Error('mockError'))
    let err
    try {
      await returnBike({ bikeId: 1 })
    } catch (error) {
      err = error
    }
    expect(err).toEqual('mockError')
  })
})
