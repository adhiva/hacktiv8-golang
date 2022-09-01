package session4

import (
	"encoding/json"
	"fmt"
	IUserService "hacktiv8/golang-basic/session-3"
	"net/http"
)

type UserHandler struct {
	UserService IUserService.IUserService
}

func NewUserHandler(userService IUserService.IUserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (nuh *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users := nuh.UserService.GetUser()
	jData, err := json.Marshal(users)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func (nuh *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Route tidak ditemukan")
		return
	}

	name := r.FormValue("name")
	nuh.UserService.Register(&IUserService.User{Name: name})
	msg := fmt.Sprintf("%s Berhasil ditambahkan", name)
	fmt.Fprint(w, msg)

}
