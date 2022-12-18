package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"module30/pkg/friends"
	"module30/pkg/user"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var users user.Users

func init() {
	users.Store = make(map[int]*user.User)
}
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/create", Create)
	r.Post("/makeFriends", MakeFriends)
	r.Get("/all", GetAll)
	r.Get("/friends/{user_id}", GetFriends)

	r.Delete("/user", DeleteUser)

	r.Put("/{user_id}", ChangeAge)

	fmt.Println("Запускаю сервер!")
	http.ListenAndServe(":3333", r)

}

func ChangeAge(w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var changeAge user.ChangeAge
	if err := json.Unmarshal(content, &changeAge); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := users.GetUserById(user_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	user.Age = changeAge.NewAge
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Возраст пользователя успешно обновлен!"))

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var deleteUser user.DeleteUser
	if err := json.Unmarshal(content, &deleteUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := users.GetUserById(deleteUser.TargetId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	for _, friendName := range user.Friends {
		friend, _ := users.GetUserByName(friendName)
		friend.DeleteFriend(user.Name)
	}
	delete(users.Store, deleteUser.TargetId)
	message := fmt.Sprintf("Deleted user with Name: %v", user.Name)
	w.Write([]byte(message))
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	user_id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userFirst, err := users.GetUserById(user_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	friends := ""
	for _, friend := range userFirst.Friends {
		friends += friend + " "
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(friends))
}

func MakeFriends(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var f friends.Friends
	if err := json.Unmarshal(content, &f); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	userFirst, err := users.GetUserById(f.SourceId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	userSecond, err := users.GetUserById(f.TargetId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	userFirst.AddFriend(userSecond)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v и %v теперь друзья!", userFirst.Name, userSecond.Name)))
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var u user.User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		user_id := users.AddUser(&u)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("User with id: %d was created", user_id)))
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(users.GetAll()))
}
