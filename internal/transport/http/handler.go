package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rest_api_course/internal/comment"
	"strconv"
)

// Handler - stores pointer to service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

//Response - an object to store response from API
type Response struct {
	Message string
}

// NewHandler - Handler constructor
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes -sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm Alive\tLOL")
	})
}

// GetAllComments - retrieves all comments from the comments service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Error while GetAllComments()")
	}
	fmt.Fprintf(w, "%+v", comments)
}

// GetComment - retrieves comments by their ID from the comments service
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Fprintf(w, "Error while ParseUInt(id)")
	}

	comment, err := h.Service.GetComment(uint32(i))
	if err != nil {
		fmt.Fprintf(w, "Error while GetComment(id)")
	}
	fmt.Fprintf(w, "%+v", comment)
}

// PostComment - adds a new comment to the comments service
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(w, "Error while PostComment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

// UpdateComment - updates a comment by ID with new comment info
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "Error while UpdateComment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

// DeleteComment - deletes a comment from the comments service by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Fprintf(w, "Error while ParseUInt(id)")
	}

	err = h.Service.DeleteComment(uint32(i))
	if err != nil {
		fmt.Fprintf(w, "Error while DeleteComment(id)")
	}
	fmt.Fprintf(w, "DELETED")
}
