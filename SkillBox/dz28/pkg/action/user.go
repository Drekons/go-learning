package action

import (
	"dz28/pkg/helper"
	"dz28/pkg/model"
	user "dz28/pkg/storage"

	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)
	userId := chi.URLParam(r, "user_id")

	var uu model.UpdateUser
	err := helper.ReadContent(w, r, &uu)
	if err != nil {
		log.Println(err)
		return
	}

	storage := user.NewUserStorage()

	u, err := storage.Find(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.SetAge(uu.NewAge)
	storage.Update(u)

	w.WriteHeader(http.StatusOK)
	if err = helper.WriteJson("{\"message\": \"Возраст пользователя успешно обновлён\"}", w); err != nil {
		log.Println(err)
		return
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)

	var du model.DeleteUser
	err := helper.ReadContent(w, r, &du)
	if err != nil {
		log.Println(err)
		return
	}

	storage := user.NewUserStorage()

	tUser, err := storage.Find(du.Target)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := tUser.GetName()
	storage.DeleteByIndex(du.Target)

	log.Printf("User id %s deleted\n", du.Target)

	w.WriteHeader(http.StatusOK)
	if err = helper.WriteJson(fmt.Sprintf("{\"message\": \"%s\"}", name), w); err != nil {
		log.Println(err)
		return
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)

	var u model.User
	err := helper.ReadContent(w, r, &u)
	if err != nil {
		log.Println(err)
		return
	}

	storage := user.NewUserStorage()

	index := storage.Put(&u)
	log.Printf("User id %s created: %+v\n", index, u)
	w.WriteHeader(http.StatusCreated)

	if err = helper.WriteJson(fmt.Sprintf("{\"ID\": \"%s\"}", index), w); err != nil {
		log.Println(err)
		return
	}
}
