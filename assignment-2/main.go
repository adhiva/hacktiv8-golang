package main

import (
	"fmt"
	"os"

	"hacktiv8/golang-basic/assignment-2/lib"

	"hacktiv8/golang-basic/assignment-2/routes"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		os.Exit(1)
	}
}

func main() {
	lib.InitDatabase()

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	r := routes.CreateRouter()
	r.Run(port)
}
