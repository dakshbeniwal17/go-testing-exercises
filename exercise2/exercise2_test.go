package exercise2

import "testing"

func TestUserExist(t *testing.T) {
	user := User{
		Name:     "Daksh Beniwal",
		Email:    "daksh@repozitory.com",
		UserName: "daksh",
	}

	err := RegisterUser(user)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}

func TestUserDoesnotExist(t *testing.T) {
	user := User{
		Name:     "Daksh Beniwal",
		Email:    "daksh1@repozitory.com",
		UserName: "daksh",
	}

	err := RegisterUser(user)
	if err != nil {
		t.Error("Expected Register User to not throw any error but got:", err)
	}
}
