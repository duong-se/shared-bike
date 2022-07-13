import { renderHook, waitFor, act } from '@testing-library/react'
import { initialize } from '@googlemaps/jest-mocks'
import { useMap } from './useMap'

describe('useMap', () => {
  beforeEach(() => {
    initialize()
  })
  it('should run and return map', async () => {
    const centerPosition = new google.maps.LatLng(50.119504, 8.638137)
    const { result } = renderHook(() => useMap(centerPosition))
    await waitFor(() => result.current.map)
    expect(result.current.map).toEqual(undefined)
    expect(result.current.mapRef).toEqual(expect.any(Function))
    act(() => {
      result.current.mapRef(document.createElement('div'))
    })
    await waitFor(() => {
      expect(result.current.map).toBeInstanceOf(google.maps.Map)
    })
  })

  it('should run and not return map', async () => {
    const centerPosition = new google.maps.LatLng(50.119504, 8.638137)
    const { result } = renderHook(() => useMap(centerPosition))
    await waitFor(() => result.current.map)
    expect(result.current.map).toEqual(undefined)
    expect(result.current.mapRef).toEqual(expect.any(Function))
    act(() => {
      result.current.mapRef(null)
    })
    await waitFor(() => {
      expect(result.current.map).toMatchSnapshot()
    })
  })
})
