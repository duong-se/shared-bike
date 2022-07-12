import { User } from "../typings/types";
import { axiosApiInstance, HTTP_STATUS } from "./axiosInstance";
import { login, register } from "./users";

describe("login", () => {
  it("should run correctly", async () => {
    const mockLoginResponse = {
      id: 1,
      name: "mockName",
      username: "mockUsername",
    } as User;
    jest.spyOn(axiosApiInstance, "post").mockResolvedValue({
      data: mockLoginResponse,
      status: HTTP_STATUS.OK,
    });
    const result = await login({
      username: "mockUsername",
      password: "mockPassword",
    });
    expect(result).toEqual(mockLoginResponse);
  });

  it("should run and throw error", async () => {
    jest.spyOn(axiosApiInstance, "post").mockRejectedValue(new Error('mockError'));
    let err
    try {
      await login({
        username: "mockUsername",
        password: "mockPassword",
      });
    } catch (error) {
      err = error
    }
    expect(err).toEqual('mockError');
  });
});


describe("register", () => {
  it("should run correctly", async () => {
    const mockLoginResponse = {
      id: 1,
      name: "mockName",
      username: "mockUsername",
    } as User;
    jest.spyOn(axiosApiInstance, "post").mockResolvedValue({
      data: mockLoginResponse,
      status: HTTP_STATUS.OK,
    });
    const result = await register({
      username: "mockUsername",
      password: "mockPassword",
      name: "mockName",
    });
    expect(result).toEqual(mockLoginResponse);
  });

  it("should run and throw error", async () => {
    jest.spyOn(axiosApiInstance, "post").mockRejectedValue(new Error('mockError'));
    let err
    try {
      await register({
        username: "mockUsername",
        password: "mockPassword",
        name: "mockName",
      });
    } catch (error) {
      err = error
    }
    expect(err).toEqual('mockError');
  });
});
