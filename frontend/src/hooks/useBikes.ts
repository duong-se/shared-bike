import { useEffect, useState } from "react"
import { tokenKey } from "./AuthProvider"

export enum BikeStatus {
  RENTED = "rented",
  AVAILABLE = "available"
}

export type Bike = {
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
  const token = localStorage.getItem(tokenKey)
  useEffect(() => {
    setLoading(true)
    const getBikesUrl = `${window.sharedBike.config.baseUrl}/bikes`
    fetch(getBikesUrl, {
      method: "GET",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json",
        "Authorization": token as string,
      },
    }).then(async resp => {
      const data = await resp.json()
      if (resp.ok) {
        return data
      }
      throw new Error(data)
    }).then(async (result) => {
      setBikes(result)
    }).catch((error) => {
      setError(error)
    }).finally(() => {
      setLoading(false)
    })
  }, [token])

  const updateListBike = async (bikeId: number, resp: Bike) => {
    const filteredBikes = bikes.filter((item) => item.id !== bikeId)
    setBikes([...filteredBikes, resp])
  }

  const returnBike = async (bikeId: number) => {
    setActionLoading(true)
    const returnBikeUrl = `${window.sharedBike.config.baseUrl}/bikes/${bikeId}/return`
    fetch(returnBikeUrl, {
      method: "PATCH",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json",
        "Authorization": token as string,
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
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json",
        "Authorization": token as string,
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
    returnBike,
    rentBike,
    error,
    bikes,
    isFetchLoading,
    isActionLoading
  }
}
