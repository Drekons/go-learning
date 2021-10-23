package user

import (
	"dz27/pkg/model"
	"errors"
	"strconv"
)

type Storage map[string]*model.User

var index = 0

func NewUserStorage() Storage {
	return make(map[string]*model.User)
}

func (us Storage) Put(u *model.User) string {
	index := us.GenerateIndex()
	us[index] = u
	return index
}

func (us Storage) GenerateIndex() string {
	index++
	return strconv.Itoa(index)
}

func (us Storage) GetByIndex(index string) (*model.User, error) {
	user := us[index]

	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (us *Storage) DeleteByIndex(index string) {
	newStorage := NewUserStorage()
	for key, u := range *us {
		if key == index {
			continue
		}
		if len(u.GetFriends()) == 0 {
			continue
		}
		for k, idx := range u.Friends {
			if idx == index {
				u.Friends = append(u.Friends[:k], u.Friends[k+1:]...)
			}
		}
		newStorage[key] = u
	}
	*us = newStorage
}
