package initializers

import (
	"log"
	"net/http"

	"github.com/isaadabbasi/go_crud/controllers"

	"github.com/gorilla/mux"
)

type callback func(string)

// InitializeServer - Initialize API server using mux router
func InitializeServer(port string, cb callback) {
	router := mux.NewRouter()
	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
	port = ":" + port
	cb(port)
	log.Fatal(http.ListenAndServe(port, router))
}
