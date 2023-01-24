import { Card, CardContent, List } from "@mui/material";
import React from "react";

import { CommentData } from "src/api/models";

import CommentInput from "./components/CommentInput";
import CommentListItem from "./components/CommentListItem";

export interface CommentListProps {
  postId: string;
  comments: CommentData[];
}

export default function CommentList({ postId, comments }: CommentListProps) {
  console.log(comments);
  return (
    <Card sx={{ marginTop: 2 }}>
      <CardContent>
        <List>
          {comments.map((comment) => {
            return (
              <CommentListItem
                key={comment.id}
                postId={postId}
                comment={comment}
              />
            );
          })}
        </List>
        <CommentInput postId={postId} />
      </CardContent>
    </Card>
  );
}
