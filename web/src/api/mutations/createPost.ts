import { useMutation, useQueryClient } from "react-query";

import { axiosClient } from "src/api";

import { DataResponse, PostData } from "../models";

export interface CreatePostParams {
  title: string;
  body: string;
  tags: string[];
}

async function createPost({ title, body, tags }: CreatePostParams) {
  const response = await axiosClient.post<DataResponse<PostData>>(`/posts`, {
    title,
    body,
    tags,
  });
  return response.data.data;
}

export default function useCreatePostMutation() {
  const queryClient = useQueryClient();
  return useMutation(createPost, {
    onSuccess: () => {
      queryClient.refetchQueries("posts");
    },
  });
}
