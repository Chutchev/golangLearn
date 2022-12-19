package user

import (
	"fmt"
	"module31/pkg/helpers"
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
	return fmt.Sprintf("Name: %v, Age: %d :)\n", u.Name, u.Age)
}

type Users struct {
	Store map[int]*User
}

func (users *Users) AddUser(u *User) int {
	current_id++
	users.Store[current_id] = u
	return current_id
}

func (users *Users) GetAll() string {
	allUsers := ""
	for index, user := range users.Store {
		allUsers += fmt.Sprintf("%d. %v", index, user.ToString())
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
	user := users.Store[userId]
	if user == nil {
		return &User{}, fmt.Errorf("User with id: %d is not found", userId)
	}
	return user, nil
}

func (users *Users) GetUserIdByName(userName string) (int, error) {
	for index, user := range users.Store {
		if user.Name == userName {
			return index, nil
		}
	}
	return -1, fmt.Errorf("User with name: %v is not found", userName)
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
