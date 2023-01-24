import { useMutation, useQueryClient } from "react-query";

import { axiosClient } from "src/api";

import { CommentData, DataResponse } from "../models";

export interface CreateCommentParams {
  postId: string;
  body: string;
}

async function createComment({ postId, body }: CreateCommentParams) {
  const response = await axiosClient.post<DataResponse<CommentData>>(
    `/posts/${postId}/comments`,
    {
      postId,
      body,
    }
  );
  return response.data.data;
}

export default function useCreateCommentMutation() {
  const queryClient = useQueryClient();
  return useMutation(createComment, {
    onSuccess: (_data, { postId }) => {
      queryClient.refetchQueries(["posts", postId]);
    },
  });
}
