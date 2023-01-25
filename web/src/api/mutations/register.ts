import { useMutation } from "react-query";

import { axiosClient } from "src/api";

import type { UserData } from "../models";

export interface RegisterParams {
  handle: string;
  email: string;
  password: string;
  avatarURL: string | null;
}

export interface RegisterData {
  user: UserData;
}

async function register({
  handle,
  email,
  password,
  avatarURL,
}: RegisterParams) {
  const response = await axiosClient.post("/auth/register", {
    handle,
    email,
    password,
    avatar_url: avatarURL,
  });
  return response.data.data;
}

export default function useRegisterMutation() {
  return useMutation(register);
}
