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
  Divider,
  Typography,
} from "@mui/material";
import { DateTime } from "luxon";
import React from "react";
import { Link } from "react-router-dom";

import useLikeCommentMutation from "src/api/mutations/likeComment";
import useCurrentUserQuery from "src/api/queries/currentUser";

import type { CommentData } from "src/api/models";

export interface CommentListItemProps {
  postId: string;
  comment: CommentData;
}

export default function CommentListItem({
  postId,
  comment,
}: CommentListItemProps) {
  const { data: currentUser } = useCurrentUserQuery();

  const { mutateAsync: likeComment } = useLikeCommentMutation();

  const isLiked = comment.likes.some(
    ({ user_id }) => currentUser?.id === user_id
  );
  const createdAt = DateTime.fromISO(comment.created_at).toLocaleString(
    DateTime.DATETIME_MED
  );

  const handleLikeClick = () => {
    likeComment({ postId, commentId: comment.id, isLiked: !isLiked })
      .then(console.log)
      .catch(console.error);
  };

  return (
    <>
      <Typography variant="body2" color="text.secondary" marginTop={1}>
        {comment.author.handle}
      </Typography>
      <Typography variant="body1" color="text.secondary">
        {comment.body}
      </Typography>
      <Button size="small" onClick={handleLikeClick}>
        {isLiked ? <ThumbUpIcon /> : <ThumbUpOutlinedIcon />}
        <Typography variant="body1" marginLeft={1}>
          {comment.likes.length}
        </Typography>
      </Button>
      <Divider sx={{ marginTop: 2 }} />
    </>
  );
}
