package routes

import (
	"hacktiv8/golang-basic/assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

func WeatherHttpHandler() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("views/*.tmpl")
	r.GET("/weather", controllers.GetWeather)

	return r
}
