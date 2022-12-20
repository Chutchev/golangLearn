package user

import (
	"encoding/json"
	"fmt"
	"module31/pkg/helpers"
	"strconv"
)

var current_id = 0

type ChangeAge struct {
	NewAge int `json:"new_age"`
}

type DeleteUser struct {
	TargetId int `json:"target_id"`
}

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
}

func (u *User) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d :)\n", u.Name, u.Age)
}

type Users struct {
	Store map[string]*User `json:"Store"`
}

func (users *Users) AddUser(u *User) int {
	current_id++
	users.Store[strconv.Itoa(current_id)] = u
	return current_id
}

func (users *Users) GetAll() string {
	allUsers := ""
	for index, user := range users.Store {
		allUsers += fmt.Sprintf("%s. %v", index, user.ToString())
	}
	return allUsers
}

func (u *User) AddFriend(anotherUser *User) {
	u.Friends = append(u.Friends, anotherUser.Name)
	anotherUser.Friends = append(anotherUser.Friends, u.Name)

	fmt.Println(u.Friends)
	fmt.Println(anotherUser.Friends)
}

func (users *Users) GetUserById(userId int) (*User, error) {
	user := users.Store[strconv.Itoa(userId)]
	if user == nil {
		return &User{}, fmt.Errorf("User with id: %d is not found", userId)
	}
	return user, nil
}

func (users *Users) GetUserIdByName(userName string) (string, error) {
	for index, user := range users.Store {
		if user.Name == userName {
			return index, nil
		}
	}
	return "", fmt.Errorf("User with name: %v is not found", userName)
}

func (users *Users) GetUserByName(userName string) (*User, error) {
	for _, user := range users.Store {
		if user.Name == userName {
			return user, nil
		}
	}
	return &User{}, fmt.Errorf("User with name: %v is not found", userName)
}

func (user *User) DeleteFriend(userName string) {
	friends := user.Friends
	user.Friends = helpers.RemoveElementFromArray(friends, userName)
	fmt.Println(user.Friends)
}

func (users *Users) Dump(filename string) error {
	mapString, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = helpers.SaveFile(string(mapString), filename)
	if err != nil {
		return err
	}
	return nil
}
