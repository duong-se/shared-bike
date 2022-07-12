import { AxiosResponse } from "axios";
import { Bike, RentBikeVariables, ReturnBikeVariables } from "../typings/types";
import { axiosApiInstance } from "./axiosInstance";

export const fetchBikes = (): Promise<Bike[]> => {
  const getBikesUrl = `${window.sharedBike.config.baseUrl}/bikes`;
  return axiosApiInstance
    .get<undefined, AxiosResponse<Bike[]>>(getBikesUrl)
    .then((resp) => resp.data)
    .catch((error) => {
      console.error(error);
      throw (error as Error).message;
    });
};

export const rentBike = ({ bikeId }: RentBikeVariables): Promise<Bike> => {
  const returnBikeUrl = `${window.sharedBike.config.baseUrl}/bikes/${bikeId}/rent`;
  return axiosApiInstance
    .patch<undefined, AxiosResponse<Bike>>(returnBikeUrl)
    .then((resp) => resp.data)
    .catch((error) => {
      console.error(error);
      throw (error as Error).message;
    });
};

export const returnBike = ({ bikeId }: ReturnBikeVariables): Promise<Bike> => {
  const returnBikeUrl = `${window.sharedBike.config.baseUrl}/bikes/${bikeId}/return`;
  return axiosApiInstance
    .patch<undefined, AxiosResponse<Bike>>(returnBikeUrl)
    .then((resp) => resp.data)
    .catch((error) => {
      console.error(error);
      throw (error as Error).message;
    });
};
