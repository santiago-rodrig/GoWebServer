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

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	user.Id = nextId
	nextId++
	users = append(users, &user)
	return user, nil
}
