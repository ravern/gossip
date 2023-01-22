import { Container } from "@mui/system";
import React from "react";

import usePostsQuery from "src/api/queries/posts";
import PostList from "src/components/PostList";
import BaseLayout from "src/layouts/Base";

export default function PostPage() {
  const { data } = usePostsQuery();
  console.log(data);
  return (
    <BaseLayout>
      <Container maxWidth="md">
        <PostList />
      </Container>
    </BaseLayout>
  );
}
