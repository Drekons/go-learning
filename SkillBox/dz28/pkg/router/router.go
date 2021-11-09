package router

import (
	"dz28/pkg/action"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init(router *chi.Mux) {
	router.Use(middleware.AllowContentEncoding("Content-Type: application/json; charset=utf-8"))

	router.Post("/create", action.UserCreate)
	router.Post("/make_friends", action.FriendsMake)
	router.Delete("/user", action.UserDelete)
	router.Get("/friends/{user_id}", action.FriendsGet)
	router.Put("/{user_id}", action.UserUpdate)
}
