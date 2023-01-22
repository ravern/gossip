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
  Typography,
} from "@mui/material";
import React from "react";

import useLikePostMutation from "src/api/mutations/likePost";
import useCurrentUserQuery from "src/api/queries/currentUser";

import type { PostData } from "src/api/models";

export interface PostListItemProps {
  post: PostData;
}

export default function PostListItem({ post }: PostListItemProps) {
  const { data: currentUser } = useCurrentUserQuery();

  const { mutateAsync: likePost } = useLikePostMutation();

  const isLiked = post.likes.some(({ user_id }) => currentUser?.id === user_id);

  const handleLikeClick = () => {
    likePost({ postId: post.id, isLiked: !isLiked })
      .then(console.log)
      .catch(console.error);
  };

  return (
    <Card sx={{ marginTop: 1 }}>
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {post.title}
        </Typography>
        {post.body != null && (
          <Typography variant="body2" color="text.secondary">
            {post.body}
          </Typography>
        )}
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
            3
          </Typography>
        </Button>
      </CardActions>
    </Card>
  );
}
