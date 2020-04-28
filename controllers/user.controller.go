package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/isaadabbasi/go_crud/entities"

	"github.com/gorilla/mux"
	"github.com/isaadabbasi/go_crud/repositories"
)

// GetUsers - Return list of Users Model
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := repositories.GetUsers()
	json.NewEncoder(w).Encode(users)
}

// GetUser - Get User by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, msg := repositories.GetUser(params["id"])
	if msg != "Not Found" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)
	} else {
		w.Write([]byte(msg))
	}
}

// DeleteUser - Delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	deleted := repositories.DeleteUser(params["id"])
	if deleted == true {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	}
}

// CreateUser - Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)

	created := repositories.CreateUser(&user)

	if created == true {
		w.WriteHeader(201)
		w.Write([]byte("Created"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	}
}

// UpdateUser - Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	deleted := repositories.UpdateUser(&user)

	if deleted == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	}
}
