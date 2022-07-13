import { fireEvent, render, screen, waitFor } from '@testing-library/react'
import { QueryClient, QueryClientProvider } from 'react-query'
import * as reactQuery from 'react-query'
import { initialize, mockInstances, Map, Marker, InfoWindow } from '@googlemaps/jest-mocks'
import {
  availableColor,
  BikeMapPage,
  customIcon,
  rentedColor,
  contentMap,
  renderCommonContent,
  renderUserHasBikeCase,
  renderUserHasNoBikeCase,
  onRentOrReturnSuccess,
  handleMarkerCallback,
  handlePopUpButtonCallback,
} from './BikeMapPage'
import * as useBikes from '../hooks/useBikes'
import { Bike, BikeStatus } from '../typings/types'
import * as useAuth from '../hooks/AuthProvider'
import * as useMap from '../hooks/useMap'

describe('BikeMapPage', () => {
  beforeEach(() => {
    initialize()
  })
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  })
  const Provider: React.FC<React.PropsWithChildren> = ({ children }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  )
  it('should render loading', async () => {
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
      isLoading: true,
      isSuccess: false,
      data: [{
        id: 1,
        name: 'mockName',
        lat: '50.123456',
        long: '8.123456',
        status: BikeStatus.RENTED,
        nameOfRenter: 'mockRenter',
        userId: 1,
      }]
    } as unknown as reactQuery.UseQueryResult<Bike[], unknown>)
    jest.spyOn(useMap, 'useMap').mockReturnValue({
      map: mockInstances.get(Map)[0] as unknown as google.maps.Map,
      mapRef: jest.fn()
    })

    render(<Provider><BikeMapPage /></Provider>)
    await waitFor(() => {
      const spinner = screen.getByRole('status')
      expect(spinner).toBeInTheDocument()
    })
  })

  it('should render map when user have bike', async () => {
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
      data: [{
        id: 1,
        name: 'mockName',
        lat: '50.123456',
        long: '8.123456',
        status: BikeStatus.RENTED,
        nameOfRenter: 'mockRenter',
        userId: 1,
      }]
    } as unknown as reactQuery.UseQueryResult<Bike[], unknown>)
    jest.spyOn(useMap, 'useMap').mockReturnValue({
      map: mockInstances.get(Map)[0] as unknown as google.maps.Map,
      mapRef: jest.fn()
    })

    render(<Provider><BikeMapPage /></Provider>)
    await waitFor(() => {
      const map = screen.getByTestId('map-container')
      expect(map).toBeInTheDocument()
    })
  })

  it('should render map when user have no bike', async () => {
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
      data: [{
        id: 1,
        name: 'mockName',
        lat: '50.123456',
        long: '8.123456',
        status: BikeStatus.AVAILABLE,
        nameOfRenter: '',
        userId: 0,
      }]
    } as unknown as reactQuery.UseQueryResult<Bike[], unknown>)
    jest.spyOn(useMap, 'useMap').mockReturnValue({
      map: mockInstances.get(Map)[0] as unknown as google.maps.Map,
      mapRef: jest.fn()
    })

    render(<Provider><BikeMapPage /></Provider>)
    await waitFor(() => {
      const map = screen.getByTestId('map-container')
      expect(map).toBeInTheDocument()
    })
  })

  it('should render error', async () => {
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
      isError: true,
      error: 'mockError',
      data: null
    } as unknown as reactQuery.UseQueryResult<Bike[], unknown>)
    jest.spyOn(useMap, 'useMap').mockReturnValue({
      map: mockInstances.get(Map)[0] as unknown as google.maps.Map,
      mapRef: jest.fn()
    })

    render(<Provider><BikeMapPage /></Provider>)
    const error = (await screen.findByText('mockError')).textContent
    expect(error).toEqual('mockError')
  })
})

