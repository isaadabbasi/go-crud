package routes

import (
	"github.com/gorilla/mux"
	"github.com/isaadabbasi/go_crud/controllers"
	"github.com/isaadabbasi/go_crud/middlewares"
)

// RegisterUserRoutes - Handle user related routes followed e.g. /api/users?/*, etc.
func RegisterUserRoutes(router *mux.Router) {

	router.HandleFunc("/api/users", middlewares.ValidateToken(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PATCH")

}
