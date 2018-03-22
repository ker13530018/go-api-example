package main

import (
	"go-api-example/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	u := user.Users{}
	// user section
	r.GET("/users/:uid", u.Get)

	r.GET("/users", u.GetAll)

	r.POST("/users", u.Post)

	r.Run(":90")
}
