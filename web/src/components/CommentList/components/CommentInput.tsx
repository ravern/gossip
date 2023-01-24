import { Box, Button, TextField } from "@mui/material";
import React, { useState } from "react";

import useCreateCommentMutation from "src/api/mutations/createComment";

export interface CommentInputProps {
  postId: string;
}

export default function CommentInput({ postId }: CommentInputProps) {
  const { mutateAsync: createComment } = useCreateCommentMutation();

  const [body, setBody] = useState("");

  const handleBodyChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setBody(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    createComment({ postId, body })
      .then(() => setBody(""))
      .catch(console.error);
  };

  return (
    <Box
      component="form"
      alignItems="center"
      sx={{ display: "flex" }}
      onSubmit={handleSubmit}
    >
      <TextField
        margin="dense"
        size="small"
        multiline
        placeholder="Add your comment"
        value={body}
        onChange={handleBodyChange}
        sx={{ flexGrow: 1 }}
      />
      <Button type="submit" variant="contained" sx={{ marginLeft: 2 }}>
        Post
      </Button>
    </Box>
  );
}
