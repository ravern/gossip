export interface DataResponse<T> {
  data: T;
}

export interface ErrorResponse {
  error: {
    message: string;
  };
}

export interface UserData {
  id: string;
  handle: string;
  email: string;
}

export interface CurrentUserData extends UserData {
  created_at: string;
  updated_at: string;
}

export interface PostData {
  id: string;
  title: string;
  body?: string;
  likes: PostLikeData[];
  tags: string[];
  created_at: string;
}

export interface PostLikeData {
  post_id: string;
  user_id: string;
  created_at: string;
}
