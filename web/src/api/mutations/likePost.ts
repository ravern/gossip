import { useMutation, useQueryClient } from "react-query";

import { axiosClient } from "src/api";

export interface LikePostParams {
  postId: string;
  isLiked: boolean;
}

async function likePost({ postId, isLiked }: LikePostParams) {
  const response = await axiosClient.post(`/posts/${postId}/likes`, {
    is_liked: isLiked,
  });
  return response.data;
}

export default function useLikePostMutation() {
  const queryClient = useQueryClient();
  return useMutation(likePost, {
    onSuccess: () => {
      queryClient.refetchQueries("posts");
    },
  });
}
