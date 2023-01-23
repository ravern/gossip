import { Container, TextareaAutosize, TextField } from "@mui/material";
import React, { useState } from "react";

import BaseLayout from "src/layouts/Base";

export default function NewPostPage() {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };

  const handleBodyChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setBody(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
  };

  return (
    <BaseLayout>
      <Container maxWidth="md">
        <form onSubmit={handleSubmit}>
          <TextField
            margin="dense"
            autoFocus
            label="Title"
            fullWidth
            variant="standard"
            value={title}
            onChange={handleTitleChange}
          />
          <TextField
            margin="dense"
            autoFocus
            label="Body"
            fullWidth
            variant="standard"
            multiline
            value={body}
            onChange={handleBodyChange}
          />
        </form>
      </Container>
    </BaseLayout>
  );
}
