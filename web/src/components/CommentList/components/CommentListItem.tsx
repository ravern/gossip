import {
  DeleteOutline as DeleteOutlinedIcon,
  ThumbUp as ThumbUpIcon,
  ThumbUpOutlined as ThumbUpOutlinedIcon,
} from "@mui/icons-material";
import { Box, Button, Divider, Typography } from "@mui/material";
import { DateTime } from "luxon";
import React from "react";

import useDeleteCommentMutation from "src/api/mutations/deleteComment";
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
  const { mutateAsync: deleteComment } = useDeleteCommentMutation();

  const isLiked = comment.likes.some(
    ({ user_id }) => currentUser?.id === user_id
  );
  const createdAt = DateTime.fromISO(comment.created_at).toLocaleString(
    DateTime.DATETIME_MED
  );

  const handleLikeClick = () => {
    if (currentUser != null) {
      likeComment({ postId, commentId: comment.id, isLiked: !isLiked })
        .then(console.log)
        .catch(console.error);
    }
  };

  const handleDeleteClick = (commentId: string) => () => {
    if (currentUser != null && confirm("Are you sure?")) {
      deleteComment({ postId, commentId })
        .then(console.log)
        .catch(console.error);
    }
  };

  return (
    <>
      <Typography variant="body2" color="text.secondary" marginTop={1}>
        {comment.author.handle}
      </Typography>
      <Typography variant="body1" color="text.secondary" marginTop={1}>
        {comment.body}
      </Typography>
      <Typography variant="body2" color="text.secondary" marginTop={1}>
        {createdAt}
      </Typography>
      <Button
        size="small"
        onClick={handleLikeClick}
        disabled={currentUser == null}
        sx={{ marginTop: 1 }}
      >
        {isLiked ? <ThumbUpIcon /> : <ThumbUpOutlinedIcon />}
        <Typography variant="body1" marginLeft={1}>
          {comment.likes.length}
        </Typography>
      </Button>
      {currentUser != null &&
        (currentUser.role === "moderator" ||
          currentUser.role === "admin" ||
          currentUser.id === comment.author.id) && (
          <Button
            size="small"
            onClick={handleDeleteClick(comment.id)}
            disabled={currentUser == null}
            sx={{ marginTop: 1 }}
          >
            <DeleteOutlinedIcon color="error" />
          </Button>
        )}
      <Divider sx={{ marginTop: 2 }} />
    </>
  );
}
