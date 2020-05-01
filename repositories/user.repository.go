package repositories

import (
	"errors"
	"sort"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/isaadabbasi/go_crud/entities"
	"github.com/isaadabbasi/go_crud/utils"
)

var users []entities.User = []entities.User{}

// Lower case init functions are builtin-methods of module
// run automagically on module load
func init() {
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

var idCursor uint16 = 1004

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
func DeleteUser(id string) bool {
	var deleted bool = false
	for idx, user := range users {
		if user.ID == id {
			deleted = true
			users = append(users[:idx], users[idx+1:]...)
		}
	}
	return deleted
}

// CreateUser - Create User
func CreateUser(user *entities.User) error {
	u := *user
	idCursor = idCursor + 1
	hash, err := utils.GetHashedPassword(u.Password)

	if err != nil {
		return errors.New("Unable to hash")
	}
	u.Password = hash
	u.ID = strconv.FormatUint(uint64(idCursor), 10)
	users = append(users, u)
	return nil
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

// HandleSignin - Handle user login
func HandleSignin(creds *entities.SigninCredentials) (*entities.AuthObject, error) {
	var user entities.User
	var auth entities.AuthObject
	for _, usr := range users {
		if usr.Username == creds.Username {
			user = usr
		}
	}
	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if compareErr != nil || user.ID == "" {
		return &auth, errors.New("Invalid Credentials")
	}
	token, err := utils.GenerateToken()
	if err != nil {
		return &auth, errors.New("Something went wrong")
	}
	auth.Token = token
	auth.User = user
	auth.User.Password = "[PROTECTED]"
	return &auth, nil
}
