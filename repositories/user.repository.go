package repositories

import (
	"sort"
	"strconv"

	"github.com/isaadabbasi/go_crud/entities"
)

var users []entities.User = []entities.User{}

// InitUserRepo - Initializes the repo and adds some mock Users
func InitUserRepo() {
	users = append(users, entities.User{
		ID:       "1001",
		Email:    "person_one@example.com",
		Password: "abcd1234",
		Role:     "admin",
		Username: "personone",
	})
	users = append(users, entities.User{
		ID:       "1002",
		Email:    "person_two@example.com",
		Password: "abcd1234",
		Role:     "user",
		Username: "persontwo",
	})
	users = append(users, entities.User{
		ID:       "1003",
		Email:    "person_three@example.com",
		Password: "abcd1234",
		Role:     "user",
		Username: "personthree",
	})
	users = append(users, entities.User{
		ID:       "1004",
		Email:    "person_four@example.com",
		Password: "abcd1234",
		Role:     "user",
		Username: "personfour",
	})
}

// GetUsers - return a list of Users
func GetUsers() []entities.User {
	return users
}

// GetUser - Returns a user if available
func GetUser(id string) (entities.User, string) {
	var rUser entities.User
	for _, user := range users {
		if user.ID == id {
			rUser = user
		}
	}
	if rUser.ID == "" {
		return rUser, "Not Found"
	}
	return rUser, ""
}

// DeleteUser - Delete User from data layer
func DeleteUser(id string) (deleted bool) {
	for idx, user := range users {
		if user.ID == id {
			deleted = true
			users = append(users[:idx], users[idx+1:]...)
		}
	}
	return
}

// CreateUser - Create User
func CreateUser(user *entities.User) bool {
	users = append(users, *user)
	return true
}

// UpdateUser - Create User
func UpdateUser(user *entities.User) (updated bool) {
	for idx, usr := range users {
		if usr.ID == user.ID {
			updated = true
			users = append(users[:idx], users[idx+1:]...)
			users = append(users, *user)
		}
	}
	sort.SliceStable(users, func(i, j int) bool {
		prevID, _ := strconv.Atoi(users[i].ID)
		nextID, _ := strconv.Atoi(users[j].ID)
		return prevID < nextID
	})
	return
}
