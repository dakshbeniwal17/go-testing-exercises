package exercise2

// without any 3rd party library, create a mock for UserExists
// edit exercise2_test.go to not call original UserExists but its mock

import (
	"fmt"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

func RegisterUser(user User) error {
	found := UserExists(user.Email)
	if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	// carry business logic and Register the user in the system
	log.Println(user)
	return nil
}
