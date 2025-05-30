package routes

import (
	"net/http"

	"github.com/Car-Rental/backend/user-service/internal/handler"
	"github.com/gorilla/mux"
)

// UserRoutes defines the routes for user management
func UserRoutes(router *mux.Router, userHandler *handler.UserHandler) {
	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/login", userHandler.LoginUser).Methods(http.MethodPost)
	// Add more routes as needed (e.g., get user, update user)
}
