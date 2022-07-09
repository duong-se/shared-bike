import { useState } from "react"

export enum BikeStatus {
  RENTED = "rented",
  AVAILABLE = "available"
}

type Bike = {
  id: number
  lat: string
  long: string
  status: BikeStatus
  userId?: number
  nameOfRenter?: string
  usernameOfRenter?: string
}

export const useBikes = () => {
  const [error, setError] = useState<string| undefined>(undefined)
  const [isFetchLoading, setLoading] = useState<boolean>(true)
  const [isActionLoading, setActionLoading] = useState<boolean>(true)
  const [bikes, setBikes] = useState<Bike[]>([])
  const fetchAllBikes = async () => {
    setLoading(true)
    const getBikesUrl = `${window.sharedBike.config.baseUrl}/bikes`
    fetch(getBikesUrl, {
      method: "GET",
      credentials: "include",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
    }).then(async (resp) => {
      setLoading(false)
      if (resp.ok) {
        const bikes = await resp.json()
        return setBikes(bikes)
      }
    }).catch((error) => {
      setLoading(false)
      setError(error)
    })
  }

  const updateListBike = async (bikeId: number, resp: Bike) => {
    const filteredBikes = bikes.filter((item) => item.id !== bikeId)
    setBikes([...filteredBikes, resp])
  }

  const returnBike = async (bikeId: number) => {
    setActionLoading(true)
    const returnBikeUrl = `${window.sharedBike.config.baseUrl}/bikes/${bikeId}/return`
    fetch(returnBikeUrl, {
      method: "PATCH",
      credentials: "include",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
    }).then(async (resp) => {
      setActionLoading(false)
      if (resp.ok) {
        const bikeUpdated = await resp.json()
        updateListBike(bikeId, bikeUpdated)
      }
    }).catch((error) => {
      setActionLoading(false)
      setError(error)
    })
  }

  const rentBike = async (bikeId: number) => {
    setActionLoading(true)
    const returnBikeUrl = `${window.sharedBike.config.baseUrl}/bikes/${bikeId}/rent`
    fetch(returnBikeUrl, {
      method: "PATCH",
      credentials: "include",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
    }).then(async (resp) => {
      setActionLoading(false)
      if (resp.ok) {
        const bikeUpdated = await resp.json()
        updateListBike(bikeId, bikeUpdated)
      }
    }).catch((error) => {
      setActionLoading(false)
      setError(error)
    })
  }

  return {
    fetchAllBikes,
    returnBike,
    rentBike,
    error,
    bikes,
    isFetchLoading,
    isActionLoading
  }
}
