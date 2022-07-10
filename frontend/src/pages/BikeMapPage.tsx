import { memo, useCallback, useEffect } from 'react';
import { MarkerClusterer } from "@googlemaps/markerclusterer";
import { useBikes, Bike, BikeStatus } from '../hooks/useBikes'
import { useMap } from '../hooks/useMap';
import { useAuth } from '../hooks/AuthProvider';
import { Spinner } from '../components/Spinner';

function customIcon(color: string) {
  return {
    path: 'M 0,0 C -2,-20 -10,-22 -10,-30 A 10,10 0 1,1 10,-30 C 10,-22 2,-20 0,0 z M -2,-30 a 2,2 0 1,1 4,0 2,2 0 1,1 -4,0',
    fillColor: color,
    fillOpacity: 1,
    strokeColor: '#000',
    strokeWeight: 2,
    scale: 1,
  }
}

const availableColor = '#2ecc71'
const rentedColor = '#34495e'

const contentMap = {
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

const renderCommonContent = (bike: Bike) => {
  return `
    <h1 class="text-2xl">Bike &raquo;${bike.id}&laquo;</h1>
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

const renderUserHasBikeCase = (
  bikes: Array<Bike>,
  bikeOfUser: Bike,
  handleReturnBike: (bikeId: number, infoWindow: google.maps.InfoWindow) => (this: HTMLElement, ev: MouseEvent) => any,
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
    marker.addListener("click", () => {
      infoWindow.open({
        anchor: marker,
        map,
        shouldFocus: false
      });
    });
    if (isUserBike) {
      infoWindow.addListener('domready', () => {
        const button = document.getElementById(`bike-action-${bike.id}`)
        if (button) {
          button?.addEventListener('click', handleReturnBike(bike.id, infoWindow))
        }
      })
    }
    return marker
  })
}

const renderUserHasNoBikeCase = (
  bikes: Array<Bike>,
  handleRentBike: (bikeId: number, infoWindow: google.maps.InfoWindow) => (this: HTMLElement, ev: MouseEvent) => any,
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
    marker.addListener("click", () => {
      infoWindow.open({
        anchor: marker,
        map,
        shouldFocus: false
      });
    });
    if (bike.status === BikeStatus.AVAILABLE) {
      infoWindow.addListener('domready', () => {
        const button = document.getElementById(`bike-action-${bike.id}`)
        if (button) {
          button?.addEventListener('click', handleRentBike(bike.id, infoWindow))
        }
      })
    }
    return marker
  })
}


export const BikeMapPage: React.FC = () => {
  const { rentBike, returnBike, bikes, isFetchLoading } = useBikes()
  const { user } = useAuth()
  const centerPosition = new google.maps.LatLng(50.119504, 8.638137)
  const { map, mapRef } = useMap(centerPosition)
  const handleRentBike = useCallback((bikeId: number,  infoWindow: google.maps.InfoWindow) => function (this: HTMLElement, ev: MouseEvent): any {
    infoWindow.close()
    rentBike(bikeId)
  }, [rentBike])
  const handleReturnBike = useCallback((bikeId: number, infoWindow: google.maps.InfoWindow) => function (this: HTMLElement, ev: MouseEvent): any {
    infoWindow.close()
    returnBike(bikeId)
  }, [returnBike])
  useEffect(() => {
    const userBike = bikes.find((item) => item.userId === user?.id)
    if (userBike) {
      const markers = renderUserHasBikeCase(bikes, userBike, handleReturnBike, map)
      new MarkerClusterer({ markers, map });
      return
    }
    const markers = renderUserHasNoBikeCase(bikes, handleRentBike, map)
    new MarkerClusterer({ markers, map });
  }, [map, bikes, handleRentBike, user?.id, handleReturnBike])

  if (isFetchLoading) {
    return <Spinner />
  }
  return (
    <div style={{ height: '100vh' }} ref={mapRef}></div>
  );
}

export default memo(BikeMapPage)
