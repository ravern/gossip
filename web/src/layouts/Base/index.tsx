import {
  AppBar,
  Box,
  Button,
  Container,
  Fab,
  Toolbar,
  Typography,
} from "@mui/material";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import useCurrentUserQuery from "src/api/queries/currentUser";
import LoginDialog from "src/components/LoginDialog";

export interface BaseLayoutProps {
  children: React.ReactNode;
}

export default function BaseLayout({ children }: BaseLayoutProps) {
  const navigate = useNavigate();
  const { data: currentUser } = useCurrentUserQuery();

  const [isLoginOpen, setIsLoginOpen] = useState(false);

  const handleCreateClick = () => {
    navigate("/posts/new");
  };

  const handleLoginClick = () => {
    setIsLoginOpen(true);
  };

  const handleLoginClose = () => {
    setIsLoginOpen(false);
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
                <Typography variant="body1">{currentUser.handle}</Typography>
              </>
            ) : (
              <>
                <Box sx={{ flexGrow: 1 }} />
                <Button color="inherit" onClick={handleLoginClick}>
                  Login
                </Button>
                <LoginDialog isOpen={isLoginOpen} onClose={handleLoginClose} />
                <Button color="inherit">Register</Button>
              </>
            )}
          </Toolbar>
        </Container>
      </AppBar>
      {children}
    </>
  );
}
