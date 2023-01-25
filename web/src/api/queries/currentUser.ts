import { useQuery } from "react-query";

import { axiosClient, LOCAL_STORAGE_KEY_ACCESS_TOKEN } from "src/api";
import { parse } from "src/jwt";

import type { CurrentUserData, DataResponse } from "../models";

export default function useCurrentUserQuery() {
  return useQuery("currentUser", async () => {
    const response = await axiosClient.get<DataResponse<CurrentUserData>>(
      "/user"
    );
    response.data.data.role = parse(
      localStorage.getItem(LOCAL_STORAGE_KEY_ACCESS_TOKEN) ?? ""
    ).role;
    return response.data.data;
  });
}
