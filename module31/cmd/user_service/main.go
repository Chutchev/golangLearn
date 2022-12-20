package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"module31/pkg/helpers"
	"module31/pkg/user"

	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var users user.Users

var abs, _ = filepath.Abs("../user_service/database.json")

func init() {
	data, err := helpers.ReadFile(abs)
	if err != nil {
		users.Store = make(map[string]*user.User)
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/create", Create)
	r.Get("/all", GetAll)
	r.Delete("/user", DeleteUser)
	r.Put("/{user_id}", ChangeAge)

	fmt.Println("Запускаю сервер!")
	http.ListenAndServe(":8081", r)

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
	err = users.Dump("database.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось сохранить изменения в БД"))
	}
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
	targetId := strconv.Itoa(deleteUser.TargetId)
	delete(users.Store, targetId)
	message := fmt.Sprintf("Deleted user with Name: %v", user.Name)
	w.Write([]byte(message))
	err = users.Dump("database.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось сохранить изменения в БД"))
	}
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
		users.Dump(abs)
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
