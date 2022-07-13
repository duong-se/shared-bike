import { useCallback, useState } from 'react'

export type UseMapReturn = {
  mapRef: (node: HTMLDivElement | null) => void
  map?: google.maps.Map
}

export const useMap = (position: google.maps.LatLng): UseMapReturn => {
  const [map, setMap] = useState<google.maps.Map>()
  const mapRef = useCallback((node: HTMLDivElement | null) => {
    if (node !== null) {
      const map = new window.google.maps.Map(node, {
        center: position,
        zoom: 8,
        fullscreenControl: false,
      })
      setMap(map)
    }
  }, [])
  return {
    map,
    mapRef,
  }
}
