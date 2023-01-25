import { Card, CardContent, List } from "@mui/material";
import React from "react";

import { CommentData } from "src/api/models";
import useCurrentUserQuery from "src/api/queries/currentUser";

import CommentInput from "./components/CommentInput";
import CommentListItem from "./components/CommentListItem";

export interface CommentListProps {
  postId: string;
  comments: CommentData[];
}

export default function CommentList({ postId, comments }: CommentListProps) {
  const { data: currentUser } = useCurrentUserQuery();

  return (
    <Card sx={{ marginTop: 2 }}>
      <CardContent>
        <List>
          {comments.length > 0 ? (
            comments.map((comment) => {
              return (
                <CommentListItem
                  key={comment.id}
                  postId={postId}
                  comment={comment}
                />
              );
            })
          ) : (
            <>There are no comments yet.</>
          )}
        </List>
        {currentUser != null && <CommentInput postId={postId} />}
      </CardContent>
    </Card>
  );
}
