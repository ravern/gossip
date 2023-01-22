import { CircularProgress, List } from "@mui/material";
import React from "react";

import usePostsQuery from "src/api/queries/posts";

import PostListItem from "./components/PostListItem";

export default function PostList() {
  const { data: posts, isLoading } = usePostsQuery();

  if (isLoading) {
    return <CircularProgress />;
  } else if (posts != null) {
    return (
      <List>
        {posts.map((post) => {
          return <PostListItem key={post.id} post={post} />;
        })}
      </List>
    );
  } else {
    throw new Error("invalid state");
  }
}
