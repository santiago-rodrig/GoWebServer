package main

import (
	"github.com/santiago-rodrig/GoWebServer/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
