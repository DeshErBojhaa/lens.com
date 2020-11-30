package models

import (
	"fmt"
	"testing"
)

func testingUserService() (*UserService, error) {
	const (
		host   = "localhost"
		port   = 5432
		user   = "tamjidahmed"
		dbname = "lens_test"
	)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}
	us.DestructiveReset()
	fmt.Println("Success")
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name:  "tamjid",
		Email: "tamjid@gmail.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Expected ID > 0, Received %d", user.ID)
	}
}
