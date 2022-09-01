package session4

import (
	"fmt"
	session3 "hacktiv8/golang-basic/session-3"
	"net/http"
)

const PORT = ":3000"

func WebServer() {
	db := []*session3.User{}
	userService := session3.NewUserService(db)
	userHandler := NewUserHandler(userService)

	http.HandleFunc("/", hello)
	http.HandleFunc("/register", userHandler.RegisterUserHandler)
	http.HandleFunc("/users", userHandler.GetUserHandler)
	http.ListenAndServe(PORT, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World, it's time to Go!"
	fmt.Fprint(w, msg)
}
