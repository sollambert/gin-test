package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sollambert/gin-test/routes"
)

func main() {
	r := gin.Default()

	r.POST("/user", routes.AddUser)

	r.Run(":8080")
}
