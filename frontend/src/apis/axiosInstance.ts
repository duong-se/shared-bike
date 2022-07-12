import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import { commonHeaders, tokenKey } from "../constants/constants";

export enum HTTP_STATUS {
  UNAUTHORIZED = 403,
  OK = 200,
}

export const axiosApiInstance = axios.create({
  headers: commonHeaders,
});

export const handleInterceptRequestError = (error: any) => {
  throw error;
}

export const handleInterceptConfig =  (config: AxiosRequestConfig<any>) => {
  const token = localStorage.getItem(tokenKey);
  if (!token) {
    return config;
  }
  config.headers = {
    ...config.headers,
    Authorization: `Bearer ${token}` as string,
  };
  return config;
}

axiosApiInstance.interceptors.request.use(
  handleInterceptConfig,
  handleInterceptRequestError
);

export const handleInterceptResponse = (response: AxiosResponse<any, any>) => {
  return response;
}

export const handleInterceptResponseError = (error: any) => {
  if (error?.response?.status === HTTP_STATUS.UNAUTHORIZED) {
    localStorage.clear();
    window.location.href = "/";
  }
  if(error instanceof AxiosError) {
    throw new Error(error.response?.data)
  }
  throw error;
}

axiosApiInstance.interceptors.response.use(
  handleInterceptResponse,
  handleInterceptResponseError,
);
