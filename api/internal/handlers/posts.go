package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ravern/gossip/v2/internal/database"
	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/ravern/gossip/v2/internal/request"
	"github.com/ravern/gossip/v2/internal/response"
	"gorm.io/gorm"
)

type postPayload struct {
	database.Post
	CommentsCount int `json:"comments_count"`
	LikesCount    int `json:"likes_count"`
}

type CreatePostBody struct {
	Title string   `json:"title" validate:"required"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags" validate:"required"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	var b CreatePostBody
	if err := request.DecodeBody(r, &b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	post := database.Post{
		Title:    b.Title,
		Body:     b.Body,
		Tags:     b.Tags,
		AuthorID: user.ID,
	}

	if result := db.Create(&post); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to create post")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, &postPayload{
		Post:          post,
		CommentsCount: 0,
		LikesCount:    0,
	})
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)

	var posts []database.Post
	if result := db.Preload("Comments").Preload("PostLikes").Preload("Author").Order("created_at DESC").Find(&posts); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to fetch posts")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, posts)
}

type getPostPayload struct {
	database.Post
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)

	id := chi.URLParam(r, "id")

	var post database.Post
	if result := db.Where("id = ?", id).Preload("Comments").Preload("PostLikes").Preload("Author").First(&post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("post not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch post")
			response.InternalServerError(w)
		}
		return
	}

	response.WriteJSON(w, &getPostPayload{Post: post})
}

type UpdatePostBody struct {
	Title *string  `json:"title"`
	Body  *string  `json:"body"`
	Tags  []string `json:"tags"`
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	id := chi.URLParam(r, "id")

	var b UpdatePostBody
	if err := request.DecodeBody(r, &b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	var post database.Post
	if result := db.Where("id = ?", id).First(&post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("post not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch post")
			response.InternalServerError(w)
		}
		return
	}

	if user.Role == database.NormalRole && post.AuthorID != user.ID {
		err := errors.New("you can only update your own post")
		logger.Warn().Err(err).Msg("unauthorized")
		response.JSONError(w, err, http.StatusUnauthorized)
		return
	}

	if b.Title != nil {
		post.Title = *b.Title
	}
	if b.Body != nil {
		post.Body = *b.Body
	}
	if b.Tags != nil {
		post.Tags = b.Tags
	}

	db.Updates(post)

	response.WriteJSON(w, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	id := chi.URLParam(r, "id")

	var post database.Post
	if result := db.Where("id = ?", id).First(&post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("post not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch post")
			response.InternalServerError(w)
		}
		return
	}

	if user.Role == database.NormalRole && post.AuthorID != user.ID {
		err := errors.New("you can only delete your own post")
		logger.Warn().Err(err).Msg("unauthorized")
		response.JSONError(w, err, http.StatusUnauthorized)
		return
	}

	db.Delete(&post)

	response.WriteJSON(w, post)
}

type LikePostBody struct {
	IsLiked bool `json:"is_liked"`
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	id := chi.URLParam(r, "id")

	var b LikePostBody
	if err := request.DecodeBody(r, &b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	var post database.Post
	if result := db.Where("id = ?", id).First(&post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("post not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch post")
			response.InternalServerError(w)
		}
		return
	}

	if b.IsLiked {
		postLike := database.PostLike{
			UserID: user.ID,
			PostID: post.ID,
		}

		if result := db.Create(&postLike); result.Error != nil {
			logger.Error().Err(result.Error).Msg("failed to like post")
			response.InternalServerError(w)
			return
		}
	} else {
		var postLike database.PostLike
		if result := db.Where("user_id = ? AND post_id = ?", user.ID, post.ID).First(&postLike); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				err := errors.New("post_like not found")
				logger.Warn().Err(err).Msg("not found")
				response.JSONError(w, err, http.StatusNotFound)
			} else {
				logger.Error().Err(result.Error).Msg("failed to fetch post_like")
				response.InternalServerError(w)
			}
			return
		}

		db.Unscoped().Delete(&postLike)
	}

	if result := db.Where("id = ?", id).Preload("Comments").Preload("PostLikes").Preload("Author").First(&post); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to fetch post")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, post)
}
