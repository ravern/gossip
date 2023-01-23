import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  TextField,
} from "@mui/material";
import React, { useState } from "react";

import { LOCAL_STORAGE_KEY_ACCESS_TOKEN } from "src/api";
import useLoginMutation from "src/api/mutations/login";

export interface LoginDialogProps {
  isOpen: boolean;
  onClose: () => void;
}

export default function LoginDialog({ isOpen, onClose }: LoginDialogProps) {
  const { mutateAsync: login } = useLoginMutation();

  const [handleOrEmail, setHandleOrEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleHandleOrEmailChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setHandleOrEmail(e.target.value);
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    login({
      handleOrEmail,
      password,
    })
      .then(({ token }) => {
        localStorage.setItem(LOCAL_STORAGE_KEY_ACCESS_TOKEN, token);
        onClose();
      })
      .catch(console.error);
  };

  return (
    <Dialog fullWidth maxWidth="xs" open={isOpen} onClose={onClose}>
      <form onSubmit={handleSubmit}>
        <DialogTitle>Welcome back!</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Please enter your credentials here.
          </DialogContentText>
          <TextField
            margin="dense"
            autoFocus
            label="Handle or Email"
            fullWidth
            variant="standard"
            value={handleOrEmail}
            onChange={handleHandleOrEmailChange}
          />
          <TextField
            margin="dense"
            label="Password"
            type="password"
            fullWidth
            variant="standard"
            value={password}
            onChange={handlePasswordChange}
          />
        </DialogContent>
        <DialogActions>
          <Button type="submit">Login</Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}
