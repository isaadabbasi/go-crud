package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/isaadabbasi/go_crud/repositories"

	"github.com/isaadabbasi/go_crud/entities"
)

// HandleSignup - Signup/register controller
func HandleSignup(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)

	err := repositories.CreateUser(&user)
	if err != nil {
		errText := err.Error()
		w.Write([]byte(errText))
		return
	}

	w.WriteHeader(201)
}

// HandleSignin - Signin/ Login Controller
func HandleSignin(w http.ResponseWriter, r *http.Request) {
	var creds entities.SigninCredentials
	json.NewDecoder(r.Body).Decode(&creds)

	auth, err := repositories.HandleSignin(&creds)

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Invalid Credentials"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(auth)
}
