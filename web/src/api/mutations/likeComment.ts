import { useMutation, useQueryClient } from "react-query";

import { axiosClient } from "src/api";

export interface LikeCommentParams {
  postId: string;
  commentId: string;
  isLiked: boolean;
}

async function likeComment({ postId, commentId, isLiked }: LikeCommentParams) {
  const response = await axiosClient.post(
    `/posts/${postId}/comments/${commentId}/likes`,
    {
      is_liked: isLiked,
    }
  );
  return response.data;
}

export default function useLikeCommentMutation() {
  const queryClient = useQueryClient();
  return useMutation(likeComment, {
    onSuccess: (_data, { postId }) => {
      queryClient.refetchQueries(["posts", postId]);
    },
  });
}
