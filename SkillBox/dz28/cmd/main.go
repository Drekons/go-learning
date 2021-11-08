package main

import (
	"dz28/pkg/helper"
	"dz28/pkg/model"
	user "dz28/pkg/storage"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var userStorage user.Storage

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "specify port to use. Defaults to 8080.")
	flag.Parse()

	userStorage = user.NewUserStorage()

	r := chi.NewRouter()
	r.Use(middleware.AllowContentEncoding("Content-Type: application/json; charset=utf-8"))

	r.Post("/create", create)
	r.Post("/make_friends", makeFriends)
	r.Delete("/user", deleteUser)
	r.Get("/friends/{user_id}", getFriends)
	r.Put("/{user_id}", updateUser)

	log.Printf("Server running on port %d ...", port)

	if err := http.ListenAndServe(":"+strconv.Itoa(port), r); err != nil {
		log.Fatalln(err)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)
	userId := chi.URLParam(r, "user_id")

	var uu model.UpdateUser
	err := helper.ReadContent(w, r, &uu)
	if err != nil {
		log.Println(err)
		return
	}

	u, err := userStorage.Find(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.SetAge(uu.NewAge)
	userStorage.Update(u)

	w.WriteHeader(http.StatusOK)
	if err = helper.WriteJson("{\"message\": \"Возраст пользователя успешно обновлён\"}", w); err != nil {
		log.Println(err)
		return
	}
}

func getFriends(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)
	userId := chi.URLParam(r, "user_id")

	u, err := userStorage.Find(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ""
	for _, friendId := range u.GetFriends() {
		f, err := userStorage.Find(friendId)
		if err != nil {
			continue
		}

		if len(response) > 0 {
			response += ",\n"
		}
		response += fmt.Sprintf("{\"ID\": \"%s\",\"name\": \"%s\"}", friendId, f.GetName())
	}

	log.Printf("Getting the user's friends id %s\n", userId)

	w.WriteHeader(http.StatusOK)
	if err = helper.WriteJson(fmt.Sprintf("[%s]", response), w); err != nil {
		log.Println(err)
		return
	}
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)

	var du model.DeleteUser
	err := helper.ReadContent(w, r, &du)
	if err != nil {
		log.Println(err)
		return
	}

	tUser, err := userStorage.Find(du.Target)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := tUser.GetName()
	userStorage.DeleteByIndex(du.Target)

	log.Printf("User id %s deleted\n", du.Target)

	w.WriteHeader(http.StatusOK)
	if err = helper.WriteJson(fmt.Sprintf("{\"message\": \"%s\"}", name), w); err != nil {
		log.Println(err)
		return
	}
}

func makeFriends(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)

	var mf model.MakeFriend
	err := helper.ReadContent(w, r, &mf)
	if err != nil {
		log.Println(err)
		return
	}

	sUser, err := userStorage.Find(mf.Source)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tUser, err := userStorage.Find(mf.Target)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sUser.AddFriend(mf.Target)
	tUser.AddFriend(mf.Source)

	userStorage.Update(sUser)
	userStorage.Update(tUser)

	log.Printf("User id %s and user id %s are friends now\n", mf.Source, mf.Target)

	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("%s и %s теперь друзья", sUser.GetName(), tUser.GetName())
	if err = helper.WriteJson(fmt.Sprintf("{\"message\": \"%s\"}", message), w); err != nil {
		log.Println(err)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)

	var u model.User
	err := helper.ReadContent(w, r, &u)
	if err != nil {
		log.Println(err)
		return
	}

	index := userStorage.Put(&u)
	log.Printf("User id %s created: %+v\n", index, u)
	w.WriteHeader(http.StatusCreated)

	if err = helper.WriteJson(fmt.Sprintf("{\"ID\": \"%s\"}", index), w); err != nil {
		log.Println(err)
		return
	}
}
