package session4

import (
	"encoding/json"
	"fmt"
	IUserService "hacktiv8/golang-basic/session-3"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (nuh *UserHandler) GetUserGinHandler(c *gin.Context) {
	users := nuh.UserService.GetUser()
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
	return
}

func (nuh *UserHandler) GetDetailUserGinHandler(c *gin.Context) {
	users := nuh.UserService.GetUser()
	index, _ := strconv.Atoi(c.Param("identifier"))
	user := users[index]

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func (nuh *UserHandler) RegisterUserGinHandler(c *gin.Context) {
	name := c.PostForm("name")
	nuh.UserService.Register(&IUserService.User{Name: name})
	msg := fmt.Sprintf("%s Berhasil ditambahkan", name)
	c.JSON(http.StatusAccepted, gin.H{
		"message": msg,
	})
	return
}
