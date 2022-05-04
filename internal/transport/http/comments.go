package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest_api_course/internal/comment"
	"strconv"
)

// GetAllComments - retrieves all comments from the comments service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error while GetAllComments()", err)
		return
	}

	if err = sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

// GetComment - retrieves comments by their ID from the comments service
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		sendErrorResponse(w, "Error while ParseUInt(id)", err)
		return
	}

	comment, err := h.Service.GetComment(uint32(i))
	if err != nil {
		sendErrorResponse(w, "Error while GetComment(id)", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// PostComment - adds a new comment to the comments service
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	newComment := comment.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		sendErrorResponse(w, "Failed to decode new comment", err)
		return
	}

	comment, err := h.Service.PostComment(newComment)
	if err != nil {
		sendErrorResponse(w, "Error while PostComment", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// UpdateComment - updates a comment by ID with new comment info
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	newComment := comment.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		sendErrorResponse(w, "Failed to decode new comment", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	commentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		sendErrorResponse(w, "Error while ParseUInt(id)", err)
		return
	}

	comment, err := h.Service.UpdateComment(uint32(commentId), newComment)
	if err != nil {
		sendErrorResponse(w, "Error while UpdateComment", err)
		return
	}

	if err = sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes a comment from the comments service by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		sendErrorResponse(w, "Error while ParseUInt(id)", err)
		return
	}

	err = h.Service.DeleteComment(uint32(i))
	if err != nil {
		sendErrorResponse(w, "Error while DeleteComment(id)", err)
		return
	}

	if err = sendOkResponse(w, Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}
