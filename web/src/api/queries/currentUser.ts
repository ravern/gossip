import { useQuery } from "react-query";

import { axiosClient } from "src/api";

import type { CurrentUserData, DataResponse } from "../models";

export default function useCurrentUserQuery() {
  return useQuery("currentUser", async () => {
    const response = await axiosClient.get<DataResponse<CurrentUserData>>(
      "/user"
    );
    return response.data.data;
  });
}
