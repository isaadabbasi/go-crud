package initializers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isaadabbasi/go_crud/routes"
)

type callback func(string)

// InitServer - Initialize API server using mux router
func InitServer(port string, cb callback) {
	router := mux.NewRouter()

	router.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)

	cb(":" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
