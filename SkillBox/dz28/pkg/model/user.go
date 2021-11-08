package model

import "fmt"

type User struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Age     string   `json:"age"`
	Friends []string `json:"friends"`
}

func (s *User) SetName(name string) {
	s.Name = name
}
func (s *User) SetAge(age string) {
	s.Age = age
}
func (s *User) AddFriend(friendId string) {
	s.Friends = append(s.Friends, friendId)
}
func (s *User) GetName() string {
	return s.Name
}
func (s *User) GetAge() string {
	return s.Age
}
func (s *User) GetFriends() []string {
	fmt.Printf("*** %+v\n", s.Friends)
	return s.Friends
}
