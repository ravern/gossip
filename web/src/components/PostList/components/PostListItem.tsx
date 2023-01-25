import {
  CommentOutlined as CommentIcon,
  DeleteOutlined as DeleteOutlinedIcon,
  ThumbUp as ThumbUpIcon,
  ThumbUpOutlined as ThumbUpOutlinedIcon,
} from "@mui/icons-material";
import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Typography,
} from "@mui/material";
import { DateTime } from "luxon";
import React from "react";
import { Link } from "react-router-dom";

import useDeletePostMutation from "src/api/mutations/deletePost";
import useLikePostMutation from "src/api/mutations/likePost";
import useCurrentUserQuery from "src/api/queries/currentUser";

import type { PostData } from "src/api/models";

export interface PostListItemProps {
  post: PostData;
}

export default function PostListItem({ post }: PostListItemProps) {
  const { data: currentUser } = useCurrentUserQuery();

  const { mutateAsync: likePost } = useLikePostMutation();
  const { mutateAsync: deletePost } = useDeletePostMutation();

  const isLiked = post.likes.some(({ user_id }) => currentUser?.id === user_id);
  const createdAt = DateTime.fromISO(post.created_at).toLocaleString(
    DateTime.DATETIME_MED
  );

  const handleLikeClick = () => {
    if (currentUser != null) {
      likePost({ postId: post.id, isLiked: !isLiked })
        .then(console.log)
        .catch(console.error);
    }
  };

  const handleDeleteClick = (postId: string) => () => {
    if (currentUser != null && confirm("Are you sure?")) {
      deletePost(postId).then(console.log).catch(console.error);
    }
  };

  return (
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
            {post.body.length > 80 ? post.body.slice(0, 80) + "..." : post.body}
          </Typography>
        )}
        <Typography variant="body2" color="text.secondary" marginTop={1}>
          {`by ${post.author.handle} on ${createdAt}`}
        </Typography>
      </CardContent>
      <CardActions>
        <Button
          size="small"
          onClick={handleLikeClick}
          disabled={currentUser == null}
        >
          {isLiked ? <ThumbUpIcon /> : <ThumbUpOutlinedIcon />}
          <Typography variant="body1" marginLeft={1}>
            {post.likes.length}
          </Typography>
        </Button>
        <Button size="small" component={Link} to={`/posts/${post.id}`}>
          <CommentIcon />
          <Typography variant="body1" marginLeft={1}>
            {post.comments.length}
          </Typography>
        </Button>
        <Box sx={{ flexGrow: 1 }} />
        {currentUser != null &&
          (currentUser.role === "moderator" ||
            currentUser.role === "admin" ||
            currentUser.id === post.author.id) && (
            <Button
              size="small"
              onClick={handleDeleteClick(post.id)}
              disabled={currentUser == null}
            >
              <DeleteOutlinedIcon color="error" />
            </Button>
          )}
      </CardActions>
    </Card>
  );
}
