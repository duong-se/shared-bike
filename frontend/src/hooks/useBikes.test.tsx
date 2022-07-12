import { renderHook, waitFor } from "@testing-library/react";
import React from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { useBikes, useRentBike, useReturnBike } from "./useBikes";
import { axiosApiInstance } from '../apis/axiosInstance'
import { Bike } from "../typings/types";

describe("useBikes", () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });
  const Provider: React.FC<React.PropsWithChildren> = ({ children }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
  it("should run and return loading false", async () => {
    jest.spyOn(axiosApiInstance, 'get').mockResolvedValue([
      {
        id: 1,
        name: "mockName"
      }
    ] as Bike[])
    const { result } = renderHook(() => useBikes({}), { wrapper: Provider });
    expect(result.current.isLoading).toEqual(true);
    await waitFor(() => result.current.isSuccess);
    await waitFor(() => {
      expect(result.current.isLoading).toEqual(false);
    })
  });
});


describe("useRentBike", () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });
  const Provider: React.FC<React.PropsWithChildren> = ({ children }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
  it("should run and return loading false", async () => {
    const { result } = renderHook(() => useRentBike({}), { wrapper: Provider });
    await waitFor(() => result.current.isSuccess);
    expect(result.current.isLoading).toEqual(false);
  });
});

describe("useReturnBike", () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });
  const Provider: React.FC<React.PropsWithChildren> = ({ children }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
  it("should run and return loading false", async () => {
    const { result } = renderHook(() => useReturnBike({}), { wrapper: Provider });
    await waitFor(() => result.current.isSuccess);
    expect(result.current.isLoading).toEqual(false);
  });
});
