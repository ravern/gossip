import { useQuery } from "react-query";

import { axiosClient } from "src/api";

import type { DataResponse, PostData } from "../models";

async function getPost(id: string) {
  const response = await axiosClient.get<DataResponse<PostData>>(
    `/posts/${id}`
  );
  return response.data.data;
}

export default function usePostQuery(id: string | undefined) {
  return useQuery(["posts", id], () => getPost(id!), { enabled: id != null });
}
