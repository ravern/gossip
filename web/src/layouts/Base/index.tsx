import {
  AppBar,
  Box,
  Button,
  Container,
  Menu,
  MenuItem,
  Toolbar,
  Typography,
} from "@mui/material";
import React, { useRef, useState } from "react";
import { useQueryClient } from "react-query";
import { Link } from "react-router-dom";

import { LOCAL_STORAGE_KEY_ACCESS_TOKEN } from "src/api";
import useCurrentUserQuery from "src/api/queries/currentUser";
import LoginDialog from "src/components/LoginDialog";
import RegisterDialog from "src/components/RegisterDialog";

export interface BaseLayoutProps {
  children: React.ReactNode;
}

export default function BaseLayout({ children }: BaseLayoutProps) {
  const queryClient = useQueryClient();

  const { data: currentUser } = useCurrentUserQuery();

  const [isLoginOpen, setIsLoginOpen] = useState(false);
  const [isRegisterOpen, setIsRegisterOpen] = useState(false);
  const [isProfileOpen, setIsProfileOpen] = useState(false);

  const profileButtonRef = useRef(null);

  const handleProfileClick = () => {
    setIsProfileOpen(true);
  };

  const handleProfileClose = () => {
    setIsProfileOpen(false);
  };

  const handleLoginClick = () => {
    setIsLoginOpen(true);
  };

  const handleLoginClose = () => {
    setIsLoginOpen(false);
  };

  const handleRegisterClick = () => {
    setIsRegisterOpen(true);
  };

  const handleRegisterClose = () => {
    setIsRegisterOpen(false);
  };

  const handleLogOutClick = () => {
    localStorage.removeItem(LOCAL_STORAGE_KEY_ACCESS_TOKEN);
    queryClient.clear();
    setIsProfileOpen(false);
  };

  return (
    <>
      <AppBar position="sticky">
        <Container maxWidth="sm">
          <Toolbar disableGutters>
            <Typography
              variant="h6"
              component={Link}
              to="/"
              sx={{ color: "inherit", textDecoration: "none" }}
            >
              Gossip
            </Typography>
            {currentUser != null ? (
              <>
                <Button
                  color="inherit"
                  sx={{ marginLeft: 2 }}
                  component={Link}
                  to="/posts/new"
                >
                  Create
                </Button>
                <Box sx={{ flexGrow: 1 }} />
                <Button
                  color="inherit"
                  ref={profileButtonRef}
                  onClick={handleProfileClick}
                >
                  {currentUser.handle}
                </Button>
                <Menu
                  anchorEl={profileButtonRef.current}
                  anchorOrigin={{
                    vertical: "bottom",
                    horizontal: "right",
                  }}
                  open={isProfileOpen}
                  onClose={handleProfileClose}
                >
                  <MenuItem onClick={handleLogOutClick} sx={{ color: "red" }}>
                    Log Out
                  </MenuItem>
                </Menu>
              </>
            ) : (
              <>
                <Box sx={{ flexGrow: 1 }} />
                <Button color="inherit" onClick={handleLoginClick}>
                  Login
                </Button>
                <LoginDialog isOpen={isLoginOpen} onClose={handleLoginClose} />
                <Button color="inherit" onClick={handleRegisterClick}>
                  Register
                </Button>
                <RegisterDialog
                  isOpen={isRegisterOpen}
                  onClose={handleRegisterClose}
                />
              </>
            )}
          </Toolbar>
        </Container>
      </AppBar>
      {children}
    </>
  );
}
