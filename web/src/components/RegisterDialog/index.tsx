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
import useRegisterMutation from "src/api/mutations/register";

export interface RegisterDialogProps {
  isOpen: boolean;
  onClose: () => void;
}

export default function RegisterDialog({
  isOpen,
  onClose,
}: RegisterDialogProps) {
  const { mutateAsync: register } = useRegisterMutation();

  const [handle, setHandle] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [avatarURL, setAvatarURL] = useState(null);

  const handleHandleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setHandle(e.target.value);
  };

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    register({
      handle,
      email,
      password,
      avatarURL,
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
        <DialogTitle>Create a new account!</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Start your gossiping here today :)
          </DialogContentText>
          <TextField
            margin="dense"
            autoFocus
            label="Handle"
            fullWidth
            variant="standard"
            value={handle}
            onChange={handleHandleChange}
          />
          <TextField
            margin="dense"
            label="Email"
            fullWidth
            variant="standard"
            value={email}
            onChange={handleEmailChange}
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
          <Button type="submit">Register</Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}
