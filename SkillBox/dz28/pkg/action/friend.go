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

func FriendsGet(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)
	userId := chi.URLParam(r, "user_id")
	storage := user.NewUserStorage()

	u, err := storage.Find(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ""
	for _, friendId := range u.GetFriends() {
		f, err := storage.Find(friendId)
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

func FriendsMake(w http.ResponseWriter, r *http.Request) {
	defer helper.BodyClose(r)

	var mf model.MakeFriend
	err := helper.ReadContent(w, r, &mf)
	if err != nil {
		log.Println(err)
		return
	}

	storage := user.NewUserStorage()

	sUser, err := storage.Find(mf.Source)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tUser, err := storage.Find(mf.Target)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sUser.AddFriend(mf.Target)
	tUser.AddFriend(mf.Source)

	storage.Update(sUser)
	storage.Update(tUser)

	log.Printf("User id %s and user id %s are friends now\n", mf.Source, mf.Target)

	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("%s и %s теперь друзья", sUser.GetName(), tUser.GetName())
	if err = helper.WriteJson(fmt.Sprintf("{\"message\": \"%s\"}", message), w); err != nil {
		log.Println(err)
		return
	}
}
