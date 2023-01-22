import { AppBar, Button, Container, Toolbar, Typography } from "@mui/material";
import React, { useState } from "react";

import useCurrentUserQuery from "src/api/queries/currentUser";
import LoginDialog from "src/components/LoginDialog";

export interface BaseLayoutProps {
  children: React.ReactNode;
}

export default function BaseLayout({ children }: BaseLayoutProps) {
  const { data: currentUser } = useCurrentUserQuery();

  const [isLoginOpen, setIsLoginOpen] = useState(false);

  const handleLoginClick = () => {
    setIsLoginOpen(true);
  };

  const handleLoginClose = () => {
    setIsLoginOpen(false);
  };

  return (
    <>
      <AppBar position="sticky">
        <Container maxWidth="md">
          <Toolbar disableGutters>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              Gossip
            </Typography>
            {currentUser != null ? (
              <>{currentUser.handle}</>
            ) : (
              <>
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
