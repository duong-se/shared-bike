import { useMutation, UseMutationOptions, useQuery, UseQueryOptions } from "react-query"
import { fetchBikes, rentBike, returnBike } from "../apis/bikes"
import { Bike, RentBikeVariables, ReturnBikeVariables } from "../typings/types"

export const useBikes = (options?: Omit<UseQueryOptions<Bike[], unknown, Bike[], string[]>, "queryKey" | "queryFn">) => {
  return useQuery(['bikes'], fetchBikes, options)
}

export const useRentBike = (options?: Omit<UseMutationOptions<Bike, unknown, RentBikeVariables, unknown>, "mutationFn">) => {
  return useMutation(rentBike, options)
}

export const useReturnBike = (options?: Omit<UseMutationOptions<Bike, unknown, ReturnBikeVariables, unknown>, "mutationFn">) => {
  return useMutation(returnBike, options)
}
