package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler - stores pointer to service
type Handler struct {
	Router *mux.Router
}

// NewHandler - Handler constructor
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes -sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up the Routes")

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm Alive\tLOL")
	})
}
