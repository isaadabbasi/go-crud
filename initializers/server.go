package initializers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/isaadabbasi/go_crud/controllers"

	"github.com/isaadabbasi/go_crud/repositories"
)

type callback func(string)

// InitServer - Initialize API server using mux router
func InitServer(port string, cb callback) {
	router := mux.NewRouter()

	// Initialize the repository
	repositories.InitUserRepo()

	router.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PATCH")

	cb(":" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
