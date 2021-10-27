package user

import (
	"dz28/pkg/model"
	"dz28/pkg/repository"
	"errors"
	"strconv"
)

type Storage struct {
	repo repository.User
}

func NewUserStorage() Storage {
	var storage Storage
	storage.repo = repository.GetUserRepo()
	return storage
}

func (us Storage) Put(u *model.User) string {
	id := us.repo.Add(u)
	return strconv.FormatInt(id, 10)
}

func (us Storage) Find(id string) (*model.User, error) {
	user := us.repo.Find(id)

	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (us Storage) DeleteByIndex(id string) {
	us.repo.Delete(id)
}
