import {
  Button,
  Card,
  CardContent,
  CircularProgress,
  Container,
  TextField,
  Typography,
} from "@mui/material";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

import useCreatePostMutation from "src/api/mutations/createPost";
import BaseLayout from "src/layouts/Base";

export default function NewPostPage() {
  const navigate = useNavigate();

  const { mutateAsync: createPost, isLoading } = useCreatePostMutation();

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

    createPost({ title, body, tags: [] })
      .then(() => {
        navigate("/");
      })
      .catch(console.error);
  };

  return (
    <BaseLayout>
      <Container maxWidth="sm">
        <Card sx={{ marginTop: 1 }}>
          <CardContent>
            <form onSubmit={handleSubmit}>
              <Typography variant="h5">Create New Post</Typography>
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
                label="Body"
                fullWidth
                variant="standard"
                multiline
                value={body}
                onChange={handleBodyChange}
              />
              <Button
                type="submit"
                variant="contained"
                fullWidth
                sx={{ marginTop: 1 }}
              >
                {isLoading ? <CircularProgress /> : <span>Create</span>}
              </Button>
            </form>
          </CardContent>
        </Card>
      </Container>
    </BaseLayout>
  );
}
