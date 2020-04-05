package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error){
	if u.ID !=0 {
		return User{}, errors.New("User must not include Id or it must not be 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
			
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)

}


func RemoveUserById(id int) error {
	for i, u := range users {
		if id == u.ID {
			users = append(users[:i], users[i+1:]...)
			return  nil
			
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)

}
