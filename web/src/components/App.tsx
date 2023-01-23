import { CssBaseline } from "@mui/material";
import React from "react";
import { QueryClientProvider } from "react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";

import { queryClient } from "src/api";
import HomePage from "src/pages/Home";
import NewPostPage from "src/pages/NewPost";
import PostPage from "src/pages/Post";

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/posts/new",
    element: <NewPostPage />,
  },
  {
    path: "/posts/:id",
    element: <PostPage />,
  },
]);

export default function App() {
  return (
    <>
      <CssBaseline />
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </>
  );
}
