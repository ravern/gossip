package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	validate "github.com/go-playground/validator/v10"
	"github.com/ravern/gossip/v2/internal/database"
	"github.com/ravern/gossip/v2/internal/middleware"
	"github.com/ravern/gossip/v2/internal/response"
	"gorm.io/gorm"
)

type createPostBody struct {
	Title string   `json:"title" validate:"required"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags" validate:"required"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)

	var b createPostBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if err := validate.New().Struct(&b); err != nil {
		logger.Warn().Err(err).Msg("bad request")
		response.JSONError(w, err, http.StatusBadRequest)
		return
	}

	post := database.Post{
		Title: b.Title,
		Body:  b.Body,
		Tags:  b.Tags,
	}

	if result := db.Create(&post); result.Error != nil {
		logger.Error().Err(result.Error).Msg("failed to create post")
		response.InternalServerError(w)
		return
	}

	response.WriteJSON(w, post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := middleware.GetDB(r)
	logger := middleware.GetLogger(r)

	var posts []database.Post
	if result := db.Find(&posts); result.Error != nil {
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

	response.WriteJSON(w, &getPostPayload{Post: post})
}
