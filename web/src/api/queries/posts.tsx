import { useQuery } from "react-query";

import { axiosClient } from "src/api";

export default function usePostsQuery() {
  return useQuery("posts", async () => {
    const response = await axiosClient.get("/posts");
    console.log(response.data);
  });
}
