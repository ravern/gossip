import { useQuery } from "react-query";

import { axiosClient } from "src/api";

import type { CommentData, DataResponse } from "../models";

async function getComments(postId: string) {
  const response = await axiosClient.get<DataResponse<CommentData[]>>(
    `/posts/${postId}/comments`
  );
  return response.data.data;
}

export default function useCommentsQuery(postId: string) {
  return useQuery(["posts", postId, "comments"], () => getComments(postId));
}
