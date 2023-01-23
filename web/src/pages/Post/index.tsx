import {
  CommentOutlined as CommentIcon,
  ThumbUp as ThumbUpIcon,
  ThumbUpOutlined as ThumbUpOutlinedIcon,
} from "@mui/icons-material";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CircularProgress,
  Typography,
} from "@mui/material";
import { Container } from "@mui/system";
import { DateTime } from "luxon";
import React from "react";
import { Link, useParams } from "react-router-dom";

import useLikePostMutation from "src/api/mutations/likePost";
import useCurrentUserQuery from "src/api/queries/currentUser";
import usePostQuery from "src/api/queries/post";
import BaseLayout from "src/layouts/Base";

export default function PostPage() {
  const { id } = useParams();

  const { data: currentUser } = useCurrentUserQuery();
  const { data: post, isLoading } = usePostQuery(id);

  const { mutateAsync: likePost } = useLikePostMutation();

  const isLiked = post?.likes.some(
    ({ user_id }) => currentUser?.id === user_id
  );
  const createdAt =
    post != null
      ? DateTime.fromISO(post.created_at).toLocaleString(DateTime.DATETIME_MED)
      : null;

  const handleLikeClick = () => {
    if (post != null) {
      likePost({ postId: post.id, isLiked: !isLiked })
        .then(console.log)
        .catch(console.error);
    }
  };

  if (isLoading) {
    return <CircularProgress />;
  } else if (post != null) {
    return (
      <BaseLayout>
        <Container maxWidth="sm">
          <Card sx={{ marginTop: 1 }}>
            <CardContent>
              <Typography
                gutterBottom
                variant="h5"
                component={Link}
                to={`/posts/${post.id}`}
                sx={{ textDecoration: "none", color: "inherit" }}
              >
                {post.title}
              </Typography>
              {post.body != null && (
                <Typography variant="body1" color="text.secondary">
                  {post.body.length > 80
                    ? post.body.slice(0, 80) + "..."
                    : post.body}
                </Typography>
              )}
              <Typography variant="body2" color="text.secondary" marginTop={1}>
                {`by ${post.author.handle} on ${createdAt}`}
              </Typography>
            </CardContent>
            <CardActions>
              <Button size="small" onClick={handleLikeClick}>
                {isLiked ? <ThumbUpIcon /> : <ThumbUpOutlinedIcon />}
                <Typography variant="body1" marginLeft={1}>
                  {post.likes.length}
                </Typography>
              </Button>
              <Button size="small">
                <CommentIcon />
                <Typography variant="body1" marginLeft={1}>
                  {post.comments.length}
                </Typography>
              </Button>
            </CardActions>
          </Card>
        </Container>
      </BaseLayout>
    );
  } else {
    throw new Error("invalid state");
  }
}
