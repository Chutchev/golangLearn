package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"module31/pkg/friends"
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
	fmt.Println(users.Store["1"])
}
func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/makeFriends", MakeFriends)
	r.Get("/friends/{user_id}", GetFriends)

	fmt.Println("Запускаю сервер!")
	http.ListenAndServe(":8080", r)

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
	err = users.Dump(abs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось сохранить изменения в БД"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v и %v теперь друзья!", userFirst.Name, userSecond.Name)))
}
