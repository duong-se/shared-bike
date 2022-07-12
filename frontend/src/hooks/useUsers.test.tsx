import { renderHook, waitFor } from "@testing-library/react";
import React from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { useLogin, useRegister } from "./useUsers";

describe("useLogin", () => {
  const queryClient = new QueryClient();
  const Provider: React.FC<React.PropsWithChildren> = ({ children }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
  it("should run and return loading false", async () => {
    const { result } = renderHook(() => useLogin({}), { wrapper: Provider });
    await waitFor(() => result.current.isSuccess);
    expect(result.current.isLoading).toEqual(false);
  });
});


describe("useRegister", () => {
  const queryClient = new QueryClient();
  const Provider: React.FC<React.PropsWithChildren> = ({ children }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
  it("should run and return loading false", async () => {
    const { result } = renderHook(() => useRegister({}), { wrapper: Provider });
    await waitFor(() => result.current.isSuccess);
    expect(result.current.isLoading).toEqual(false);
  });
});
