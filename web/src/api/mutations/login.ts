import { useMutation } from "react-query";

import { axiosClient } from "src/api";

import type { UserData } from "../models";

export interface LoginParams {
  handleOrEmail: string;
  password: string;
}

export interface LoginData {
  user: UserData;
}

async function login({ handleOrEmail, password }: LoginParams) {
  const response = await axiosClient.post("/auth/login", {
    handle_or_email: handleOrEmail,
    password,
  });
  return response.data.data;
}

export default function useLoginMutation() {
  return useMutation(login);
}
