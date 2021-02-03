package main

import (
	"bigqiqi/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/passport/v1")
	{
		v1.POST("/login", user.Login)
	}
	router.Run(":8080")
}
