import { useMutation, useQueryClient } from "react-query";

import { axiosClient } from "src/api";

import { DataResponse, PostData } from "../models";

async function deletePost(postId: string) {
  const response = await axiosClient.delete<DataResponse<PostData>>(
    `/posts/${postId}`
  );
  return response.data.data;
}

export default function useDeletePostMutation() {
  const queryClient = useQueryClient();
  return useMutation(deletePost, {
    onSuccess: () => {
      queryClient.refetchQueries("posts");
    },
  });
}