describe('customIcon', () => {
  it('should render custom icon with available color', () => {
    const icon = customIcon(availableColor)
    expect(icon).toMatchSnapshot()
  })
  it('should render custom icon with rented color', () => {
    const icon = customIcon(rentedColor)
    expect(icon).toMatchSnapshot()
  })
})

describe('contentMap', () => {
  it('should render content map for available bike', () => {
    const map = contentMap[BikeStatus.AVAILABLE]
    const view = map.renderButton({
      id: 1,
      name: 'mockName',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    } as Bike)
    expect(view).toContain('bike-action-1')
    expect(view).toContain('RENT BIKE')
  })

  it('should render content map for rented bike', () => {
    const map = contentMap[BikeStatus.RENTED]
    const view = map.renderButton({
      id: 1,
      name: 'mockName',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    } as Bike)
    expect(view).toContain('bike-action-1')
    expect(view).toContain('RENT BIKE')
  })

  it('should render content map for retur bike', () => {
    const map = contentMap.returnBike
    const view = map.renderButton({
      id: 1,
      name: 'mockName',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    } as Bike)
    expect(view).toContain('bike-action-1')
    expect(view).toContain('RETURN BIKE')
  })
})

describe('renderCommonContent', () => {
  it('should renderCommonContent', () => {
    const view = renderCommonContent({
      id: 1,
      name: 'mockName',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    } as Bike)
    expect(view).toContain('mockName')
  })
})

describe('renderUserHasBikeCase', () => {
  beforeEach(() => {
    initialize()
  })
  it('should render bikes', () => {
    const mockReturnBikeMutate = jest.fn()
    const map = new google.maps.Map(document.createElement('div'))
    const view = renderUserHasBikeCase(
      [
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
      ] as Array<Bike>,
      {
        id: 1,
        name: 'mockName',
        lat: '50.123456',
        long: '8.123456',
        status: BikeStatus.RENTED,
        nameOfRenter: 'mockRenter',
        userId: 1,
      },
      mockReturnBikeMutate,
      map,
    )
    const infoWindowMocks = mockInstances.get(InfoWindow)
    const markerMocks = mockInstances.get(Marker)
    expect(markerMocks).toHaveLength(2)
    expect(markerMocks[0].setIcon).toHaveBeenCalledTimes(1)
    expect(infoWindowMocks[0].setContent).toHaveBeenCalledTimes(1)
    expect(infoWindowMocks[0].addListener).toHaveBeenCalledTimes(1)
    expect(markerMocks[0].addListener).toHaveBeenCalledTimes(1)
    expect(markerMocks[1].setIcon).toHaveBeenCalledTimes(1)
    expect(infoWindowMocks[1].setContent).toHaveBeenCalledTimes(1)
    expect(markerMocks[1].addListener).toHaveBeenCalledTimes(1)
    expect(infoWindowMocks[1].addListener).toHaveBeenCalledTimes(0)
    expect(view).toHaveLength(2)
  })

  it('should return market empty', () => {
    const mockReturnBikeMutate = jest.fn()
    const view = renderUserHasBikeCase(
      [
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
      ] as Array<Bike>,
      {
        id: 1,
        name: 'mockName',
        lat: '50.123456',
        long: '8.123456',
        status: BikeStatus.RENTED,
        nameOfRenter: 'mockRenter',
        userId: 1,
      },
      mockReturnBikeMutate,
      undefined,
    )
    expect(view).toHaveLength(0)
  })
})


