import { useMutation, useQueryClient } from "react-query";

import { axiosClient } from "src/api";

import { CommentData, DataResponse } from "../models";

export interface DeleteCommentParams {
  postId: string;
  commentId: string;
}

async function deleteComment({ postId, commentId }: DeleteCommentParams) {
  const response = await axiosClient.delete<DataResponse<CommentData>>(
    `/posts/${postId}/comments/${commentId}`
  );
  return response.data.data;
}

export default function useDeleteCommentMutation() {
  const queryClient = useQueryClient();
  return useMutation(deleteComment, {
    onSuccess: (_data, { postId }) => {
      queryClient.refetchQueries(["posts", postId]);
    },
  });
}
