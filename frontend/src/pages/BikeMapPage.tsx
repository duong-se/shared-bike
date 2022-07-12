import { memo, useEffect } from 'react';
import { MarkerClusterer } from "@googlemaps/markerclusterer";
import { useMap } from '../hooks/useMap';
import { useAuth } from '../hooks/AuthProvider';
import { Spinner } from '../components/Spinner';
import { Bike, BikeStatus, RentBikeVariables, ReturnBikeVariables } from '../typings/types';
import { MutateOptions, QueryClient, useQueryClient } from 'react-query';
import { AlertError } from '../components/AlertError';
import { useBikes, useRentBike, useReturnBike } from '../hooks/useBikes';

export const customIcon = (color: string) => {
  return {
    path: 'M 0,0 C -2,-20 -10,-22 -10,-30 A 10,10 0 1,1 10,-30 C 10,-22 2,-20 0,0 z M -2,-30 a 2,2 0 1,1 4,0 2,2 0 1,1 -4,0',
    fillColor: color,
    fillOpacity: 1,
    strokeColor: '#000',
    strokeWeight: 2,
    scale: 1,
  }
}

export const availableColor = '#2ecc71'
export const rentedColor = '#34495e'

export const contentMap = {
  [BikeStatus.AVAILABLE]: {
    icon: customIcon(availableColor),
    renderButton: (bike: Bike) => {
      return `
      <div class="flex justify-end">
        <button
          id="bike-action-${bike.id}"
          type="button"
          class="w-1/2 py-2.5 rounded-3xl mt-4 btn-primary"
        >
          <div class="flex flex-row items-center justify-center">
            <div class="font-light">RENT BIKE</div>
          </div>
        </button>
      </div>
      `
    }
  },
  [BikeStatus.RENTED]: {
    icon: customIcon(rentedColor),
    renderButton: (bike: Bike) => {
      return `
      <div class="flex justify-end">
        <button
          id="bike-action-${bike.id}"
          type="button"
          disabled
          class="w-1/2 py-2.5 rounded-3xl mt-4 btn-disabled"
        >
          <div class="flex flex-row items-center justify-center">
            <div class="font-light">RENT BIKE</div>
          </div>
        </button>
      </div>
      `
    }
  },
  returnBike: {
    icon: customIcon(availableColor),
    renderButton: (bike: Bike) => {
      return `
        <div class="flex justify-end">
          <button
            id="bike-action-${bike.id}"
            type="button"
            class="w-1/2 py-2.5 rounded-3xl mt-4 btn-primary"
          >
            <div class="flex flex-row items-center justify-center">
              <div class="font-light">RETURN BIKE</div>
            </div>
          </button>
        </div>
        </div>
      `
    }
  },
}

export const renderCommonContent = (bike: Bike) => {
  return `
    <h1 class="text-2xl">Bike &raquo;${bike.name}&laquo;</h1>
    <h2 class="font-light text-lg">This bike for rent</h2>
    <div class="container mx-2">
      <ol class="list-decimal py-2 px-4">
        <li>Click on &ldquo;Rent Bike&rdquo;</li>
        <li>Bicycle lock will unlock automatically</li>
        <li>Adjust saddle height</li>
      </ol>
    </div>
  `
}

export const renderUserHasBikeCase = (
  bikes: Array<Bike>,
  bikeOfUser: Bike,
  returnBikeMutate: (variables: ReturnBikeVariables, options?: MutateOptions<Bike, unknown, ReturnBikeVariables, unknown> | undefined) => void,
  map?: google.maps.Map,
): Array<google.maps.Marker> => {
  if (!map) {
    return []
  }
  return bikes.map((bike) => {
    const isUserBike = bikeOfUser.id === bike.id
    const marker = new google.maps.Marker({
      position: { lat: Number(bike.lat), lng: Number(bike.long) },
      map: map,
      title: "title"
    })
    const infoWindow = new google.maps.InfoWindow();
    let { icon, renderButton } = contentMap[BikeStatus.RENTED]
    if (isUserBike) {
      icon = contentMap.returnBike.icon
      renderButton = contentMap.returnBike.renderButton
    }
    marker.setIcon(icon)
    const popUpContent = `
      <div class="p-2">
        ${renderCommonContent(bike)}
        ${renderButton(bike)}
      </div>
    `
    infoWindow.setContent(popUpContent)
    marker.addListener("click", handleMarkerCallback(infoWindow, marker, map));
    if (isUserBike) {
      infoWindow.addListener('domready', handlePopUpButtonCallback(returnBikeMutate, bike, infoWindow))
    }
    return marker
  })
}

