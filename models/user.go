package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("user can't have an ID")
	}

	user.ID = nextId
	nextId++
	users = append(users, &user)
	return user, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("user with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) (User, error) {
	userIdx := -1

	for i, u := range users {
		if u.ID == id {
			userIdx = i

			break
		}
	}

	if userIdx > -1 {
		user := users[userIdx]
		users = append(users[:userIdx], users[(userIdx+1):]...)

		return *user, nil
	}

	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}
