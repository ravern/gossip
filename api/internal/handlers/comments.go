package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/ravern/gossip/v2/internal/database"
	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/ravern/gossip/v2/internal/request"
	"github.com/ravern/gossip/v2/internal/response"
	"gorm.io/gorm"
)

type CreateCommentBody struct {
	Body string `json:"body" validate:"required"`
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	postID := chi.URLParam(r, "postID")

	var b CreateCommentBody
	if err := request.DecodeBody(r, &b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	var post database.Post
	if result := db.Where("id = ?", postID).First(&post); result.Error != nil {
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

	comment := database.Comment{
		Body:     b.Body,
		AuthorID: user.ID,
		PostID:   uuid.MustParse(postID),
	}

	if result := db.Create(&comment); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to create comment")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, comment)
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)

	postID := chi.URLParam(r, "postID")

	var comments []database.Comment
	if result := db.Where("post_id = ?", postID).Find(&comments); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to fetch comments")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, comments)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)

	postID := chi.URLParam(r, "postID")
	id := chi.URLParam(r, "id")

	var comment database.Comment
	if result := db.Where("post_id = ? AND id = ?", postID, id).First(&comment); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("comment not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch comment")
			response.InternalServerError(w)
		}
		return
	}

	response.WriteJSON(w, comment)
}

type UpdateCommentBody struct {
	Body *string `json:"body"`
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	postID := chi.URLParam(r, "postID")
	id := chi.URLParam(r, "id")

	var b UpdateCommentBody
	if err := request.DecodeBody(r, &b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	var comment database.Comment
	if result := db.Where("post_id = ? AND id = ?", postID, id).First(&comment); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("comment not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch comment")
			response.InternalServerError(w)
		}
		return
	}

	if user.Role == database.NormalRole && comment.AuthorID != user.ID {
		err := errors.New("you can only update your own comment")
		logger.Warn().Err(err).Msg("unauthorized")
		response.JSONError(w, err, http.StatusUnauthorized)
		return
	}

	if b.Body != nil {
		comment.Body = *b.Body
	}

	db.Updates(comment)

	response.WriteJSON(w, comment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	postID := chi.URLParam(r, "postID")
	id := chi.URLParam(r, "id")

	var comment database.Comment
	if result := db.Where("post_id = ? AND id = ?", postID, id).First(&comment); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("comment not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch comment")
			response.InternalServerError(w)
		}
		return
	}

	if user.Role == database.NormalRole && comment.AuthorID != user.ID {
		err := errors.New("you can only delete your own comment")
		logger.Warn().Err(err).Msg("unauthorized")
		response.JSONError(w, err, http.StatusUnauthorized)
		return
	}

	db.Delete(&comment)

	response.WriteJSON(w, comment)
}

type LikeCommentBody struct {
	IsLiked bool `json:"is_liked"`
}

func LikeComment(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)
	user := middleware.GetUser(r)

	postID := chi.URLParam(r, "postID")
	id := chi.URLParam(r, "id")

	var b LikeCommentBody
	if err := request.DecodeBody(r, &b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	var comment database.Comment
	if result := db.Where("id = ? AND post_id = ?", id, postID).First(&comment); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("comment not found")
			logger.Warn().Err(err).Msg("not found")
			response.JSONError(w, err, http.StatusNotFound)
		} else {
			logger.Error().Err(result.Error).Msg("failed to fetch comment")
			response.InternalServerError(w)
		}
		return
	}

	if b.IsLiked {
		commentLike := database.CommentLike{
			UserID:    user.ID,
			CommentID: comment.ID,
		}

		if result := db.Create(&commentLike); result.Error != nil {
			logger.Error().Err(result.Error).Msg("failed to like post")
			response.InternalServerError(w)
			return
		}
	} else {
		var commentLike database.CommentLike
		if result := db.Where("user_id = ? AND comment_id = ?", user.ID, comment.ID).First(&commentLike); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				err := errors.New("comment_like not found")
				logger.Warn().Err(err).Msg("not found")
				response.JSONError(w, err, http.StatusNotFound)
			} else {
				logger.Error().Err(result.Error).Msg("failed to fetch comment_like")
				response.InternalServerError(w)
			}
			return
		}

		db.Unscoped().Delete(&commentLike)
	}

	if result := db.Where("id = ?", id).Preload("CommentLikes").First(&comment); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to fetch post")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, comment)
}
