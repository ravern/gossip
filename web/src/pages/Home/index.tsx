import React from "react";

import usePostsQuery from "src/api/queries/posts";

export default function HomePage() {
  const { data } = usePostsQuery();
  console.log(data);
  return <>Test</>;
}
