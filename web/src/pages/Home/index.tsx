import { Container } from "@mui/system";
import React from "react";

import PostList from "src/components/PostList";
import BaseLayout from "src/layouts/Base";

export default function HomePage() {
  return (
    <BaseLayout>
      <Container maxWidth="sm">
        <PostList />
      </Container>
    </BaseLayout>
  );
}