describe('renderUserHasNoBikeCase', () => {
  beforeEach(() => {
    initialize()
  })
  it('should render bikes', () => {
    const mockReturnBikeMutate = jest.fn()
    const map = new google.maps.Map(document.createElement('div'))
    const view = renderUserHasNoBikeCase(
      [
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
          status: BikeStatus.AVAILABLE,
          nameOfRenter: '',
          userId: 0,
        },
      ] as Array<Bike>,
      mockReturnBikeMutate,
      map,
    )
    const infoWindowMocks = mockInstances.get(InfoWindow)
    const markerMocks = mockInstances.get(Marker)
    expect(markerMocks).toHaveLength(2)
    expect(markerMocks[0].setIcon).toHaveBeenCalledTimes(1)
    expect(infoWindowMocks[0].addListener).toHaveBeenCalledTimes(0)
    expect(markerMocks[0].addListener).toHaveBeenCalledTimes(1)
    expect(markerMocks[1].setIcon).toHaveBeenCalledTimes(1)
    expect(markerMocks[1].addListener).toHaveBeenCalledTimes(1)
    expect(infoWindowMocks[1].addListener).toHaveBeenCalledTimes(1)
    expect(view).toHaveLength(2)
  })

  it('should return empty', () => {
    const mockReturnBikeMutate = jest.fn()
    const view = renderUserHasNoBikeCase(
      [
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
      ] as Array<Bike>,
      mockReturnBikeMutate,
      undefined,
    )
    expect(view).toHaveLength(0)
  })
})

describe('onRentOrReturnSuccess', () => {
  it('should run success', () => {
    const queryClient = new QueryClient()
    const spy = jest.spyOn(queryClient, 'setQueryData').mockReturnValue([
      {
        id: 1,
        name: 'mockName1',
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
      }
    ])
    jest.spyOn(queryClient, 'getQueryData').mockReturnValue([
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
    ])
    const fn = onRentOrReturnSuccess(queryClient)
    fn({
      id: 1,
      name: 'mockName1',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    })
    expect(spy).toBeCalled()
  })

  it('should run for no currentBike', () => {
    const queryClient = new QueryClient()
    const spy = jest.spyOn(queryClient, 'setQueryData').mockReturnValue([
      {
        id: 1,
        name: 'mockName1',
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
      }
    ])
    jest.spyOn(queryClient, 'getQueryData').mockReturnValue(undefined)
    const fn = onRentOrReturnSuccess(queryClient)
    fn({
      id: 1,
      name: 'mockName1',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    })
    expect(spy).not.toBeCalled()
  })
})

describe('handleMarkerCallback', () => {
  beforeEach(() => {
    initialize()
  })
  it('should run and call call open fn', () => {
    const map = new google.maps.Map(document.createElement('div'))
    const infoWindow = new google.maps.InfoWindow()
    const marker = new google.maps.Marker()
    const fn = handleMarkerCallback(infoWindow, marker, map)
    fn()
    const infoWindowMocks = mockInstances.get(InfoWindow)
    expect(infoWindowMocks[0].open).toHaveBeenCalledTimes(1)
  })
})

describe('handlePopUpButtonCallback', () => {
  const MockComponent = () => <button data-testid="test-button-1" id="bike-action-1">mock</button>
  beforeEach(() => {
    initialize()
  })
  it('should run and call open fn', () => {
    const infoWindow = new google.maps.InfoWindow()
    const mockFn = jest.fn()
    const mockBike = {
      id: 1,
      name: 'mockName1',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    }
    render(<MockComponent />)
    const fn = handlePopUpButtonCallback(mockFn, mockBike, infoWindow)
    fn()
    const markerButton = screen.getByTestId('test-button-1')
    fireEvent.click(markerButton)
    const infoWindowMocks = mockInstances.get(InfoWindow)
    expect(mockFn).toBeCalled()
    expect(infoWindowMocks[0].close).toHaveBeenCalledTimes(1)
  })

  it('should run and cannot find button', () => {
    const infoWindow = new google.maps.InfoWindow()
    const mockFn = jest.fn()
    const mockBike = {
      id: 1,
      name: 'mockName1',
      lat: '50.123456',
      long: '8.123456',
      status: BikeStatus.RENTED,
      nameOfRenter: 'mockRenter',
      userId: 1,
    }
    render(<div />)
    const fn = handlePopUpButtonCallback(mockFn, mockBike, infoWindow)
    fn()
  })
})
