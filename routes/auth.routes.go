package routes

import (
	"github.com/gorilla/mux"
	"github.com/isaadabbasi/go_crud/controllers"
)

// RegisterAuthRoutes - Handle Signin, Signup, Forget password, etc.
func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/signin", controllers.HandleSignin).Methods("POST")
	router.HandleFunc("/api/signup", controllers.HandleSignup).Methods("POST")
}
