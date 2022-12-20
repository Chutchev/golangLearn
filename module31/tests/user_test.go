package tests

import (
	"fmt"
	"module31/pkg/user"
	"testing"
)

func TestUserToString(t *testing.T) {
	username := "Ivan"
	age := 23
	friends := make([]string, 0)

	message := fmt.Sprintf("Name: %s, Age: %d :)\n", username, age)
	user := user.User{
		Name:    username,
		Age:     age,
		Friends: friends}

	if message != user.ToString() {
		t.Logf("%v != %v", message, user.ToString())
		t.Fatalf("%v != %v", message, user.ToString())
	}
}

func TestAddFriend(t *testing.T) {
	firstUserName := "Ivan"
	secondUserName := "Andrey"
	firstUser := user.User{
		Name:    firstUserName,
		Age:     23,
		Friends: make([]string, 0)}
	secondUser := user.User{
		Name:    secondUserName,
		Age:     24,
		Friends: make([]string, 0)}

	firstUser.AddFriend(&secondUser)

	if len(firstUser.Friends) != 1 {
		t.Logf("Длина списка первого юзера друзей = %d", len(firstUser.Friends))
		t.Fatalf("Длина списка первого юзера друзей = %d", len(firstUser.Friends))
	}

	if len(secondUser.Friends) != 1 {
		t.Logf("Длина списка второго юзера друзей = %d", len(secondUser.Friends))
		t.Fatalf("Длина списка второго юзера друзей = %d", len(secondUser.Friends))
	}

	if firstUser.Friends[0] != secondUserName {
		t.Logf("Пользователи: %v и %v не друзья. Друг первого пользователя: %v", firstUser.Name, secondUser, firstUser.Friends[0])
		t.Errorf("Пользователи: %v и %v не друзья.", firstUser.Name, secondUser)
	}

	if secondUser.Friends[0] != firstUserName {
		t.Logf("Пользователи: %v и %v не друзья. Друг второго пользователя: %v", firstUser.Name, secondUser, secondUser.Friends[0])
		t.Errorf("Пользователи: %v и %v не друзья.", firstUser.Name, secondUser)
	}
}

func TestAddUser(t *testing.T) {
	var users user.Users
	users.Store = make(map[string]*user.User, 0)
	firstUserName := "Ivan"
	firstUser := user.User{
		Name:    firstUserName,
		Age:     23,
		Friends: make([]string, 0)}

	users.AddUser(&firstUser)
	if len(users.Store) != 1 {
		t.Log("Пользователь не добавился в хранилище")
		t.Fatal("Пользователь не добавился в хранилище")
	}
}

func TestGetUserById(t *testing.T) {
	userID := 1
	var users user.Users
	users.Store = make(map[string]*user.User, 0)
	firstUserName := "Ivan"
	secondUserName := "Andrey"
	firstUser := user.User{
		Name:    firstUserName,
		Age:     23,
		Friends: make([]string, 0)}
	secondUser := user.User{
		Name:    secondUserName,
		Age:     24,
		Friends: make([]string, 0)}

	users.AddUser(&firstUser)
	users.AddUser(&secondUser)

	resultUser, _ := users.GetUserById(userID)
	if resultUser != &firstUser {
		t.Logf("Пользователь с id: %d не равен правильному", userID)
		t.Fatalf("Пользователь с id: %d не равен правильному", userID)
	}

}
