package main

import (
	"fmt"
	"github.com/santiago-rodrig/GoWebServer/models"
)

func main() {
	user := models.User{
		Id:        2,
		FirstName: "Santiago",
		LastName:  "Rodriguez",
	}

	fmt.Println(user)
}
