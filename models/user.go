package models

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)
