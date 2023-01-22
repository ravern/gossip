import { useQuery } from "react-query";

import { axiosClient } from "src/api";

import type { DataResponse, PostData } from "../models";

async function getPosts() {
  const response = await axiosClient.get<DataResponse<PostData[]>>("/posts");
  return response.data.data;
}

export default function usePostsQuery() {
  return useQuery("posts", getPosts);
}
