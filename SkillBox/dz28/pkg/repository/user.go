package repository

import (
	"dz28/pkg/model"
	"fmt"
	"log"
	"strings"
)

const table string = "users"

type User struct {
}

func GetUserRepo() User {
	var userRepo User
	return userRepo
}

func (u User) Add(user *model.User) int64 {
	id, err := GetSqlLite().Insert(
		table,
		"name, age, friends",
		fmt.Sprintf(
			"'%s', '%s', '%s'",
			user.Name,
			user.Age,
			strings.Join(user.Friends, ","),
		),
	)

	if err != nil {
		return 0
	}

	return id
}

func (u User) Find(id string) *model.User {
	row := GetSqlLite().SelectRow(table, fmt.Sprintf("id = %s", id))
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}

	var user model.User
	var friends string

	if err := row.Scan(nil, &user.Name, &user.Age, &friends); err != nil {
		log.Println(err)
		return nil
	}

	user.Friends = strings.Split(friends, ",")

	return &user
}

func (u User) Delete(id string) {
	rows, err := GetSqlLite().FindInSet(table, "friends", id)

	if err != nil {
		for rows.Next() {
			var fId, friends string
			err := rows.Scan(&id, nil, nil, &friends)
			if err != nil {
				log.Println(err)
				continue
			}

			friendsArr := strings.Split(friends, ",")
			for k, f := range friendsArr {
				if f == id {
					friendsArr = append(friendsArr[:k], friendsArr[k+1:]...)
				}
			}

			_, err = GetSqlLite().Update(
				table,
				fId,
				map[string]string{"friends": strings.Join(friendsArr, ",")},
			)
			if err != nil {
				log.Println(err)
			}
		}
	}

	_, err = GetSqlLite().DeleteById(table, id)
	if err != nil {
		log.Println(err)
	}
}
