package router

import (
	"net/http"

	"github.com/Akashkarmokar/golang_server_setup/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	// Create news Route
	r.HandleFunc("POST /news", handler.PostNew())

	// Get All News
	r.HandleFunc("GET /news", handler.GetAllNews())

	// Get news By ID
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID())

	// Update news by ID.
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID())

	// Delete news by ID.
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID())
	return r
}
