import { AxiosResponse } from "axios";
import { LoginResponse, LoginVariables, RegisterResponse, RegisterVariables } from "../typings/types";
import { axiosApiInstance } from "./axiosInstance";

export const login = ({
  username,
  password,
}: LoginVariables): Promise<LoginResponse> => {
  const loginUrl = `${window.sharedBike.config.baseUrl}/users/login`;
  return axiosApiInstance
    .post<LoginVariables, AxiosResponse<LoginResponse>>(loginUrl, {
      username,
      password,
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error(error);
      throw (error as Error).message;
    });
};

export const register = ({
  username,
  password,
  name,
}: RegisterVariables): Promise<RegisterResponse> => {
  const registerUrl = `${window.sharedBike.config.baseUrl}/users/register`;
  return axiosApiInstance
    .post<RegisterVariables, AxiosResponse<RegisterResponse>>(registerUrl, {
      username,
      password,
      name
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error(error);
      throw (error as Error).message;
    });
};
