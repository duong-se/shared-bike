import { useCallback, useEffect, useState } from 'react';
import { Map, tileLayer, Marker } from 'leaflet'
import { useBikes } from '../hooks/useBikes'


export const BikeMapPage: React.FC = () => {
  const { fetchAllBikes, rentBike, returnBike, bikes, isFetchLoading } = useBikes()
  const [map, setMap] = useState<Map>()
  const position = {
    lat: 50.119504,
    lng: 8.638137,
  }
  const mapRef = useCallback((node: HTMLDivElement | null) => {
    if (node !== null) {
      const map = new Map(node, { center: position, zoom: 13 }).setView(position, 13)
      tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
        maxZoom: 18
      }).addTo(map);
      setMap(map)
    }
  }, [])
  useEffect(() => {
    fetchAllBikes()
  }, [])
  useEffect(() => {
    const markers = bikes.map((bike) => {
      const marker = new Marker({ lat: Number(bike.lat), lng: Number(bike.long) }).addTo(map as Map)
    })
  }, [bikes])
  useEffect(() => {
    return () => {
      map?.remove()
    }
  }, [map])

  if (isFetchLoading) {
    return <div>Loading...</div>
  }
  return (
    <div style={{ height: '100vh' }} ref={mapRef}></div>
    //   <MapContainer style={{ height: '100vh' }} center={position} zoom={13} scrollWheelZoom={false}>
    //   <TileLayer
    //     attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    //     url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
    //   />
    //   {bikes.map((item) => {
    //     const bikePosition = {
    //       lat: Number(item.lat),
    //       lng: Number(item.long)
    //     }
    //     return (
    //       <Marker position={bikePosition}>
    //         <Popup>
    //           <RentBikePopUp
    //             isRented={item.status === BikeStatus.RENTED}
    //             isDisabledAction={false}
    //             onRent={() => {
    //               rentBike(item.id)
    //             }}
    //             onReturn={() => {
    //               returnBike(item.id)
    //             }}
    //           />
    //         </Popup>
    //     </Marker>
    //     )
    //   })}
    // </MapContainer>
  );
}

export default BikeMapPage
