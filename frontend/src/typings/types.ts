export type LoginVariables = {
  username: string
  password: string
}

export type LoginResponse = {
  accessToken: string
}

export type User = {
  id: number
  name: string
  username: string
}

export type RentBikeVariables = {
  bikeId: number
}

export type ReturnBikeVariables = {
  bikeId: number
}

export enum BikeStatus {
  RENTED = "rented",
  AVAILABLE = "available"
}

export type Bike = {
  id: number
  name: string
  lat: string
  long: string
  status: BikeStatus
  userId?: number
  nameOfRenter?: string
  usernameOfRenter?: string
}

export type RegisterVariables = {
  username: string
  password: string
  name: string
}

export type RegisterResponse = {
  accessToken: string
}