export const handlePopUpButtonCallback = (
  fn: (variables: ReturnBikeVariables | RentBikeVariables, options?: MutateOptions<Bike, unknown, ReturnBikeVariables | RentBikeVariables, unknown> | undefined) => void,
  bike: Bike,
  infoWindow: google.maps.InfoWindow,
) => () => {
  const button = document.getElementById(`bike-action-${bike.id}`)
  if (button) {
    button?.addEventListener('click', () => {
      fn({ bikeId: bike.id })
      infoWindow.close()
    })
  }
}

export const handleMarkerCallback = (infoWindow: google.maps.InfoWindow, marker: google.maps.Marker, map: google.maps.Map) => () => {
  infoWindow.open({
    anchor: marker,
    map,
    shouldFocus: false
  });
}


export const renderUserHasNoBikeCase = (
  bikes: Array<Bike>,
  rentBikeMutate: (variables: RentBikeVariables, options?: MutateOptions<Bike, unknown, RentBikeVariables, unknown> | undefined) => void,
  map?: google.maps.Map,
): Array<google.maps.Marker> => {
  if (!map) {
    return []
  }
  return bikes.map((bike) => {
    const { icon, renderButton } = contentMap[bike.status]
    const marker = new google.maps.Marker({
      position: { lat: Number(bike.lat), lng: Number(bike.long) },
      map: map,
      title: "title"
    })
    marker.setIcon(icon)
    const popUpContent = `
      <div class="p-2">
        ${renderCommonContent(bike)}
        ${renderButton(bike)}
      </div>
    `
    const infoWindow = new google.maps.InfoWindow({
      content: popUpContent,
    });
    marker.addListener("click", handleMarkerCallback(infoWindow, marker, map));
    if (bike.status === BikeStatus.AVAILABLE) {
      infoWindow.addListener('domready', handlePopUpButtonCallback(rentBikeMutate, bike, infoWindow))
    }
    return marker
  })
}

export const onRentOrReturnSuccess = (queryClient: QueryClient) => (data: Bike) => {
  const currentBikes = queryClient.getQueryData<Bike[]>(['bikes'])
  if (currentBikes) {
    const filteredBikes = currentBikes.filter((bike) => bike.id !== data.id)
    const newBikes = [...filteredBikes, data]
    queryClient.setQueryData(['bikes'], newBikes)
  }
}

export const BikeMapPage: React.FC = () => {
  const { user } = useAuth()
  const queryClient = useQueryClient()
  const {
    data: bikes,
    isLoading: isFetchBikesLoading,
    isError: isFetchBikesError,
    error: fetchBikesError
  } = useBikes()
  const {
    mutate: rentBikeMutate,
  } = useRentBike({
    onSuccess: onRentOrReturnSuccess(queryClient)
  })
  const {
    mutate: returnBikeMutate,
  } = useReturnBike({
    onSuccess: onRentOrReturnSuccess(queryClient)
  })
  const centerPosition = new google.maps.LatLng(50.119504, 8.638137)
  const { map, mapRef } = useMap(centerPosition)

  useEffect(() => {
    if (bikes) {
      const userBike = bikes.find((item) => item.userId === user?.id)
      if (userBike) {
        const markers = renderUserHasBikeCase(bikes, userBike, returnBikeMutate, map)
        new MarkerClusterer({ markers, map });
        return
      }
      const markers = renderUserHasNoBikeCase(bikes, rentBikeMutate, map)
      new MarkerClusterer({ markers, map });
      return
    }
  }, [map, bikes, user?.id, user, rentBikeMutate, returnBikeMutate])

  if (isFetchBikesLoading) {
    return <Spinner />
  }
  if (isFetchBikesError) {
    return <AlertError error={fetchBikesError as string} />
  }
  return (
    <div data-testid="map-container" style={{ height: '100vh' }} ref={mapRef}></div>
  );
}

export default memo(BikeMapPage)
