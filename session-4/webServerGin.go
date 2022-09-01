package session4

import (
	session3 "hacktiv8/golang-basic/session-3"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinRouter() {
	db := []*session3.User{}
	userService := session3.NewUserService(db)
	userHandler := NewUserHandler(userService)

	router := gin.Default()

	groupRouter := router.Group("/gin-handler")
	groupRouter.GET("/hello-gin", HelloGin)
	groupRouter.POST("/register", userHandler.RegisterUserGinHandler)
	groupRouter.GET("/user/:identifier", userHandler.GetDetailUserGinHandler)
	groupRouter.GET("/user", userHandler.GetUserGinHandler)

	router.Run(PORT)
}

func HelloGin(c *gin.Context) {
	msg := "Hello World, it's time to Go!"
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	return
}
